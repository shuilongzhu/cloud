package db

import (
	"cloud/common/pojo"
	"cloud/error_code"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

func DataInit() {
	DeviceChannelIdInit()
	DeviceOnlineEnumInit()
	ChannelOnlineEnumInit()
}

// DeviceOnlineEnumInit @description: 初始化在线设备心跳
func DeviceOnlineEnumInit() {
	deviceIds := make([]string, 0)

	code := MyDb.CommonSelect(&deviceIds, pojo.Field{
		TableName: "device_info",
		Pluck:     "device_sn",
		Where:     "status = ? and delete_tag = ?"}, 1, 0)
	if error_code.Success != code {
		log.Errorf("code:%v,DeviceOnlineEnum init_data fail", code)
		return
	}

	var timestamp = time.Now().Unix()
	for _, deviceId := range deviceIds {
		pojo.DeviceOnlineEnum.Store(deviceId, pojo.DeviceResourceReq{TimeStampC: timestamp})
	}
}

// ChannelOnlineEnumInit @description: 初始化在线通道
func ChannelOnlineEnumInit() {
	channelInfos := make([]pojo.ChannelInfo, 0)

	code := MyDb.CommonSelect(&channelInfos, pojo.Field{
		TableName: "channel_info",
		Pluck:     "device_id,channel_id",
		Where:     "channel_status = ? and delete_tag = ?"}, 1, 0)
	if error_code.Success != code {
		log.Errorf("code:%v,ChannelOnlineEnum init_data fail", code)
		return
	}

	var timestamp = time.Now().Unix()
	for _, channelInfo := range channelInfos {
		pojo.SetCOE(channelInfo.DeviceId, channelInfo.ChannelId, timestamp)
	}
}

func DeviceChannelIdInit() {
	deviceInfos := make([]pojo.DeviceInfo, 0)
	channelInfos := make([]pojo.ChannelInfo, 0)

	code := MyDb.CommonSelect(&deviceInfos, pojo.Field{
		TableName: "device_info",
		Pluck:     "id,device_sn",
		Where:     "delete_tag = ?"}, 0)
	if error_code.Success != code {
		log.Errorf("code:%v,DeviceChannelIdInit fail", code)
		return
	}

	for _, deviceInfo := range deviceInfos {
		pojo.DeviceIdEnum[deviceInfo.DeviceSn] = deviceInfo.Id
	}

	code = MyDb.CommonSelect(&channelInfos, pojo.Field{
		TableName: "channel_info",
		Pluck:     "id,device_id,channel_id",
		Where:     "delete_tag = ?"}, 0)
	if error_code.Success != code {
		log.Errorf("code:%v,DeviceChannelIdInit fail", code)
		return
	}

	for _, channelInfo := range channelInfos {
		pojo.ChannelIdEnum[channelInfo.DeviceId+channelInfo.ChannelId] = channelInfo.Id
	}
}

// UserGroupSelect @description:根据user_group_id查询用户组
// @return code
// @return result
func UserGroupSelect(userGroupId string) (code int, result pojo.UserGroup) {
	code = MyDb.CommonSelect(&result, pojo.Field{TableName: "user_group", Pluck: "id", Limit: 1, Where: "user_group_id = ?"}, userGroupId)
	return
}

// UserGroupCreate @description: 创建用户组
// @parameter userGroup
// @return code
func UserGroupCreate(userGroup pojo.UserGroup) (code int) {
	code = MyDb.CommonCreate("user_group", &userGroup)
	return
}

// UserInfoSelect @description: 根据user_name查询用户信息
// @parameter userName
// @return code
// @return result
func UserInfoSelect(userName string) (code int, result pojo.UserInfo) {
	code = MyDb.CommonSelect(&result, pojo.Field{TableName: "user_info", Pluck: "id,user_id,user_group_id,user_password,role,avatar_storage_address,number,emil", Limit: 1, Where: "user_name= ?"}, userName)
	return
}

// UserInfoCreate @description: 创建用户
// @parameter userInfo
// @return code
func UserInfoCreate(userInfo pojo.UserInfo) (code int) {
	code = MyDb.CommonCreate("user_info", &userInfo)
	return
}

// DeviceGroupSelect @description: 根据device_group_id查询设备分组信息
// @parameter deviceGroupId
// @return code
// @return result
func DeviceGroupSelect(deviceGroupId string) (code int, result pojo.DeviceGroup) {
	result = pojo.DeviceGroup{}
	code = MyDb.CommonSelect(&result, pojo.Field{
		TableName: "device_group",
		Pluck:     "id,device_group_name",
		Limit:     1,
		Where:     "device_group_id= ? and delete_tag = ?"}, deviceGroupId, 0)
	return
}

// DeviceGroupCreate @description: 创建设备分组
// @parameter deviceGroup
// @return code
func DeviceGroupCreate(deviceGroup pojo.DeviceGroup) (code int) {
	code = MyDb.CommonCreate("device_group", &deviceGroup)
	return
}

// RequestRecordCreate @description:新增项目请求记录
// @parameter requestRecord
// @return code
func RequestRecordCreate(requestRecord pojo.RequestRecord) (code int) {
	code = MyDb.CommonCreate("request_record", &requestRecord)
	return
}

// AuthDeviceIdSelect @description: 根据入参条件查询此用户下的授权设备id
// @parameter userId
// @parameter deviceGroupId
// @parameter deviceId
// @return int
// @return []string
func AuthDeviceIdSelect(userId, deviceGroupId, deviceId string) (code int, result []string) {
	result = make([]string, 0)
	where := fmt.Sprintf("device_info.delete_tag = %d and user_authorization.delete_tag = %d", 0, 0)

	if "" != userId {
		where = fmt.Sprintf("%s and user_authorization.user_id = %s", where, userId)
	}
	if "" != deviceGroupId {
		where = fmt.Sprintf("%s and user_authorization.device_group_id = %s", where, deviceGroupId)
	}
	if "" != deviceId {
		where = fmt.Sprintf("%s and user_authorization.device_id = %s", where, deviceId)
	}
	code = MyDb.CommonSelect(&result, pojo.Field{
		TableName: "device_info",
		Pluck:     "distinct device_info.device_sn",
		Joins:     "left join user_authorization on device_info.device_sn=user_authorization.device_id",
		Sort:      "device_info.device_sn desc",
		Where:     where})
	return
}

// DeviceIsAuth @description: 设备是否授权
// @parameter deviceId
// @parameter userId
// @return code
func DeviceIsAuth(deviceId, userId string) (code int) {
	state := 0

	code = MyDb.CommonSelect(&state, pojo.Field{
		TableName: "user_authorization",
		Pluck:     "1",
		Limit:     1,
		Where:     "device_id = ? and user_id = ? and delete_tag = ?"}, deviceId, userId, 0)

	if 0 == state { //未授权
		code = error_code.ErrDeviceNotAuth
	}

	return
}

// ChannelInfoSelect @description: 根据条件获取通道信息
// @parameter deviceId
// @return code
// @return result
func ChannelInfoSelect(deviceId, channelId, channelStatus string, pageId, pageCount int) (code int, result []pojo.ChannelInfo) {
	result = make([]pojo.ChannelInfo, 0)

	where := fmt.Sprintf("delete_tag = %d", 0)

	if "" != deviceId {
		where = fmt.Sprintf("%s and device_id = %s", where, deviceId)
	}
	if "" != channelId {
		where = fmt.Sprintf("%s and channel_id = %s", where, channelId)
	}
	if "" != channelStatus {
		where = fmt.Sprintf("%s and channel_status = %s", where, channelStatus)
	}

	code = MyDb.CommonSelect(&result, pojo.Field{
		TableName: "channel_info",
		Pluck:     "device_id,device_name,channel_id,channel_name,channel_ip,channel_username,channel_access_type,take_way,channel_address,channel_url,channel_model,setup_time,channel_status,take_address,DATE_FORMAT(access_time,'%Y-%m-%d %H:%i:%s') as access_time,DATE_FORMAT(created_time,'%Y-%m-%d %H:%i:%s') as created_time",
		Sort:      "created_time asc",
		Offset:    (pageId - 1) * pageCount,
		Limit:     pageCount,
		Where:     where})

	return
}

// UserAuthorizationSelect @description: 查询用户设备授权信息
// @parameter deviceId
// @parameter userId
// @return code
// @return result
func UserAuthorizationSelect(deviceId, userId string) (code int, result pojo.UserAuthorization) {
	result = pojo.UserAuthorization{}

	code = MyDb.CommonSelect(&result, pojo.Field{
		TableName: "user_authorization",
		Pluck:     "device_name,device_group_id",
		Where:     "device_id = ? and user_id = ? and delete_tag = ?"}, deviceId, userId, 0)
	return
}
