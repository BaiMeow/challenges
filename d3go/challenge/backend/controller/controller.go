package controller

import (
	"d3go/model"
	"d3go/service/auth"
	"d3go/service/upload"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"path"
)

type Resp struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	Data       any    `json:"data"`
}

var ErrFormatError = "format error"
var ErrInternalServer = "internal server error"

func Login(c *gin.Context) {
	u := &model.User{}
	err := c.ShouldBindJSON(u)
	if err != nil {
		c.JSON(200, Resp{
			StatusCode: -1,
			StatusMsg:  ErrFormatError,
		})
		return
	}
	permission, err := auth.Login(u)
	if err != nil {
		c.JSON(200, Resp{
			StatusCode: -1,
			StatusMsg:  ErrInternalServer,
		})
		return
	}
	session := sessions.Default(c)
	switch permission {
	case auth.UnAuthed:
		c.JSON(200, Resp{
			StatusCode: -1,
			StatusMsg:  "login fail",
		})
	case auth.User:
		session.Set("admin", false)
		session.Save()
		c.JSON(200, Resp{
			StatusCode: 0,
			StatusMsg:  "login success",
		})
	case auth.Admin:
		session.Set("admin", true)
		session.Save()
		c.JSON(200, Resp{
			StatusCode: 0,
			StatusMsg:  "login as admin success",
		})
	}
}

func Register(c *gin.Context) {
	var u model.User
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(200, Resp{
			StatusCode: -1,
			StatusMsg:  ErrFormatError,
		})
		return
	}
	err = auth.Register(&u)
	if err != nil {
		c.JSON(200, Resp{
			StatusCode: -1,
			StatusMsg:  ErrInternalServer,
		})
		return
	}
	c.JSON(200, Resp{
		StatusCode: 0,
		StatusMsg:  "register success",
	})
}

func Upload(c *gin.Context) {
	f, err := c.FormFile("file")
	if err != nil {
		c.JSON(500, Resp{
			StatusCode: -1,
			StatusMsg:  "upload fail",
		})
		return
	}

	if (f.Header.Get("Content-Type") != "application/zip" && f.Header.Get("Content-Type") != "application/x-zip-compressed") || path.Ext(f.Filename) != ".zip" {
		c.JSON(500, Resp{
			StatusCode: -1,
			StatusMsg:  "not a zip file",
		})
		return
	}

	uu := uuid.New()

	zipPath := path.Join("upload", uu.String()+".zip")
	if err := c.SaveUploadedFile(f, zipPath); err != nil {
		c.JSON(500, Resp{
			StatusCode: -1,
			StatusMsg:  "save zip fail",
		})
		return
	}

	tree, err := upload.Unzip(zipPath, path.Join("unzipped", uu.String()))
	if err != nil {
		c.JSON(500, Resp{
			StatusCode: -1,
			StatusMsg:  "upload fail",
		})
		return
	}

	c.JSON(200, Resp{
		StatusCode: 0,
		StatusMsg:  "upload success",
		Data:       tree.Children,
	})
}
