package pojo

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	UserId      string `json:"user_id"`
	UserGroupId string `json:"user_group_id"`
	UserName    string `json:"user_name"`
	RoleName    string `json:"role_name"`
	RoleId      int    `json:"role_id"`
	UserNum     string `json:"user_num"`
	jwt.StandardClaims
}

type Field struct {
	TableName string `json:"table_name"` //表名称
	Pluck     string `json:"pluck"`      //要查询的列，*：查询全部字段
	Joins     string `json:"joins"`      //关联的表和关联条件
	Group     string `json:"group"`      //分组
	Sort      string `json:"sort"`       //排序,例:id desc
	Offset    int    `json:"offset"`     //分页从第几个开始查
	Limit     int    `json:"limit"`      //分页一页查几条
	Where     string `json:"where"`      //where条件 user_name = ? and password = ?
}

// DeviceGroup 设备分组表 device_group
type DeviceGroup struct {
	Id              int64     `gorm:"primary_key; column:id"`   //`id` int NOT NULL AUTO_INCREMENT  '自增ID',
	DeviceGroupId   string    `gorm:"column:device_group_id"`   //`device_group_id` varchar(32)设备分组ID',
	UserGroupId     string    `gorm:"column:user_group_id"`     //`user_group_id` varchar(32)用户分组id',
	DeviceGroupName string    `gorm:"column:device_group_name"` //`device_group_name` varchar(32)设备分组名称',
	DeviceGroupNote string    `gorm:"column:device_group_note"` //`device_group_note` varchar(128)设备分组描述',
	CreatedTime     time.Time `gorm:"column:created_time"`      //`created_time` datetime创建时间',
	UpdateTime      time.Time `gorm:"column:update_time"`       //`update_time` datetime更新时间',
	DeleteTag       int       `gorm:"column:delete_tag"`        //`delete_tag` int DEFAULT '0'  '删除 0:正常；1：已删除；',
}

// DeviceInfo 设备信息表 device_info
type DeviceInfo struct {
	Id                    int64  `json:"id" gorm:"AUTO_INCREMENT; primary_key; column:id"`  //id
	DeviceSn              string `json:"device_id" gorm:"column:device_sn"`                 //设备ID
	DeviceName            string `json:"name" gorm:"column:name"`                           //设备名称
	DeviceGbId            string `json:"gb_id" gorm:"column:gb_id"`                         //设备国标ID
	DeviceGroupId         string `json:"group_id" gorm:"column:group_id"`                   //设备分组ID
	UserGroupId           string `json:"user_group_id" gorm:"column:user_group_id"`         //用户分组id
	DeviceGroupName       string `json:"device_group_name" gorm:"column:device_group_name"` //设备分组名称
	ProductType           string `json:"product_type" gorm:"column:product_type"`           //产品类型 1:计算盒子；2:布控球；3:AI枪机 4:普通摄像头[手机] 5：服务器',6抓拍机
	ProductImgUrl         string `json:"product_img_url" gorm:"column:product_img_url"`     //产品图片存储地址
	DeviceStatus          int    `json:"status" gorm:"column:status"`                       //设备状态 1:在线；2:不在线；3:未激活；4暂停服务
	DeviceInstallPosition string `json:"location" gorm:"column:location"`                   //设备安装位置
	CreatedTime           string `json:"created_time" gorm:"column:created_time"`           //创建时间
	UpdateTime            string `json:"update_time" gorm:"column:update_time"`             //更新时间
	DeleteTag             int    `json:"delete_tag" gorm:"column:delete_tag"`               //删除 0:正常；1:已删除
	AlternateField1       string `json:"alternate_field1" gorm:"column:alternate_field1"`   //用字段1,2:已激活；3:未激活；4暂停服务
	DeviceType            int    `json:"type" gorm:"column:type"`                           //设备类型 1 信通网关 2艾瑞抓拍机
	Uuid                  string `json:"uuid" gorm:"column:uuid"`                           //设备UUID
	IsOpen                int    `json:"is_open" gorm:"column:is_open"`                     //是否开启(1开启；2关闭)
	Licence               string `json:"licence" gorm:"column:licence"`                     //许可证
	AiriaShip             string `json:"airia_ship" gorm:"column:airia_ship"`               //ship版本信息
}

type ChannelInfo struct {
	Id                int64     `json:"id" gorm:"column:id"`
	DeviceId          string    `json:"device_id" gorm:"column:device_id"`
	DeviceName        string    `json:"device_name" gorm:"column:device_name"`
	ChannelId         string    `json:"channel_id" gorm:"column:channel_id"`
	ChannelName       string    `json:"channel_name" gorm:"column:channel_name"`
	ChannelIp         string    `json:"channel_ip" gorm:"column:channel_ip"`
	ChannelUsername   string    `json:"channel_username" gorm:"column:channel_username"`
	ChannelPassword   string    `json:"channel_password" gorm:"column:channel_password"`
	ChannelAccessType string    `json:"channel_access_type" gorm:"column:channel_access_type"`
	TakeWay           string    `json:"take_way" gorm:"column:take_way"`
	ChannelAddress    string    `json:"channel_address" gorm:"column:channel_address"`
	ChannelUrl        string    `json:"channel_url" gorm:"column:channel_url"`
	ChannelModel      string    `json:"channel_model" gorm:"column:channel_model"`
	SetupTime         string    `json:"setup_time" gorm:"column:setup_time"`
	ChannelStatus     string    `json:"channel_status" gorm:"column:channel_status"`
	AccessTime        string    `json:"access_time" gorm:"column:access_time"`
	CreatedTime       string    `json:"created_time" gorm:"column:created_time"`
	DeleteTime        time.Time `json:"delete_time" gorm:"column:delete_time"`
	DeleteTag         int       `json:"delete_tag" gorm:"column:delete_tag"`
	TakeAddress       string    `json:"take_address" gorm:"column:take_address"`
	IsFalse           int       `json:"is_false" gorm:"column:is_false"`
	AlternateField2   string    `json:"alternate_field2" gorm:"column:alternate_field2"`
	Longitude         string    `json:"longitude" gorm:"column:longitude"`
	Latitude          string    `json:"latitude" gorm:"column:latitude"`
	Height            string    `json:"height" gorm:"column:height"`
}

// UserInfo 用户信息表 user_info
type UserInfo struct {
	Id                       int       `gorm:"AUTO_INCREMENT; primary_key; column:id"` //`id` int NOT NULL AUTO_INCREMENT  '自增ID',
	UserId                   string    `gorm:"column:user_id"`                         //`user_id` varchar(32) NOT NULL  '用户id',
	UserGroupId              string    `gorm:"column:user_group_id"`                   //`user_group_id` varchar(32) NOT NULL  '用户分组id',
	UserGroupName            string    `gorm:"column:user_group_name"`                 //`user_group_name` varchar(32)用户组名称',
	UserName                 string    `gorm:"column:user_name"`                       //`user_name` varchar(32)用户姓名',
	Nickname                 string    `gorm:"column:nickname"`                        //`nickname` varchar(32)用户昵称',
	Role                     int       `gorm:"column:role"`                            //`role` int角色类型 0：普通用户；1：系统超级管理员；2:运维观察员；3:模型管理员',[不能更新为普通用户 4为普通用户]
	ParentNodeId             string    `gorm:"column:parent_node_id"`                  //`parent_node_id` varchar(32)父节点ID',
	Number                   string    `gorm:"column:number"`                          //`number` varchar(32)手机号',
	UserPassword             string    `gorm:"column:user_password"`                   //`user_password` varchar(32)用户密码 md5加密',
	Email                    string    `gorm:"column:emil"`                            //`emil` varchar(32)邮箱',
	Note                     string    `gorm:"column:note"`                            //`note` varchar(512)备注信息',
	CreatedTime              time.Time `gorm:"column:created_time"`                    //`created_time` datetime创建时间',
	UpdateTime               time.Time `gorm:"column:update_time"`                     //`update_time` datetime更新时间',
	LoginTime                time.Time `gorm:"column:login_time"`                      //`login_time` datetime登录时间',
	LoginIp                  string    `gorm:"column:login_ip"`                        //`login_ip` varchar(32)登录IP',
	AvatarStorageAddress     string    `gorm:"column:avatar_storage_address"`          //`avatar_storage_address` varchar(128)头像存储地址',
	IdPhotoAddress           string    `gorm:"column:Id_photo_address"`                //`Id_photo_address` varchar(128)身份证照片存储地址',
	IdNumber                 string    `gorm:"column:Id_number"`                       //`Id_number` varchar(32)身份证号码',
	IdCardVerificationStatus int       `gorm:"column:Id_card_verification_status"`     //`Id_card_verification_status` int身份证验证状态 0：待提交；1：待审核；2：审核通过；3：审核驳回，需修改信息；4：审核拒绝;',
	UserState                int       `gorm:"column:user_state"`                      //`user_state` int用户状态 1：正常；2：已禁用；',
	WeChatID                 string    `gorm:"column:WeChat_ID"`                       //`WeChat_ID` varchar(32)微信ID',
	QQId                     string    `gorm:"column:qq_id"`                           //`qq_id` varchar(32)QQ ID',
	AlternateField1          string    `gorm:"column:alternate_field1"`                //`alternate_field1` varchar(32)备用字段1',
	AlternateField2          string    `gorm:"column:alternate_field2"`                //`alternate_field2` varchar(32)备用字段2',
	AlternateField3          string    `gorm:"column:alternate_field3"`                //`alternate_field3` varchar(32)备用字段3',
}

// UserGroup 用户组表 user_group
type UserGroup struct {
	Id              int       `gorm:"AUTO_INCREMENT; primary_key; column:id"` //`id` int NOT NULL AUTO_INCREMENT  '自增ID',
	UserGroupId     string    `gorm:"column:user_group_id"`                   //`user_group_id` varchar(32)用户组id',
	UserGroupName   string    `gorm:"column:user_group_name"`                 //`user_group_name` varchar(32)用户组名称',
	UserGroupNote   string    `gorm:"column:user_group_note"`                 //`user_group_note` varchar(128)用户组描述',
	CreatedTime     time.Time `gorm:"column:created_time"`                    //`created_time` datetime创建时间',
	UpdateTime      time.Time `gorm:"column:update_time"`                     //`update_time` datetime更新时间',
	DeleteTag       int       `gorm:"column:delete_tag"`                      //`delete_tag` int删除标记 0:正常；1删除',
	AlternateField1 string    `gorm:"column:alternate_field1"`                //`alternate_field1` varchar(32)备用字段1',
	AlternateField2 string    `gorm:"column:alternate_field2"`                //`alternate_field2` varchar(32)备用字段2',
}

// UserAuthorization 用户授权记录表 user_authorization
type UserAuthorization struct {
	Id            int       `json:"id" gorm:"AUTO_INCREMENT; primary_key; column:id"` //`id` int NOT NULL AUTO_INCREMENT  '自增ID',
	DeviceId      string    `json:"device_id" gorm:"column:device_id"`                //`device_id` varchar(32)设备ID',
	UserId        string    `json:"user_id" gorm:"column:user_id"`                    //`user_id` varchar(32)用户id',
	DeviceName    string    `json:"device_name" gorm:"column:device_name"`
	DeviceGroupId string    `json:"device_group_id" gorm:"column:device_group_id"`
	UserGroupId   string    `json:"user_group_id" gorm:"column:user_group_id"`
	CreatedTime   time.Time `json:"created_time" gorm:"column:created_time"` //`created_time` datetime创建授权时间',
	DeleteTime    time.Time `json:"delete_time" gorm:"column:delete_time"`   //`delete_time` datetime删除授权时间',
	UpdateTime    time.Time `json:"update_time" gorm:"column:update_time"`   //`update_time` datetime更新授权时间',
	DeleteTag     int       `json:"delete_tag" gorm:"column:delete_tag"`     //`delete_tag` int删除标记 0:正常；1：已删除；',
}

// UserLoginInfo 用户登录记录表 user_login_info
type UserLoginInfo struct {
	Id          int       `gorm:"AUTO_INCREMENT; primary_key; column:id"` //`id` int NOT NULL AUTO_INCREMENT  '自增ID',
	UserId      string    `gorm:"column:user_id"`                         //`user_id` varchar(32)用户id',
	UserGroupId string    `gorm:"column:user_group_id"`                   //`user_group_id` varchar(32)用户分组id',
	LoginTime   time.Time `gorm:"column:login_time"`                      //`login_time` datetime登录时间',
	LoginDevice int       `gorm:"column:login_device"`                    //`login_device` int登录设备 1:pc端；2:android；3:iOS;',
	LoginIp     string    `gorm:"column:login_ip"`                        //`login_ip` varchar(32)登录IP',
	CreatedTime time.Time `gorm:"column:created_time"`                    //`created_time` datetime创建时间',
}

type RequestRecord struct {
	Id            int64     `json:"id" gorm:"primary_key; column:id"`
	ReqMethod     string    `json:"req_method" gorm:"column:req_method"`
	ReqUrl        string    `json:"req_url" gorm:"column:req_url"`
	ReqParam      string    `json:"req_param" gorm:"column:req_param"`
	Code          int       `json:"code" gorm:"column:code"`
	ClientIp      string    `json:"client_ip" gorm:"column:client_ip"`
	UserName      string    `json:"user_name" gorm:"column:user_name"`
	StartTime     time.Time `json:"start_time" gorm:"column:start_time"`
	ExecutionTime int64     `json:"execution_time" gorm:"column:execution_time"`
	DeleteTag     int       `json:"delete_tag" gorm:"column:delete_tag"` //删除标记 0:正常；1：已删除
	CreateTime    time.Time `json:"create_time" gorm:"column:create_time"`
}
