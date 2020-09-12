package meddleware

import (
	"gin/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authorization(c *gin.Context) {

	if _, ok := c.Request.Header["Authorization"]; ok {

		tokenString := c.Request.Header["Authorization"][len(c.Request.Header["Authorization"])-1][7:] //Bearer

		authorizationMap := common.ParseHStoken(tokenString)

		if _, ok := authorizationMap["uid"]; ok {
			c.Set("user", authorizationMap)
			c.Next()
			return

		}

	}
	c.Abort()
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "无权限!",
	})
	return

}
