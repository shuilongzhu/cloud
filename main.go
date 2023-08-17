package main

import (
	"cloud/common/init_data"
	"cloud/common/timed_task"
	"cloud/common/utils/error"
	"cloud/config"
	"cloud/router"
	"flag"
	"fmt"
	"os"
)

type App struct {
	Host       string
	Port       string
	ConfigPath string
}

var configPath string
var showVersion bool

// @title           VideoCompressCloud API
// @version         3.6.5
// @description     video-compress-cloud Project API
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      172.10.50.41:5022
// @BasePath  /cloud/api

// @securityDefinitions.basic  BasicAuth

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
// @description					Description for what is this security definition being used

// @securitydefinitions.oauth2.application  OAuth2Application
// @tokenUrl                                https://example.com/oauth/token
// @scope.write                             Grants write access
// @scope.admin                             Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit  OAuth2Implicit
// @authorizationUrl                     https://example.com/oauth/authorize
// @scope.write                          Grants write access
// @scope.admin                          Grants read and write access to administrative information

// @securitydefinitions.oauth2.password  OAuth2Password
// @tokenUrl                             https://example.com/oauth/token
// @scope.read                           Grants read access
// @scope.write                          Grants write access
// @scope.admin                          Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode  OAuth2AccessCode
// @tokenUrl                               https://example.com/oauth/token
// @authorizationUrl                       https://example.com/oauth/authorize
// @scope.admin                            Grants read and write access to administrative information

func init() {
	flag.BoolVar(&showVersion, "version", false, "print program build version")
	flag.StringVar(&configPath, "c", "conf/config.toml", "path of configure file.")
	flag.Parse()
}

func main() {
	appServer := &App{}
	if showVersion {
		fmt.Println(config.Version)
		os.Exit(0)
	}
	appServer.ConfigPath = configPath
	//启动项目
	appServer.Start()
}

func (app *App) Start() {
	//  项目异常捕获
	defer error.CatchGoError("Main", true)

	//初始化全局变量
	if err := config.MainInit(app.ConfigPath); nil != err {
		return
	}

	//初始化连接数据库
	if err := config.ConnectMysql(); nil != err {
		return
	}

	//初始化数据
	go func() {
		defer error.CatchGoError("CommonInit", true)
		init_data.CommonInit()
	}()

	//定时任务
	go func() {
		defer error.CatchGoError("GoCron", true)
		timed_task.GoCron()
	}()

	router.RouteInit()
}
