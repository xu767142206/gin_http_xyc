package main

import (
	"gin/common"
	"gin/config"
	_ "gin/log"
	"gin/router"
)

func main() {
	defer common.GetDb().Close()
	router.Router.Run(":" + config.Conf.Server.Port)

}
