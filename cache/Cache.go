package cache

import (
	"os"

	"github.com/astaxie/beego/cache"
	// _ "github.com/astaxie/beego/cache/redis"
)

var fileCache cache.Cache

//初始化缓存
func init() {

	dir, _ := os.Getwd()
	fileCache, _ = cache.NewCache("file", `{"CachePath":"`+dir+`/runtime/cache","FileSuffix":".cache","DirectoryLevel":"2","EmbedExpiry":"120"}`)

}

func GetFileCache() cache.Cache {
	return fileCache
}
