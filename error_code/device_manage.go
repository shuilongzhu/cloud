package error_code

const (
	ErrDeviceTreeSelect = DeviceManageErrorCodeBase + 1 + iota
	ErrIsExistDeviceGroup
	ErrIsNotExistDeviceGroup
	ErrUngroupNotAllowOperate
	ErrIsExistDeviceGroupName
	ErrDeviceNotAuth
	ErrDeviceNotRegister
)

func init() {
	StatusCode[ErrDeviceTreeSelect] = "设备树设备通道查询错误"
	StatusCode[ErrIsExistDeviceGroup] = "该设备组已经存在"
	StatusCode[ErrIsNotExistDeviceGroup] = "该设备组不存在，请先添加"
	StatusCode[ErrUngroupNotAllowOperate] = "未分组设备分组不允许操作"
	StatusCode[ErrIsExistDeviceGroupName] = "已经存在此设备分组名称，请换一个"
	StatusCode[ErrDeviceNotAuth] = "该设备未授权，请先授权"
	StatusCode[ErrDeviceNotRegister] = "该设备序列号未注册，请先注册"
}
