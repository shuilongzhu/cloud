package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

// MainInit @description: 初始化项目中的全局变量
// @parameter path
// @return err
func MainInit(path string) (err error) {

	err = readConfig(path)
	if nil != err {
		log.Errorln("main init fail! error is:", err)
	}
	return
}

// readConfig @description: 读取配置文件中的配置信息
// @parameter path
// @return err
func readConfig(path string) (err error) {
	conf, err := ini.Load(path) //加载配置文件
	if nil != err {
		log.Errorln("load config file fail! error is:", err)
		return
	}

	conf.BlockMode = false

	err = conf.Section("log").MapTo(LogConfigData) //logConfig解析成结构体
	if nil != err {
		log.Println("LogConfigData parse fail!")
		return
	}

	err = conf.Section("mysql").MapTo(MysqlConfigData) //MysqlConfig解析成结构体
	if nil != err {
		log.Println("MysqlConfigData parse fail!")
		return
	}

	err = conf.Section("base_info").MapTo(BaseInfoConfigData) //BaseInfoConfig解析成结构体
	if nil != err {
		log.Println("BaseInfoConfigData parse fail!")
		return
	}

	return
}
