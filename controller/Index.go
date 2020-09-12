package controller

import (
	"gin/cache"
	"gin/common"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	bm := cache.GetFileCache()
	if bm.IsExist("index") {
		cacheResult := bm.Get("index")
		c.JSON(common.Success("获取成功!", cacheResult))
		return
	}
	user, _ := c.Get("user")
	if usermap, ok := user.(jwt.MapClaims); ok {
		if uid, ok := usermap["uid"].(float64); ok {
			uid := int(uid)
			// var result map[string]interface{}
			// common.GetDb().Table("template").Where("id > ?", uid).Find(&result)

			//逻辑
			/*
				start

				end
			*/
			bm.Put("index", uid, 10*time.Minute)
			time.Sleep(4 * time.Second)
			c.JSON(common.Success("获取成功!", uid))
			return
		}

	}
	c.JSON(common.ServiceError("map解析错误！", ""))
}
