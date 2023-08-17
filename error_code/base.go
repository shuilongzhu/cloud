package error_code

// 标注化服务层错误码的统一,100001(六位)
// error 错误信息属于内部错误具体信息
// mapString 是针对外用户输出信息

const (
	Success = 200

	CommonErrorCodeBase       = 0 + iota*100000 //公共基础错误码
	UserManageErrorCodeBase                     //用户管理基础错误码
	DeviceManageErrorCodeBase                   //设备管理基础错误码
)

var StatusCode = map[int]string{}

func init() {
	StatusCode[Success] = "Successfully"
}
