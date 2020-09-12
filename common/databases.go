package common

import (
	"fmt"
	"gin/config"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", config.Conf.Mysql.Username+":"+config.Conf.Mysql.Password+"@tcp("+config.Conf.Mysql.Host+":"+config.Conf.Mysql.Port+")/"+config.Conf.Mysql.Dbname+"?charset="+config.Conf.Mysql.Charset+"&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(0)
	}
	//设置全局表名禁用复数
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(config.Conf.Mysql.MaxOpenConns)

	// defer db.Close()
}
func GetDb() *gorm.DB {
	return db
}
