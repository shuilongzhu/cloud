package pojo

import (
	"mime/multipart"
)

type AddDeviceGroupReq struct {
	DeviceGroupName string `json:"device_group_name" binding:"required"` //设备组分组名称
	UserGroupId     string `json:"user_group_id" binding:"required"`     //用户分组id
}

type QueryDeviceGroupInfoReq struct {
	DeviceGroupName string `json:"device_group_name"`           //设备组名称、
	CurrentPageId   int    `json:"current_page_id" example:"1"` //当前页面id
	PageCount       int    `json:"page_count" example:"10"`     //每页条数
}

// 添加模型事件告警模板入参
type AddEventTemplateReq struct {
	DeviceId string `json:"device_id" example:"1111111"` //设备ID
	EvName   string `json:"ev_name" example:"抽烟告警模板"`    //模板名称
}

// 设备Id入参
type DeviceIdReq struct {
	DeviceId string `json:"device_id" example:"1111111"` //设备ID
}

// 删除时间计划入参
type DeleteTimeScheduleReq struct {
	DeviceId string `json:"device_id" uri:"device_id" example:"1111111"` //设备ID
	PlanId   string `json:"plan_id" uri:"plan_id" example:"1111111"`     //计划ID
}

// 删除模型入参
type DeleteDeviceModelReq struct {
	DeviceId string `json:"device_id" example:"1111111"` //设备ID
	ModelId  string `json:"model_id" example:"1111111"`  //模型ID
}

// 删除模型事件告警模板入参
type DeleteEventTemplateReq struct {
	DeviceId   string `json:"device_id" example:"1111111"`   //设备ID
	TemplateId string `json:"template_id" example:"1111111"` //模板ID
}

// QueryDeviceTreeReq 查询设备树目录结构入参
type QueryDeviceTreeReq struct {
	DeviceStatus string `form:"device_status" example:"0"`     //设备状态
	ProductType  string `form:"product_type" example:"边缘计算盒子"` //产品外观
}

// GetChannelConfigReq 获取通道配置信息
type GetChannelConfigReq struct {
	DeviceId string `json:"device_id"` //设备ID
	DBpath   string `json:"dbpath"`    //数据库文件的存储路径
	DBname   string `json:"dbname"`
	ChanId   int    `json:"id"` //通道id
}

// 获取通道告警事件设置
type GetChannelEventsReq struct {
	DeviceId string `json:"device_id"` //设备ID
	DBpath   string `json:"dbpath"`    //数据库文件的存储路径
	DBname   string `json:"dbname"`
	ChanId   int    `json:"id"` //通道id
}

type SetChannelEventReq struct {
	DeviceId             string         `json:"device_id"` //设备ID
	DBpath               string         `json:"dbpath"`    //数据库文件的存储路径
	DBname               string         `json:"dbname"`
	Id                   int            `json:"id"` //事件规则id
	Name                 string         `json:"name"`
	ScheduleId           int            `json:"schedule_id"`
	ScheduleName         string         `json:"schedule_name"`
	NeedEvent            int            `json:"need_event"`
	ChnId                int            `json:"chn_id"`
	ChannelId            int            `json:"channel_id"`
	BusiName             string         `json:"busi_name"`
	DefaultThres         float64        `json:"default_thres"`
	WarnThres            float64        `json:"warn_thres"`
	WarnLevel            int            `json:"warn_level"`
	WarnAggergatePeriod  int            `json:"warn_aggergate_period"`
	WarnAggergateCounts  int            `json:"warn_aggergate_counts"`
	WarnMediaFormat      int            `json:"warn_media_format"`
	WarnVideoBeforeEvent int            `json:"warn_video_before_event"`
	WarnVideoAfterEvent  int            `json:"warn_video_after_event"`
	NeedSync             int            `json:"need_sync"`
	WarnAreaCondition    int            `json:"warn_area_condition"`
	Snapshot             string         `json:"snapshot"`
	Areas                []ChannelAreas `json:"areas"`
}

type ChannelAreas struct {
	Area []ChannelEventAreas `json:"area"`
	Name string              `json:"name"`
}

type ChannelEventAreas struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// 设置通道配置信息
type SetChannelConfigReq struct {
	DeviceId                 string `json:"device_id"` //设备ID
	DBpath                   string `json:"dbpath"`    //数据库文件的存储路径
	DBname                   string `json:"dbname"`
	ChanId                   int    `json:"id"` //通道id
	ChnId                    int    `json:"chn_id"`
	ImageScheduleId          int    `json:"image_schedule_id"`
	VideoScheduleId          int    `json:"video_schedule_id"`
	NeedCollectImage         int    `json:"need_collect_image"`
	ImageCollectInterval     int    `json:"image_collect_interval"`
	ImageCollectIntervalUnit int    `json:"image_collect_interval_unit"`
	NeedImageSync            int    `json:"need_image_sync"`
	NeedCollectVideo         int    `json:"need_collect_video"`
	VideoCollectDuration     int    `json:"video_collect_duration"`
	VideoCollectInterval     int    `json:"video_collect_interval"`
	VideoCollectIntervalUnit int    `json:"video_collect_interval_unit"`
	NeedVideoSync            int    `json:"need_video_sync"`
	StorageMode              int    `json:"storage_mode"`
	StorageLocation          string `json:"storage_location"`
	StorageWarnLifetime      int    `json:"storage_warn_lifetime"`
	StorageVideoLifetime     int    `json:"storage_video_lifetime"`
	StorageImageLifetime     int    `json:"storage_image_lifetime"`
	Models                   []struct {
		Name         string   `json:"name"`
		SdkId        int      `json:"sdk_id"`
		Tags         []string `json:"tags"`
		Interval     int      `json:"interval"`
		IntervalUnit int      `json:"interval_unit"`
		Status       int      `json:"status"`
	} `json:"models"`
}

type AddTimeScheduleReq struct {
	DeviceId string `json:"device_id"` //设备ID
	DBpath   string `json:"dbpath"`
	DBname   string `json:"dbname"`
	Name     string `json:"name"`
	Plan     Plan   `json:"plan"`
}

type UpdateTimeScheduleReq struct {
	DeviceId string `json:"device_id"` //设备ID
	DBpath   string `json:"dbpath"`
	DBname   string `json:"dbname"`

	Id   int    `json:"id"`
	Name string `json:"name"`
	Plan Plan   `json:"plan"`
}

type DeleteTimeScheduleSend struct {
	DeviceId string `json:"device_id"` //设备ID
	DBpath   string `json:"dbpath"`
	DBname   string `json:"dbname"`
	PlanId   int    `json:"id"`
}

type DeviceGroupReq struct {
	DeviceGroupName string `json:"device_group_name"` //设备组ID、
	DeviceGroupId   string `json:"device_group_id" binding:"required" example:"10001"`
}

type RegroupReq struct {
	Regroups []Regroup `json:"regroups"`
}

type Regroup struct {
	DeviceId string `json:"device_id" example:"1"` //设备ID
	//DeviceName      string `json:"device_name" example:"1"`           //设备名称
	//DeviceGroupName string `json:"device_group_name" example:"10001"` //设备组名称
	DeviceGroupId string `json:"device_group_id" example:"10001"` //设备组ID、
}

type UpdateDeviceInfoReq struct {
	DeviceId             string `json:"device_id" binding:"required"` //设备ID
	DeviceName           string `json:"device_name"`                  //设备名称
	DeviceGroupId        string `json:"device_group_id"`              //设备组ID、
	DeviceInstallAddress string `json:"device_install_address"`       //设备安装位置
}

type DeviceRegisterReq struct {
	DeviceId   string `json:"device_id" binding:"required" example:"0506060122120001"`                  //设备ID
	DeviceUUId string `json:"device_uuid" binding:"required" example:"b86cac44ed0dbf447d5703a9c1c1f72"` //device_uuid
	DeviceDesc string `json:"device_desc" example:"1"`                                                  //安卓"android" 盒子"emmc" linux是"dmi"
}

type IsRegisteredReq struct {
	DeviceUUId string `json:"device_uuid" binding:"required" example:"b86cac44ed0dbf447d5703a9c1c1f72"` //设备uuid
}

type ReceiveDBFileReq struct {
	DBFile     string `json:"db_file" binding:"required"`   //
	DBFileMd5  string `json:"md5"`                          //
	DeviceId   string `json:"device_id" binding:"required"` //设备ID
	UpdateType int    `json:"type"`                         //自动上传、手动上传
	ModifyTime int    `json:"modify_time"`
}

type MqttResultResp struct {
	Code            int    `json:"code"  example:"200"`
	GenerateRandKey string `json:"generate_rand_key"`
	Message         string `json:"message" example:"响应成功信息"`
}

type ReceiveDeviceEventReq struct {
	DeviceId   string `json:"device_id"`   //设备ID
	EventType  string `json:"event_type"`  //事件类型 CPU资源占用
	EventNote  string `json:"event_name"`  // CPU资源占用90%
	CreateTime int    `json:"create_time"` //
}

// OpenChannelReq 请求开启通道
type OpenChannelReq struct {
	DeviceId string `json:"device_id"` //设备ID
	DBpath   string `json:"dbpath"`    //数据库文件的存储路径
	DBname   string `json:"dbname"`    //数据库文件的名称
	ChanId   int    `json:"chn_id"`
	IsEnable int    `json:"isenable"` //是否启用
}

type OpenChannelResp struct {
	Code   int             `json:"code"`   //状态码
	Errmsg string          `json:"errmsg"` //错误信息
	Data   OpenChannelData `json:"data"`
}

type OpenChannelData struct {
	ChanId   int `json:"chn_id"`
	Isenable int `json:"isenable"`
	State    int `json:"state"`
}

// 请求编辑通道配置信息
type UpdateChannelInfoReq struct {
	DeviceId   string `json:"device_id"` //设备ID
	DBpath     string `json:"dbpath"`    //数据库文件的存储路径
	DBname     string `json:"dbname"`    //数据库文件的名称
	ChanId     int    `json:"chn_id"`
	ChanName   string `json:"chn_name"`
	MediaType  string `json:"media_type"`
	Transproto string `json:"transproto"`
	Url        string `json:"url"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Location   string `json:"location"`
	Models     []struct {
		Name   string `json:"name"`
		SdkId  int    `json:"sdk_id"`
		Status int    `json:"status"`
	} `json:"models"`
	Isenable int `json:"isenable"` //是否开启   表示是否开启 0 未开启 1开启
}
type UpdateChannelInfoResp struct {
	Code   int                   `json:"code"`   //状态码
	Errmsg string                `json:"errmsg"` //错误信息
	Data   UpdateChannelInfoData `json:"data"`
}
type UpdateChannelInfoData struct {
	ChanId   int `json:"chn_id"`
	Isenable int `json:"isenable"`
	State    int `json:"state"`
}

// 查询设备告警事件列表[分页]
type QueryDeviceEventsReq struct {
	DeviceId      string `json:"device_id" binding:"required"` //设备ID
	EventType     string `json:"event_type"`                   //事件类型 CPU资源占用
	CurrentPageId int    `json:"current_page_id" example:"1"`  //当前页面id
	PageCount     int    `json:"page_count" example:"1"`       //每页条数

	StartTime string `json:"start_time" example:"2021-08-09 24:00:00"` // 默认全部""
	EndTime   string `json:"end_time" example:"2021-08-09 24:00:00"`   // 默认全部""
}

// GuardClockIn 哨兵-扫码上报设备SN
type GuardClockIn struct {
	DeviceId string `json:"device_id"` //设备ID
	Action   string `json:"action"`
}

// GuardClockQuery 设备页面查询打卡记录
type GuardClockQuery struct {
	DeviceId  string `json:"device_id"` //设备ID
	UserId    string `json:"user_id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Time      string `json:"time"`
	Page      int    `json:"page"`
	Size      int    `json:"size"`
}

// GuardClockQueryOnUser 用户页面打卡记录获取
type GuardClockQueryOnUser struct {
	UserId   string `json:"user_id"`
	DeviceId string `json:"device_id"` //设备ID
	//ProductDesign string `json:"product_design"`
	//Hardware string `json:"hardware"`
	Time         string `json:"time"`
	Status       int    `json:"status"`
	DeviceStatus int    `json:"device_status"`
	Page         int    `json:"page"`
	Size         int    `json:"size"`
}

// UploadMp3FileArgs 上传Mp3文件的Args
type UploadMp3FileArgs struct {
	ActionId int64  `json:"id,string"`
	DeviceSn string `json:"device_sn"`
	Url      string `json:"url"`
}

// UploadMp3FileCmd 上传mp3文件的cmd
type UploadMp3FileCmd struct {
	Cmd      string            `json:"cmd"`
	Args     UploadMp3FileArgs `json:"args"`
	DeviceId string            `json:"device_id"`
}

// GuardPtzCmdArgs 哨兵云台操作请求
type GuardPtzCmdArgs struct {
	DeviceSn string `json:"device_sn"`
	IsPhone  int    `json:"is_phone"`
	Type     int    `json:"type"`
	Param2   int    `json:"param2"`
	Stop     int    `json:"stop"`
	ActionId int64  `json:"id,string"`
}

// PtzCmd 哨兵云台操作MQTT下发消息格式
type PtzCmd struct {
	Cmd      string          `json:"cmd"`
	Args     GuardPtzCmdArgs `json:"args"`
	DeviceId string          `json:"device_id"`
}

type PtzResult struct {
	Code   int             `json:"code"`
	ErrMsg string          `json:"errmsg"`
	Data   PtzResponseData `json:"data"`
}

type PtzResponseData struct {
	DeviceSn string `json:"device_sn"`
	ActionId int64  `json:"action_id,string"`
}

type UploadMp3Files struct {
	DeviceId string                `form:"device_id" json:"device_id"`
	File     *multipart.FileHeader `form:"file" json:"file"`
	ActionId int64
}

type GetBatchSetChannelStorageType struct {
	Code   int            `json:"code"`
	Errmsg string         `json:"errmsg"`
	Data   ChannelStorage `json:"data"`
}

// BatchSetChannelStorageType 桥接结构体  设置批量采集存储
type BatchSetChannelStorageType struct {
	GetCollectionSettingAndStroageSetting
	ChannelStorage
}
type ChannelStorage struct {
	StorageMode          int            `json:"storage_mode"`
	StorageWarnLifetime  int            `json:"storage_warn_lifetime"`
	StorageVideoLifetime int            `json:"storage_video_lifetime"`
	StorageImageLifetime int            `json:"storage_image_lifetime"`
	StorageTypes         []StorageTypes `json:"storage_types"`
}
type StorageTypes struct {
	Type            int    `json:"type"`
	StorageLocation string `json:"storage_location"`
}
type BridgingResp struct {
	Code   int    `json:"code"`
	Errmsg string `json:"errmsg"`
}

// GetCollectionSettingAndStroageSetting 获取采集批量设置
type GetCollectionSettingAndStroageSetting struct {
	DeviceId string `json:"device_id"`
	DbPath   string `json:"dbpath"`
	DbName   string `json:"dbname"`
}

// CollectionSettingResp 获取的采集设置
type CollectionSettingResp struct {
	BridgingResp
	Data CollectionSetting `json:"data"`
}
type CollectionSetting struct {
	ImageScheduleId          int `json:"image_schedule_id"`
	VideoScheduleId          int `json:"video_schedule_id"`
	NeedCollectImage         int `json:"need_collect_image"`
	ImageCollectInterval     int `json:"image_collect_interval"`
	ImageCollectIntervalUnit int `json:"image_collect_interval_unit"`
	NeedImageSync            int `json:"need_image_sync"`
	NeedCollectVideo         int `json:"need_collect_video"`
	VideoCollectDuration     int `json:"video_collect_duration"`
	VideoCollectInterval     int `json:"video_collect_interval"`
	VideoCollectIntervalUnit int `json:"video_collect_interval_unit"`
	NeedVideoSync            int `json:"need_video_sync"`
}

// SetCollection 设置采集批量设置
type SetCollection struct {
	GetCollectionSettingAndStroageSetting
	CollectionSetting
}

type SetCollection2 struct {
	DeviceIds []string `json:"device_ids"`
	DbPath    string   `json:"dbpath"`
	DbName    string   `json:"dbname"`
	CollectionSetting
}

// CameraCarOnline 车载相机 心跳
type CameraCarOnline struct {
	DeviceIds []string `json:"device_id"`
	OnLine    []int    `json:"on_line"`
}

// ReceivedCameraAlarmInfos 车载相机 告警图片
type ReceivedCameraAlarmInfos struct {
	DeviceId   string `json:"device_id"`
	CreateTime int64  `json:"create_time"`
	EventName  string `json:"event_name"`
	BusiName   string `json:"busi_name"`
	ImgPath    string `json:"img_path"`
}

// 接受告警图片数据
type QueryAlarmImgInfosReq struct {
	DeviceInstallAddress string `json:"device_install_address" form:"device_install_address"` //  '设备安装区域地址',
	ChanInstallAddress   string `json:"chan_install_address" form:"chan_install_address"`

	MediaFormat int `json:"media_format" form:"media_format"` //区分是图片还是视频 0的话告警视频就是空 1的话就有告警视频

	OriginalImage string  `json:"original_image" form:"original_image"` //原图
	AlarmImage    string  `json:"alarm_image" form:"alarm_image"`       //告警图片
	ResizeRatio   float64 `json:"resizeRatio" form:"resizeRatio"`       // 图像的压缩率
	DeviceGBID    string  `json:"device_gb_id" form:"device_gb_id"`     //设备国标ID
	ChannelGBID   string  `json:"channel_gb_id" form:"channel_gb_id"`   //通道国标ID
	DeviceName    string  `json:"device_name" form:"device_name"`       //设备名称
	DeviceID      string  `json:"device_sn" form:"device_sn"`           //设备ID
	WorkId        int     `json:"work_id" form:"work_id"`               //作业ID

	DeviceId          string `json:"device_id" form:"device_id"`   //设备ID
	ChannelID         int    `json:"channel_id" form:"channel_id"` //通道ID
	ChanName          string `json:"channel_name" form:"channel_name"`
	Timestamp         int64  `json:"timestamp" form:"timestamp"`   //报警时间
	Event             string `json:"event" form:"event"`           // 表示 event是helmet
	EventName         string `json:"event_name" form:"event_name"` // event_name 是未佩戴安全帽这样
	WarnImageLifeTime int    `json:"warnImageLifeTime" form:"warnImageLifeTime"`
	WarnLevel         int    `json:"warn_level" form:"warn_level"`   //告警等级0、1、2——————1、2、3;
	AlarmVideo        string `json:"alarm_video" form:"alarm_video"` //告警图片
	Data              []Data `json:"data" form:"data"`
}
type Data struct {
	X          int           `json:"x"`         // 检测框位置 x
	Y          int           `json:"y"`         // 检测框位置 y
	Width      int           `json:"width"`     // 检测框位置 width
	Height     int           `json:"height"`    // 检测框位置 height
	Score      float64       `json:"score"`     // 检测对象置信度得分
	TypeId     int           `json:"type_id"`   // 检测对象类别值
	TypeName   string        `json:"type_name"` // 检测对象类别名
	TrackId    int           `json:"track_id"`  // 检测对象追踪id编号
	AlgoType   int           `json:"algo_type"` // 算法类型 -1未定义，0 表示检测算法，1 表示分类算法
	Attributes []Attributes  `json:"attributes"`
	Child      []interface{} `json:"child"`
}

type Attributes struct {
	AttributeName string  `json:"attribute_name"` // 检测对象属性名
	AttributeId   int     `json:"attribute_id"`
	Value         int     `json:"value"` // 检测对象属性值
	Score         float64 `json:"score"` // 检测对象置信度得分
}

type ChannelGisReq struct {
	Type    int `json:"type" example:"1"`     //设备ID,0:全部;1:摄像头盒子设备;2:手机单兵设备
	IsAlarm int `json:"is_alarm" example:"1"` //设备是否告过警,0:全部;1:产生告警;2:未产生告警
}

type UpdateChannelGisReq struct {
	DeviceId string `json:"device_id" binding:"required"`
	ChanId   string `json:"chan_id" binding:"required"`
	X        string `json:"x"`
	Y        string `json:"y"`
	Z        string `json:"z"`
}

type AiriaShipVersionReq struct {
	DeviceId    string `json:"device_id" binding:"required"` //设备ID
	ShipVersion string `json:"ship_version"`                 //airiaShip软件版本信息
}
