package error_code

const (
	ErrUserNotExist = UserManageErrorCodeBase + 1 + iota
	ErrUserBeDisabled
	ErrUserPassword
	ErrUserExit
	ErrRegisterErr
	ErrIsNotExistUserGroup
)

func init() {
	StatusCode[ErrUserNotExist] = "该用户不存在，请先创建此用户"
	StatusCode[ErrUserBeDisabled] = "该用户已被禁用"
	StatusCode[ErrUserPassword] = "登陆密码错误，请重新输一遍"
	StatusCode[ErrUserExit] = "该用户已存在"
	StatusCode[ErrRegisterErr] = "注册失败"
	StatusCode[ErrIsNotExistUserGroup] = "该用户组不存在,请先添加此用户组"
}
