package pojo

import "time"

type QueryDeviceGroupInfoResp struct {
	TotalCount      int                   `json:"total_count" example:"100"` //总记录条数
	PageNum         int                   `json:"page_num" example:"10"`     //总页数
	DeviceGroupInfo []DeviceGroupInfoResp `json:"device_group_info"`
}

type DeviceGroupInfoResp struct {
	DeviceGroupId string `json:"device_group_id" gorm:"column:device_group_id"`                         //设备分组id
	GroupName     string `json:"group_name" gorm:"column:device_group_name"`                            //分组名称
	DeviceCount   int    `json:"device_count"`                                                          //设备数量
	ChannelCount  int    `json:"channel_count"`                                                         //通道摄像头数量
	CreatedTime   string `json:"created_time" gorm:"column:created_time" example:"2021-11-01 13:04:05"` //创建时间
}

type DeviceOverviewResp struct {
	DeviceCount        int     `json:"device_count"`         //设备数量
	OnlineDeviceCount  int     `json:"online_device_count"`  //在线设备数量
	OnlineDeviceRatio  float64 `json:"online_device_ratio"`  //在线设备占比
	ChannelCount       int     `json:"channel_count"`        //通道摄像头数量
	OnlineChannelCount int     `json:"online_channel_count"` //在线通道摄像头数量
	OnlineChannelRatio float64 `json:"online_channel_ratio"` //在线通道摄像头占比
}

type QueryResourceDataResp struct {
	DeviceId    string `json:"device_id" example:"1"`
	CpuUsage    string `json:"cpu_percent"`  //CPU利用率
	MemoryUsage string `json:"memory_usage"` //内存使用率
	DiskUsage   string `json:"disk_usage"`   //硬盘使用率 MB
	UpdateTime  string `json:"update_time"`  //更新时间
}

type DeviceInfo struct {
	DeviceName         string `json:"device_name" example:"1"`  //设备名称  全部：默认为空
	DeviceID           string `json:"device_id" example:"1"`    //设备ID、 全部：默认为空
	ProductType        int    `json:"product_type" example:"1"` //产品类型 1:计算盒子；2:布控球；100：全部【默认】
	ProductName        string `json:"product_name"`
	HardwareModel      string `json:"hardware_model" example:"1"`          //硬件型号
	SoftwareName       string `json:"software_name" example:"1"`           //内置软件名称 全部：默认为空
	SystemArchitecture string `json:"system_architecture" example:"arm64"` //架构 arm64、x86 全部：默认为空
	ChipPlatform       string `json:"chip_platform"`                       //芯片平台
	HardwareVersion    string `json:"hardware_version"`                    //硬件版本
	DeviceStatus       int    `json:"device_status" example:"1"`           // '设备状态 1:在线；2:不在线；0：全部【默认】',
	AuthorizationTime  string `json:"authorization_time" example:"1"`      //授权的时间
	AiriaShip          string `json:"airia_ship" example:"1"`              //单兵中airiaShip版本信息
}

type QueryDeviceInfosReq struct {
	UserId          string `json:"user_id" example:"10001"`                        //已授权设备
	PatriarchUserId string `json:"patriarch_user_id" example:"10001"`              //设备源
	DeviceGroupId   string `json:"device_group_id" example:"10001"`                //设备组ID、
	CurrentPageId   int    `json:"current_page_id" binding:"required" example:"1"` //当前页面id
	PageCount       int    `json:"page_count" binding:"required" example:"10"`     //每页条数
	DeviceInfo
}
type AuthorizeDeviceInfoTable struct {
	Id                int       `gorm:"column:id" json:"id"`
	DeviceId          string    `gorm:"column:device_sn" json:"device_id"`
	DeviceName        string    `gorm:"column:name" json:"device_name"`
	DeviceGroupName   string    `gorm:"column:device_group_name"`
	ProductDesign     string    `gorm:"column:product_design" json:"product_name"`
	ProductName       string    `gorm:"column:product_name" json:"product_name2"`
	ProductType       string    `gorm:"column:product_type" json:"hardware_model"`
	XpuType           string    `gorm:"column:xpu_type" json:"chip_platform"`
	XpuVersion        string    `gorm:"column:xpu_version" json:"hardware_version"`
	DeviceStatus      int       `gorm:"column:status" json:"device_status"`
	CreatedTime       time.Time `gorm:"column:created_time"`
	AuthorizationTime string    `json:"authorization_time"`
}

type DeviceModuleDisplay struct {
	DeviceGroupId   string `json:"device_group_id"  example:"11000"`
	DeviceGroupName string `json:"device_group_name"  example:"分组1"`
	ServerCount     int    `json:"server_count"  example:"1"`     //服务器数量
	BoxCount        int    `json:"box_count"  example:"1"`        //盒子数量
	BallCount       int    `json:"ball_count"  example:"1"`       //布控球数量
	AiBoltCount     int    `json:"ai_bolt_count"  example:"1"`    //AI枪机数量
	CameraCount     int    `json:"camera_count"  example:"1"`     //摄像头数量
	PhoneCount      int    `json:"phone_count"  example:"1"`      // 手机数量
	GuardCount      int    `json:"guard_count" example:"1"`       // 手机数量
	CarCameraCount  int    `json:"car_camera_count"  example:"1"` // 手机数量
	CreatedTime     string `json:"created_time"  example:"1"`     //`created_time` datetime DEFAULT NULL COMMENT '创建时间',
}

type DeviceBasicInfo struct {
	ProductType     int    `json:"product_type"  example:"1"`         //'产品类型 1:计算盒子；2:布控球；',
	ProductTypeName string `json:"product_type_name"  example:"计算盒子"` //'产品类型 1:计算盒子；2:布控球；',
	TotalCount      int    `json:"total_count"  example:"100"`        //该类型的产品数量
	OnlineCount     int    `json:"online_count"  example:"10"`        //该类型的产品数量
	Percent         string `json:"percent"  example:"10"`             //百分比
}
type QueryDeviceInfosResp struct {
	DeviceInfos []DeviceInfos
	TotalCount  int `json:"total_count" example:"1244547"` //总记录条数
	PageNum     int `json:"page_num" example:"1244547"`    //总页数
}
type DeviceInfos struct {
	DeviceInfo
	DeviceGroupName string `json:"device_group_name"  example:"11000"`
	RecordDateTime  string `json:"record_datetime"  example:"2021-11-01 13:04:05"` //注册时间
}

type GetChannelConfigResp struct {
	Code   int                  `json:"code"`   //状态码
	Errmsg string               `json:"errmsg"` //错误信息
	Data   GetChannelConfigData `json:"data"`
}

type GetChannelConfigData struct {
	Id                       int         `json:"id"`
	ImageScheduleId          int         `json:"image_schedule_id"`
	VideoScheduleId          int         `json:"video_schedule_id"`
	ImageScheduleName        string      `json:"image_schedule_name"`
	VideoScheduleName        string      `json:"video_schedule_name"`
	NeedCollectImage         int         `json:"need_collect_image"`
	ImageCollectInterval     int         `json:"image_collect_interval"`
	ImageCollectIntervalUnit int         `json:"image_collect_interval_unit"` //image_collect_interval_unit字段：图片采集间隔单位，0为秒，1为分钟，2为小时
	NeedImageSync            int         `json:"need_image_sync"`
	NeedCollectVideo         int         `json:"need_collect_video"`
	VideoCollectDuration     int         `json:"video_collect_duration"`
	VideoCollectInterval     int         `json:"video_collect_interval"`
	VideoCollectIntervalUnit int         `json:"video_collect_interval_unit"`
	NeedVideoSync            int         `json:"need_video_sync"`
	StorageMode              int         `json:"storage_mode"` //storage_mode字段：存储模式，0为循环存储，1为存满停止
	StorageLocation          string      `json:"storage_location"`
	StorageWarnLifetime      int         `json:"storage_warn_lifetime"`
	StorageVideoLifetime     int         `json:"storage_video_lifetime"`
	StorageImageLifetime     int         `json:"storage_image_lifetime"`
	MediaUrl                 string      `json:"media_url"`
	Models                   []GetModels `json:"models"` //models字段：模型信息
}

type GetModels struct {
	Name         string   `json:"name"`
	SdkId        int      `json:"sdk_id"`
	Tags         []string `json:"tags"`
	Interval     int      `json:"interval"`
	IntervalUnit int      `json:"interval_unit"` //models里interval_unit字段：识别间隔单位，0为秒，1为分钟
	Status       int      `json:"status"`
}

type QueryTimeScheduleResp struct {
	Code   int                     `json:"code"`   //状态码
	Errmsg string                  `json:"errmsg"` //错误信息
	Data   []QueryTimeScheduleData `json:"data"`
}

type QueryTimeScheduleData struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Plan Plan   `json:"plan"`
}

type Plan struct {
	Field1 []int `json:"1"`
	Field2 []int `json:"2"`
	Field3 []int `json:"3"`
	Field4 []int `json:"4"`
	Field5 []int `json:"5"`
	Field6 []int `json:"6"`
	Field7 []int `json:"7"`
}

// 获取通道配置信息
type GetChannelEventsResp struct {
	Code   int                    `json:"code"`   //状态码
	Errmsg string                 `json:"errmsg"` //错误信息
	Data   []GetChannelEventsData `json:"data"`
}

type GetChannelEventsData struct {
	Id                   int     `json:"id"`
	ScheduleId           int     `json:"schedule_id"`
	ScheduleName         string  `json:"schedule_name"`
	NeedEvent            int     `json:"need_event"`
	ChannelId            int     `json:"channel_id"`
	Name                 string  `json:"name"`
	BusiName             string  `json:"busi_name"`
	DefaultThres         float64 `json:"default_thres"`
	WarnThres            float64 `json:"warn_thres"`
	WarnLevel            int     `json:"warn_level"`
	WarnAggergatePeriod  int     `json:"warn_aggergate_period"`
	WarnAggergateCounts  int     `json:"warn_aggergate_counts"`
	WarnMediaFormat      int     `json:"warn_media_format"`
	WarnVideoBeforeEvent int     `json:"warn_video_before_event"`
	WarnVideoAfterEvent  int     `json:"warn_video_after_event"`
	NeedSync             int     `json:"need_sync"`
	Snapshot             string  `json:"snapshot"`
	WarnAreaCondition    int     `json:"warn_area_condition"`
	Areas                []Areas `json:"areas"`
}

type Areas struct {
	Name string      `json:"name"`
	Area interface{} `json:"area"`
}

type DeviceResp struct {
	Code   int    `json:"code"`   //状态码
	Errmsg string `json:"errmsg"` //错误信息
}
type DeviceGroupListResp struct {
	DeviceGroupId   string `json:"device_group_id"  example:"11000"`
	DeviceGroupName string `json:"device_group_name"  example:"分组1"`
	Level           string `json:"level"  example:"group"`
}

type QueryDeviceEventsResq struct {
	TotalCount   int           `json:"total_count" example:"1244547"` //总记录条数
	PageNum      int           `json:"page_num" example:"1244547"`    //总页数
	DeviceEvents []DeviceEvent `json:"device_events"`
}
type DeviceEvent struct {
	EventType  string `json:"event_type"`  //事件类型 CPU资源占用
	EventNote  string `json:"event_name"`  // CPU资源占用90%
	CreateTime string `json:"create_time"` //
}

type GuardClockLogInfo struct {
	ActionId int64  `json:"action_id,string,omitempty"`
	Action   string `json:"action,omitempty"`
	AcTime   string `json:"create_time,omitempty"`
}

// GuardActionLogResp 打卡操作记录返回信息
type GuardActionLogResp struct {
	ActionId   int64  `json:"action_id,string"`
	DeviceSn   string `json:"device_sn"`
	DeviceName string `json:"device_name"`
	UserId     string `json:"user_id"`
	UserName   string `json:"user_name"`
	UserNum    string `json:"user_num"`
	RoleName   string `json:"role_name"`
	Action     string `json:"action"`
	AcTime     string `json:"ac_time"`
	NextAction int64  `json:"next_action,string"`
}

// 打卡记录分页返回
type GuardClockInLogPage struct {
	Count int                  `json:"count"`
	Log   []GuardActionLogResp `json:"log"`
}

type QueryChannelsResp struct {
	Code   int               `json:"code"`   //状态码
	Errmsg string            `json:"errmsg"` //错误信息
	Data   QueryChannelsData `json:"data"`
}

type QueryChannelsData struct {
	TotalCount int            `json:"total_count"`
	CurrPageno int            `json:"curr_pageno"`
	Pagesize   int            `json:"pagesize"`
	Orderby    string         `json:"orderby"`
	IsAsc      string         `json:"isAsc"`
	Channels   []ChannelsData `json:"channels"`
}

type ChannelsData struct {
	ChanId     int          `json:"chn_id"`
	Isenable   int          `json:"isenable"` //通道开关 isenable字段：表示是否开启 0 未开启 1开启
	ChanName   string       `json:"chn_name"`
	ModifyTime string       `json:"modify_time"`
	Url        string       `json:"url"` //取流地址
	MediaType  string       `json:"media_type"`
	Transproto string       `json:"transproto"`
	Username   string       `json:"username"`
	Password   string       `json:"password"`
	Location   string       `json:"location"`
	MediaUrl   string       `json:"media_url"` //直播流地址
	State      int          `json:"state"`     //通道状态 state字段：表示通道状态 0 在线显示绿色 1 离线显示红色 2 关闭显示灰色
	IsMod      int          `json:"isMod"`     //isMod字段：表示是否修改 0 未修改（通道树列表不展示该通道行）1 修改过的需要展示
	Models     []ModelsData `json:"models"`
}

type ModelsData struct {
	Name   string `json:"name"`
	SdkId  int    `json:"sdk_id"`
	Status int    `json:"status"`
}

type ChannelGisResp struct {
	DChannel []ChannelResp `json:"d_channels"`
	PChannel []ChannelResp `json:"p_channels"`
}

type ChannelResp struct {
	Id         string `json:"id"` //主键id
	DeviceId   string `json:"device_id" gorm:"column:device_id"`
	ChanId     string `json:"chan_id" gorm:"column:channel_id"`
	ChanName   string `json:"name" gorm:"column:channel_name"`
	ChanStatus int    `json:"chan_status" gorm:"column:channel_status"` //'通道状态',1:在线；2:不在线；',
	Level      string `json:"level" gorm:"column:level"`
	Longitude  string `json:"longitude" gorm:"column:longitude"`
	Latitude   string `json:"latitude" gorm:"column:latitude"`
	Height     string `json:"height" gorm:"column:height"`
	IsAlarm    int    `json:"is_alarm"` //通道是否告过警 1:产生告警;2:未产生告警
}

type IsRegisteredResp struct {
	DeviceId  string `json:"device_id"`
	Base64Str string `json:"base64_str"`
}

type GetPhoneDeviceIdResp struct {
	DeviceId   string `json:"device_id" gorm:"column:device_sn"`
	DeviceName string `json:"device_name" gorm:"column:name"`
}
