package controller

import (
	"gin/common"
	"gin/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

//用户列表
func UserList(c *gin.Context) {
	user, _ := c.GetQuery("user")
	result := model.GetUserInstance().GetUserInfoArr(user)
	c.JSON(common.Success("ok", result))
}

//删除用户
func UserDel(c *gin.Context) {

	idstr, _ := c.GetQuery("id")
	id, _ := strconv.Atoi(idstr)
	if id != 0 {
		user := &model.User{Id: id}
		user.Del()
		c.JSON(common.Success("删除成功!", ""))
		return
	}
	c.JSON(common.ClientError("id 未定义!", ""))
}
