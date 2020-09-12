package controller

import (
	"gin/common"
	"gin/model"
	"time"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	account := c.PostForm("user")
	pwd := c.PostForm("pwd")

	if account == "" || pwd == "" {
		c.JSON(common.ClientError("账号或者密码为空！", ""))
		return
	}

	user := &model.User{UserName: account, PassWord: pwd}
	result := user.Find()
	if result.Id == 0 {
		c.JSON(common.ClientError("未找到该账号", ""))
		return
	}
	token, err := common.CreateToken(result.UserName, result.PassWord, result.Id)
	if err != nil {
		c.JSON(common.ServiceError("token fiall", ""))
		return
	}
	c.JSON(common.Success("获取成功！", map[string]interface{}{
		"token":   token,
		"endtime": time.Now().Add(time.Minute * 15).Unix(),
	}))

}

func Register(c *gin.Context) {
	account := c.PostForm("user")
	pwd := c.PostForm("pwd")

	if account == "" || pwd == "" {
		c.JSON(common.ClientError("账号或者密码为空！", ""))
		return
	}

	userFind := &model.User{UserName: account}
	userResult := userFind.Find()
	if userResult.Id > 0 {
		c.JSON(common.ClientError("该账号已经注册", ""))
		return
	}

	user := &model.User{UserName: account, PassWord: pwd}
	user.Add()
	if user.Id > 0 {
		c.JSON(common.Success("创建成功", model.UserOto{Id: user.Id, UserName: user.UserName, CreatedAt: user.CreatedAt}))
		return
	}

	c.JSON(common.ServiceError("创建账号失败!", ""))
}
