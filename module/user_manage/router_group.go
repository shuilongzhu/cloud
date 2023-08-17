package user_manage

import (
	"cloud/module/user_manage/controller"
	"github.com/gin-gonic/gin"
)

// UserAPIInit @description:用户管理路由，baseUrl:/cloud/api/user
// @parameter user
func UserAPIInit(user *gin.RouterGroup) {
	user.POST("/login", controller.Login)       // 用户登录
	user.POST("/register", controller.Register) // 用户注册
}
