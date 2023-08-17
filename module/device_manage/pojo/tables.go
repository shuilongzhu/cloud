package pojo

import "time"

type ProductDesignList struct {
	ProductDesign string `gorm:"column:product_design"`
}

type DeviceInfosToTableORM struct {
	Id                    int64     `gorm:"primary_key; column:id"`
	DeviceSn              string    `gorm:"column:device_sn"`         //`device_id` varchar(32)设备ID',
	DeviceName            string    `gorm:"column:name"`              //`device_name` varchar(32)设备名称',
	DeviceGbId            string    `gorm:"column:gb_id"`             //`device_gb_id` varchar(32)设备国标ID',
	DeviceGroupId         string    `gorm:"column:group_id"`          //`device_group_id` varchar(32)设备分组ID',
	UserGroupId           string    `gorm:"column:user_group_id"`     //`user_group_id` varchar(32)用户分组id',
	DeviceGroupName       string    `gorm:"column:device_group_name"` //`device_group_name` varchar(32)设备分组名称',
	ProductType           string    `gorm:"column:product_type"`      //`product_type` 产品型号：FWB01
	ProductImgUrl         string    `gorm:"column:product_img_url"`   //`product_img_url` varchar(32)产品图片存储地址',
	DeviceStatus          int       `gorm:"column:status"`            //`device_status` int   设备状态 1:在线；2:不在线；3:未激活；4暂停服务
	DeviceInstallPosition string    `gorm:"column:location"`          //`device_install_position` varchar(32)设备安装位置',
	CreatedTime           time.Time `gorm:"column:created_time"`      //`created_time` datetime创建时间',
	UpdateTime            time.Time `gorm:"column:update_time"`       //`update_time` datetime更新时间',
	DeleteTag             int       `gorm:"column:delete_tag"`        //`delete_tag` int DEFAULT '0'  '删除 0:正常；1:已删除；',
	AlternateField1       string    `gorm:"column:alternate_field1"`  //`alternate_field1` varchar(32)备用字段1',2:已激活；3:未激活；4暂停服务
	DeviceType            int       `gorm:"column:type"`              //  `device_type` int DEFAULT NULL COMMENT '设备类型',1 信通网关 2艾瑞抓拍机
	Uuid                  string    `gorm:"column:uuid"`
	IsOpen                int       `gorm:"column:is_open"`
	Licence               string    `gorm:"column:licence"`
}

type DeviceInfosAndProductInfos struct {
	DeviceSn              string    `gorm:"column:device_sn"`         //`device_id` varchar(32)设备ID',
	DeviceName            string    `gorm:"column:name"`              //`device_name` varchar(32)设备名称',
	DeviceGbId            string    `gorm:"column:gb_id"`             //`device_gb_id` varchar(32)设备国标ID',
	DeviceGroupId         string    `gorm:"column:group_id"`          //`device_group_id` varchar(32)设备分组ID',
	UserGroupId           string    `gorm:"column:user_group_id"`     //`user_group_id` varchar(32)用户分组id',
	DeviceGroupName       string    `gorm:"column:device_group_name"` //`device_group_name` varchar(32)设备分组名称',
	ProductType           string    `gorm:"column:product_type"`      //`product_type` 产品型号：FWB01
	ProductImgUrl         string    `gorm:"column:product_img_url"`   //`product_img_url` varchar(32)产品图片存储地址',
	DeviceStatus          int       `gorm:"column:status"`            //`device_status` int   设备状态 1:在线；2:不在线；3:未激活；4暂停服务
	DeviceInstallPosition string    `gorm:"column:location"`          //`device_install_position` varchar(32)设备安装位置',
	CreatedTime           time.Time `gorm:"column:created_time"`      //`created_time` datetime创建时间',
	UpdateTime            time.Time `gorm:"column:update_time"`       //`update_time` datetime更新时间',
	DeleteTag             int       `gorm:"column:delete_tag"`        //`delete_tag` int DEFAULT '0'  '删除 0:正常；1:已删除；',
	AlternateField1       string    `gorm:"column:alternate_field1"`  //`alternate_field1` varchar(32)备用字段1',2:已激活；3:未激活；4暂停服务
	DeviceType            int       `gorm:"column:type"`              //  `device_type` int DEFAULT NULL COMMENT '设备类型',1 信通网关 2艾瑞抓拍机
	Uuid                  string    `gorm:"column:uuid"`
	IsOpen                int       `gorm:"column:is_open"`
	Licence               string    `gorm:"column:licence"`
	ProductDesign         string    `gorm:"column:product_design"` // 产品形态  手机、服务器、边缘计算盒
	ProductName           string    `gorm:"column:product_name"`
	ProductLogo           string    `gorm:"column:product_logo"` // 产品图片地址,
	XpuType               string    `gorm:"column:xpu_type"`
	XpuVersion            string    `gorm:"column:xpu_version"`
	AiriaShip             string    `gorm:"column:airia_ship"`
}

// 设备分组表 device_group
type DeviceGroupTables struct {
	Id              int64     `gorm:"primary_key; column:id"`   //`id` int NOT NULL AUTO_INCREMENT  '自增ID',
	DeviceGroupId   string    `gorm:"column:device_group_id"`   //`device_group_id` varchar(32)设备分组ID',
	UserGroupId     string    `gorm:"column:user_group_id"`     //`user_group_id` varchar(32)用户分组id',
	DeviceGroupName string    `gorm:"column:device_group_name"` //`device_group_name` varchar(32)设备分组名称',
	DeviceGroupNote string    `gorm:"column:device_group_note"` //`device_group_note` varchar(128)设备分组描述',
	CreatedTime     time.Time `gorm:"column:created_time"`      //`created_time` datetime创建时间',
	UpdateTime      time.Time `gorm:"column:update_time"`       //`update_time` datetime更新时间',
	DeleteTag       int       `gorm:"column:delete_tag"`        //`delete_tag` int DEFAULT '0'  '删除 0:正常；1：已删除；',
}

// 设备事件信息表 device_event_info
type DeviceEventInfo struct {
	Id int `gorm:"AUTO_INCREMENT; primary_key; column:id"` //`id` int NOT NULL AUTO_INCREMENT  '自增ID 用于标识事件',

	DeviceId      string    `gorm:"column:device_id"`       //`device_id` varchar(32)设备ID',
	DeviceName    string    `gorm:"column:device_name"`     //`device_name` varchar(32)设备名称',
	EventTypeName string    `gorm:"column:event_type_name"` //`event_type_name` varchar(32) CHARACTER SET utf8事件类型',
	EventNote     string    `gorm:"column:event_note"`      //`event_note` varchar(128) CHARACTER SET utf8事件描述 事件名称',
	EventStatus   int       `gorm:"column:event_status"`    //`event_status` int事件状态',
	DisposeTime   string    `gorm:"column:dispose_time"`    //`dispose_time` varchar(32) CHARACTER SET utf8处理时间',
	TimeStamp     int       `gorm:"column:time_stamp"`      //时间戳
	CreatedTime   time.Time `gorm:"column:created_time"`    //`created_time` datetime创建时间',
	UpdateTime    time.Time `gorm:"column:update_time"`     //`update_time` datetime更新时间',
}

// 设备通道信息表 channel_info
type ChannelInfo struct {
	Id int `gorm:"AUTO_INCREMENT; primary_key; column:id"` //`id` int NOT NULL AUTO_INCREMENT  '自增ID',

	DeviceId          string    `gorm:"column:device_id"`           //`device_id` varchar(32)设备ID',
	DeviceName        string    `gorm:"column:device_name"`         //`device_name` varchar(32)设备名称',
	ChannelId         string    `gorm:"column:channel_id"`          //`channel_id` varchar(32)通道ID',
	ChannelName       string    `gorm:"column:channel_name"`        //`channel_name` varchar(32)通道名称',
	ChannelIp         string    `gorm:"column:channel_ip"`          //`channel_ip` varchar(32)通道IP地址',
	ChannelUsername   string    `gorm:"column:channel_username"`    //`channel_username` varchar(32)通道用户名',
	ChannelPassword   string    `gorm:"column:channel_password"`    //`channel_password` varchar(32)通道登录密码',
	ChannelAccessType string    `gorm:"column:channel_access_type"` //`channel_access_type` varchar(32)通道接入类型',
	TakeWay           string    `gorm:"column:take_way"`            //`take_way` varchar(32)取流方式',
	ChannelAddress    string    `gorm:"column:channel_address"`     //`channel_address` varchar(32)通道安装地址',
	ChannelUrl        string    `gorm:"column:channel_url"`         //`channel_url` varchar(32)通道直播流地址',
	ChannelModel      string    `gorm:"column:channel_model"`       //`channel_model` varchar(32)通道选择模型',
	SetupTime         string    `gorm:"column:setup_time"`          //`setup_time` varchar(32)配置时间',
	ChannelStatus     string    `gorm:"column:channel_status"`      //`channel_status` varchar(32)通道状态',1:在线显示绿色；2:不在线显示红色；3关闭显示灰色'
	AccessTime        time.Time `gorm:"column:access_time"`         //`access_time` datetime接入时间',
	CreatedTime       time.Time `gorm:"column:created_time"`        //`created_time` datetime创建时间',
	DeleteTime        time.Time `gorm:"column:delete_time"`         //`delete_time` datetime删除时间',
	DeleteTag         int       `gorm:"column:delete_tag"`          //`delete_tag` int删除标记 0:正常；1：已删除；',
	TakeAddress       string    `gorm:"column:take_address"`        //`alternate_field1` varchar(32)取流地址',
	AlternateField2   string    `gorm:"column:alternate_field2"`    //`alternate_field2` varchar(32)备用字段2',
}

// GuardActionLog 打卡操作记录表
type GuardActionLog struct {
	ActionId   int64  `gorm:"column:action_id"`
	DeviceSn   string `gorm:"column:device_sn"`
	DeviceName string `gorm:"column:device_name"`
	UserId     string `gorm:"column:user_id"`
	UserName   string `gorm:"column:user_name"`
	UserNum    string `gorm:"column:user_num"`
	RoleName   string `gorm:"column:role_name"`
	Action     string `gorm:"column:action"`
	AcTime     int64  `gorm:"column:ac_time"`
	NextAction int64  `gorm:"column:next_action"`
}
type GuardActionAddTable struct {
	ActionId   int64  `gorm:"column:action_id"`
	DeviceSn   string `gorm:"column:device_sn"`
	UserId     string `gorm:"column:user_id"`
	UserName   string `gorm:"column:user_name"`
	UserNum    string `gorm:"column:user_num"`
	RoleName   string `gorm:"column:role_name"`
	Action     string `gorm:"column:action"`
	AcTime     int64  `gorm:"column:ac_time"`
	NextAction int64  `gorm:"column:next_action"`
}

// GuardClockLogAndDeviceInfos 打卡操作记录关联设备详情信息表
type GuardClockLogAndDeviceInfos struct {
	ActionId     int64  `gorm:"column:action_id" json:"-"`
	DeviceSn     string `gorm:"column:device_sn" json:"device_sn"`
	DeviceName   string `gorm:"column:device_name" json:"device_name"`
	Location     string `gorm:"column:location" json:"location"`
	DeviceStatus int    `gorm:"column:status" json:"device_status"`
	AcTime       string `gorm:"column:-" json:"ac_time"`
	Count        int64  `gorm:"column:-" json:"count"`
	Status       int    `gorm:"column:-" json:"status"`
}

type ClockInLogAndDeviceInfosResp struct {
	List  []GuardClockLogAndDeviceInfos `json:"list"`
	Count int                           `json:"count"`
}
