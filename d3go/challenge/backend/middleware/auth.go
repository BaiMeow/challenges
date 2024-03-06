package middleware

import (
	"d3go/controller"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		IsAdmin, ok := session.Get("admin").(bool)
		if !ok {
			c.AbortWithStatusJSON(200, controller.Resp{
				StatusCode: -1,
				StatusMsg:  "unauthorized",
			})
			return
		}
		if !IsAdmin {
			c.AbortWithStatusJSON(200, controller.Resp{
				StatusCode: -2,
				StatusMsg:  "only admin can access",
			})
			return
		}
	}
}
