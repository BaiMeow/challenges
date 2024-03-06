package main

import (
	"context"
	"log"
	"net/url"
	"os"
	"regexp"
	"sync"
	"text/template"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/html"
)

var re = regexp.MustCompile(`script|file|on`)

var lock sync.Mutex

func main() {
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.NoSandbox, chromedp.DisableGPU)...)
	defer cancel()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		tmplStr := c.Query("tmpl")
		if tmplStr == "" {
			tmplStr = defaultTmpl
		} else {
			if re.MatchString(tmplStr) {
				c.String(403, "tmpl contains invalid word")
				return
			}
			if len(tmplStr) > 50 {
				c.String(403, "tmpl is too long")
				return
			}
			tmplStr = html.EscapeString(tmplStr)
		}
		tmpl, err := template.New("resp").Parse(tmplStr)
		if err != nil {
			c.String(500, "parse template error: %v", err)
			return
		}
		if err := tmpl.Execute(c.Writer, c); err != nil {
			c.String(500, "execute template error: %v", err)
		}
	})

	r.GET("/bot", func(c *gin.Context) {
		rawURL := c.Query("url")
		u, err := url.Parse(rawURL)
		if err != nil {
			c.String(403, "url is invalid")
			return
		}
		if u.Host != "127.0.0.1:8080" {
			c.String(403, "host is invalid")
			return
		}
		go func() {
			lock.Lock()
			defer lock.Unlock()

			ctx, cancel := chromedp.NewContext(allocCtx,
				chromedp.WithBrowserOption(chromedp.WithDialTimeout(10*time.Second)),
			)
			defer cancel()
			ctx, _ = context.WithTimeout(ctx, 20*time.Second)
			if err := chromedp.Run(ctx,
				chromedp.Navigate(u.String()),
				chromedp.Sleep(time.Second*10),
			); err != nil {
				log.Println(err)
			}
		}()
		c.String(200, "bot will visit it.")
	})

	r.GET("/flag", func(c *gin.Context) {
		if c.RemoteIP() != "127.0.0.1" {
			c.String(403, "you are not localhost")
			return
		}
		flag, err := os.ReadFile("/flag")
		if err != nil {
			c.String(500, "read flag error")
			return
		}
		c.SetCookie("flag", string(flag), 3600, "/", "", false, true)
		c.Status(200)
	})
	r.Run(":8080")
}

const defaultTmpl = `
<!DOCTYPE html>
<html>
<head>
	<title>YOU ARE</title>
</head>
<body>
	<div>欢迎来自 {{.ClientIP}} 的朋友</div>
	<div>你的 User-Agent 是 {{.GetHeader "User-Agent"}}</div>
	<div>flag在bot手上，想办法偷过来</div>
</body>
`
