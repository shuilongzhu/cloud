package service

import (
	commonDb "cloud/common/db"
	commonPojo "cloud/common/pojo"
	"cloud/common/utils/ctx"
	"cloud/common/utils/data_convert"
	"cloud/error_code"
	"cloud/module/user_manage/dao"
	"cloud/module/user_manage/pojo"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

// Login @description: 用户登录逻辑处理
// @parameter req
// @return code
func Login(req pojo.LoginReq) (code int, result pojo.LoginResp) {
	result = pojo.LoginResp{}

	//1、查询用户名称是否存在
	code, userInfo := commonDb.UserInfoSelect(req.UserName)
	if error_code.Success != code {
		return
	}

	if 0 == userInfo.Id { //此用户不存在
		code = error_code.ErrUserNotExist
		return
	}

	if 2 == userInfo.UserState { //用户状态 1：正常；2：已禁用
		code = error_code.ErrUserBeDisabled
		return
	}

	//2. 校验密码
	if userInfo.UserPassword != data_convert.GetMD5Encode(req.Password) {
		code = error_code.ErrUserPassword
		return
	}

	//3.用户登陆记录
	if code = dao.UserLoginInfoCreate(userInfo.UserId, 1); error_code.Success != code {
		return
	}

	//4. 用户token生成
	claims := &commonPojo.Claims{
		UserId:      userInfo.UserId,
		UserGroupId: userInfo.UserGroupId,
		UserName:    req.UserName,
		RoleId:      userInfo.Role,
		RoleName:    commonPojo.RoleEnum[userInfo.Role],
		UserNum:     userInfo.Number,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1",  // 签名颁发者
			Subject:   "user token", //签名主题
		},
	}
	code, result.TokenString = ctx.SettingToken(claims)
	if error_code.Success != code {
		return
	}

	// 5.数据整理返回
	result.UserId = userInfo.UserId
	result.UserName = req.UserName
	result.UserGroupId = userInfo.UserGroupId
	result.RoleName = commonPojo.RoleEnum[userInfo.Role]
	result.UserLogo = userInfo.AvatarStorageAddress
	result.Num = userInfo.Number
	result.Email = userInfo.Email
	if userInfo.Role == 1 {
		result.IsAdmin = true
	} else {
		result.IsAdmin = false
	}

	return
}

// Register @description: 用户注册逻辑处理
// @parameter req
// @return code
func Register(req pojo.RegisterReq) (code int) {
	//查询是否已经注册
	code, userInfo := commonDb.UserInfoSelect(req.UserName)
	if error_code.Success != code {
		return
	}

	if 0 != userInfo.Id { //此用户已存在
		return error_code.ErrUserExit
	}
	userInfo.UserName = req.UserName
	userInfo.UserId = commonPojo.IntPrefix + strconv.FormatInt(data_convert.NextId(), 10)
	userInfo.UserGroupId = req.UserGroupId
	userInfo.Nickname = req.NickName
	userInfo.UserGroupName = req.UserGroupName
	userInfo.Role = req.Role
	userInfo.Number = req.Num
	userInfo.UserPassword = data_convert.GetMD5Encode(req.UserPassword)
	userInfo.Email = req.Email
	userInfo.CreatedTime = time.Now()
	userInfo.UpdateTime = time.Now()
	userInfo.LoginTime = time.Now()
	userInfo.UserState = 1 //用户正常 1：正常 2：禁用
	//插入用户信息
	return commonDb.UserInfoCreate(userInfo)
}
