package main

import (
	"crypto/rand"
	"d3go/config"
	"d3go/controller"
	"d3go/db"
	"d3go/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jpillora/overseer"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func prog(state overseer.State) {
	r := gin.Default()
	InitRouter(r)
	server := http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	go func() {
		if err := server.Serve(state.Listener); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	<-state.GracefulShutdown
	if err := server.Shutdown(nil); err != nil {
		log.Fatal(err)
	}
}

func main() {
	config.Init()
	db.Init()
	if config.Conf.AutoUpdate {
		log.Printf("Auto update enabled")
		err := overseer.RunErr(overseer.Config{
			Program: prog,
			Address: ":8080",
			Fetcher: &config.Fetch,
		})
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		r := gin.Default()
		InitRouter(r)
		if err := r.Run(":8080"); err != nil {
			log.Fatal(err)
		}
	}
}

func InitRouter(r *gin.Engine) {
	var rad [32]byte
	rand.Read(rad[:])
	store := cookie.NewStore(rad[:])
	r.Use(sessions.Sessions("mysession", store))
	r.POST("/login", controller.Login)
	r.POST("/register", controller.Register)
	r.GET("/*filepath", ServeFile)
	r.HEAD("/*filepath", ServeFile)
	admin := r.Group("/admin")
	admin.Use(middleware.Auth())
	admin.POST("/upload", controller.Upload)
}

func ServeFile(c *gin.Context) {
	// unzipped file server
	p := c.Param("filepath")
	if strings.HasPrefix(p, "/unzipped") {
		if len(p) == 9 {
			p = "/"
		} else {
			p = p[9:]
		}
		c.FileFromFS(p, http.Dir("./unzipped"))
		return
	}
	// embed static file server
	p = "/static/" + c.Param("filepath")
	c.FileFromFS(p, http.FS(Static))
	return
}
