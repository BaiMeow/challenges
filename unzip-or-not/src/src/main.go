package main

import (
	"io"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	// disable the output of gin
	rand.Seed(time.Now().UnixNano())
}

func main() {
	if err := os.MkdirAll("uploads", 0777); err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Check the attachment for next step.")
	})
	r.POST("/unzip", func(c *gin.Context) {
		fh, err := c.FormFile("file")
		if err != nil {
			log.Println(err)
			c.String(500, "读取文件失败")
			return
		}

		savepath := filepath.Clean(filepath.Join("uploads", strings.ReplaceAll(fh.Filename, "../", "")))

		file, err := fh.Open()
		if err != nil {
			log.Println(err)
			c.String(500, "读取文件失败")
			return
		}
		defer file.Close()

		saveto, err := os.OpenFile(savepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			log.Println(err)
			c.String(500, "保存文件失败")
			return
		}
		defer saveto.Close()

		_, err = io.Copy(saveto, file)
		if err != nil {
			log.Println(err)
			c.String(500, "保存文件失败")
			return
		}

		res, err := Unzip(savepath)
		if err != nil {
			log.Println(err)
			c.String(500, "解压失败")
			return
		}

		// show unzip log
		c.String(200, string(res))
	})

	r.Run(":8080")
}

func Unzip(zipfile string) ([]byte, error) {
	unzipdir := filepath.Join("output", strconv.Itoa(rand.Int()))
	output, err := exec.Command("unzip", "-n", "-v", zipfile, "-d", unzipdir).CombinedOutput()
	if err != nil {
		return nil, err
	}

	return output, nil
}
