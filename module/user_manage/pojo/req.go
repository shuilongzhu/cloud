package pojo

type LoginReq struct {
	UserName         string `json:"user_name" binding:"required"`       //用户名
	Password         string `json:"password" binding:"required"`        //密码
	VerificationCode string `json:"verification_code" example:"123456"` //验证码
	LoginIp          string `json:"login_ip" example:"127.0.0.1"`       //请求ip地址
}

type RegisterReq struct {
	UserName      string `json:"user_name" binding:"required"`
	UserPassword  string `json:"user_password" binding:"required"`
	Num           string `json:"num" binding:"required"`
	Email         string `json:"email" example:"abc123@163.com"`
	UserGroupName string `json:"user_group_name" binding:"required"`
	UserGroupId   string `json:"user_group_id" binding:"required"`
	Role          int    `json:"role" binding:"required"`
	NickName      string `json:"nick_name" binding:"required"`
}
