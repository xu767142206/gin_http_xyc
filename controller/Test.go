package controller

import (
	"encoding/json"
	"fmt"
	"gin/common"
	"gin/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	user, _ := c.Get("user")

	if usermap, ok := user.(jwt.MapClaims); ok {
		result, err := json.Marshal(usermap)
		if err != nil {
			c.JSON(common.ServiceError("解析失败!", err))
			return
		}
		userOto := &model.UserOto{}
		json.Unmarshal(result, &userOto)
		fmt.Println(userOto)
		c.JSON(common.Success("获取成功!", userOto))
		return

	}
	c.JSON(common.ServiceError("解析用户失败!", ""))

}
