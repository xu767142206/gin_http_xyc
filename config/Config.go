package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type BaseConf struct {
	Server        ServerEntiy `yaml:"server"`
	Mysql         MysqlEntiy  `yaml:"mysql"`
	Authorization string      `yaml:"authorization"`
}

type MysqlEntiy struct {
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Dbname       string `yaml:"dbname"`
	Charset      string `yaml:"charset"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
}

type ServerEntiy struct {
	Port string `yaml:"port"`
}

var Conf *BaseConf = &BaseConf{}

func init() {

	dir, _ := os.Getwd()

	yamlFile, err := ioutil.ReadFile(dir + "/conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		os.Exit(0)
	}

	err = yaml.Unmarshal(yamlFile, Conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		os.Exit(0)
	}

}
