package router

import (
	"cloud/config"
	"cloud/docs"
	"cloud/middleware"
	"cloud/module/device_manage"
	"cloud/module/user_manage"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os"
)

// RouteInit @description: 路由组初始化
func RouteInit() {
	log.Println("服务端 IpAddress：", config.BaseInfoConfigData.IpAddress)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(middleware.Cors(), middleware.Auth(), middleware.ApiRequestRecord()) //gin中间件统一处理

	if true { //配置文件中swagger_state为true配置swagger
		handleSwagger(router)
	}

	// 用户管理
	apiUserManage := router.Group("/cloud/api/user")
	user_manage.UserAPIInit(apiUserManage)
	// 设备管理
	apiDeviceManage := router.Group("/cloud/api/device")
	device_manage.DeviceAPIInit(apiDeviceManage)

	http.Handle("/", router)

	runErr := router.Run(config.BaseInfoConfigData.IpAddress)
	if runErr != nil {
		log.Errorln("RouteInit 无法运行新进程！error:", runErr)
		os.Exit(0)
		return
	}
	log.Println("正运行进程！")
}

// handleSwagger 动态处理swagger展示，测试开发环境才展示swagger,生产环境不暴露swagger,待完善
func handleSwagger(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	docs.SwaggerInfo.Host = config.BaseInfoConfigData.IpAddress
}
