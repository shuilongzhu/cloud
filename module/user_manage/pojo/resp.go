package pojo

type LoginResp struct {
	UserId      string `json:"user_id"`
	UserName    string `json:"user_name"`
	UserGroupId string `json:"user_group_id"`
	TokenString string `json:"token"`
	RoleName    string `json:"role_name"`
	IsAdmin     bool   `json:"is_admin"`
	UserLogo    string `json:"user_logo"`
	Num         string `json:"num"`
	Email       string `json:"email"`
}
