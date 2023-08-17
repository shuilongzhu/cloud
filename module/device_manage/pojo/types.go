package pojo

type DeviceGroup struct {
	Id              string   `json:"id"` //主键id
	DeviceGroupId   string   `json:"device_group_id" gorm:"column:device_group_id"`
	DeviceGroupName string   `json:"device_group_name" gorm:"column:device_group_name"`
	Level           string   `json:"level" gorm:"column:level"`
	Devices         []Device `json:"children"`
}

func (DeviceGroup) TableName() string {
	return "device_group"
}

type Device struct {
	Id                   string    `json:"id"` //主键id
	DeviceGroupId        string    `json:"device_group_id" gorm:"column:group_id"`
	DeviceSn             string    `json:"device_id" gorm:"column:device_sn"`
	DeviceName           string    `json:"name" gorm:"column:name"`
	DeviceStatus         int       `json:"device_status" gorm:"column:status"` // '设备状态 1:在线；2:不在线；',
	Level                string    `json:"level" gorm:"column:device_level"`
	DeviceInstallAddress string    `json:"device_install_address" gorm:"column:location"`
	Channels             []Channel `json:"children" gorm:"foreignKey:device_id;references:device_sn""`
}

func (Device) TableName() string {
	return "device_info"
}

type Channel struct {
	Id         string `json:"id"` //主键id
	DeviceId   string `json:"device_id" gorm:"column:device_id"`
	ChanId     string `json:"chan_id" gorm:"column:channel_id"`
	ChanName   string `json:"name" gorm:"column:channel_name"`
	ChanStatus int    `json:"chan_status" gorm:"column:channel_status"` //'通道状态',1:在线；2:不在线；',
	Level      string `json:"level" gorm:"column:level"`
}

func (Channel) TableName() string {
	return "channel_info"
}

type DeviceState struct {
	State int `json:"state" gorm:"column:state"`
	Count int `json:"count" gorm:"column:count"`
}

type ChannelState struct {
	ChannelStatus int `json:"channel_status" gorm:"column:channel_status"`
	Count         int `json:"count" gorm:"column:count"`
}
