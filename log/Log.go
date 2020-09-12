package log

import (
	"os"

	"github.com/astaxie/beego/logs"
)

func init() {
	dir, _ := os.Getwd()
	logs.SetLogger(logs.AdapterFile, `{"filename":"`+dir+`/runtime/log/notice.log"}`)
	logs.Async()
	/*
		使用用例
		// logs.Debug("my book is bought in the year of ", 2016)

	*/

}
