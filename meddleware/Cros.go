package meddleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cros(c *gin.Context) {

	method := c.Request.Method

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")

	//放行所有OPTIONS方法
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
	c.Next()
	return
}
