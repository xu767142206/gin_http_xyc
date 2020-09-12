package router

import (
	"gin/controller"
	"gin/meddleware"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {

	Router = gin.New()

	Router.Use(meddleware.Cros)

	Router.POST("/login", controller.Login)
	Router.POST("/register", controller.Register)
	Router.Use(meddleware.Authorization)
	Router.GET("/test", controller.Test)
	Router.GET("/index", controller.Index)
	Router.GET("/admin/user/list", controller.UserList)
	Router.GET("/admin/user/del", controller.UserDel)
	// Router.POST("/somePost", posting)
	// Router.PUT("/somePut", putting)
	// Router.DELETE("/someDelete", deleting)
	// Router.PATCH("/somePatch", patching)
	// Router.HEAD("/someHead", head)
	// Router.OPTIONS("/someOptions", options)

}
