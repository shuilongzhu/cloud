package dao

import (
	"cloud/common/db"
	"cloud/common/pojo"
	"time"
)

// UserLoginInfoCreate @description: 用户登陆记录
// @parameter userId
// @parameter loginType
// @return code
func UserLoginInfoCreate(userId string, loginType int) (code int) {
	info := pojo.UserLoginInfo{
		UserId:      userId,
		LoginTime:   time.Now(),
		LoginDevice: loginType,
		CreatedTime: time.Now(),
	}
	code = db.MyDb.CommonCreate("user_login_info", &info)
	return
}
