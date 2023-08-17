package dao

import (
	"cloud/common/db"
	comPojo "cloud/common/pojo"
	"cloud/error_code"
	"cloud/module/device_manage/pojo"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

const (
	//selectDeviceInfoJoins 查询设备信息关联用户授权表
	selectDeviceInfoJoins = "left join user_authorization on device_info.device_sn = user_authorization.device_id left join product_info on device_info.product_type = product_info.product_type"
	// selectPreloadDeviceInfoField 预加载查询设备信息所需特定属性字段
	selectPreloadDeviceInfoField = "device_info.id,device_info.group_id,device_info.device_sn,device_info.status,device_info.location,user_authorization.device_name as name,product_info.product_level AS device_level"
	// whereDeviceInfo 设备树设备信息查询条件
	whereDeviceInfo = "user_authorization.user_id = ? and user_authorization.device_group_id = ? and user_authorization.delete_tag = ? and device_info.delete_tag = ?"
)

// DeviceGroupByUserGroupIdSelect @description: 根据user_group_id查询设备分组
// @parameter userId
// @return code
// @return result
func DeviceGroupByUserGroupIdSelect(userGroupId string) (code int, result []pojo.DeviceGroup) {
	result = make([]pojo.DeviceGroup, 0)

	code = db.MyDb.CommonSelect(&result, comPojo.Field{
		TableName: "device_group",
		Pluck:     "id,device_group_id,device_group_name,'group' as level",
		Where:     "delete_tag = ? and user_group_id = ?"}, 0, userGroupId)
	return
}

// DeviceGroupUngroupSelect @description: 根据user_group_id获取设备未分组id
// @parameter userGroupId
// @return code
// @return deviceGroupId
func DeviceGroupUngroupSelect(userGroupId string) (code int, deviceGroupId string) {
	code = db.MyDb.CommonSelect(&deviceGroupId, comPojo.Field{
		TableName: "device_group",
		Pluck:     "device_group_id",
		Where:     "delete_tag = ? and user_group_id = ? and device_group_name = ?"}, 0, userGroupId, "未分组设备")
	return
}

// DeviceChannelInfoSelect @description: 根据条件查询设备树设备通道具体信息
// @parameter deviceGroupId
// @parameter deviceStatus
// @parameter productDesign
// @parameter userId
// @return int
// @return []pojo.Device
func DeviceChannelInfoSelect(deviceGroupId string, deviceStatus, productDesign, userId string) (int, []pojo.Device) {
	result := make([]pojo.Device, 0)

	operate := db.MyDb.Db.Joins(selectDeviceInfoJoins).Select(selectPreloadDeviceInfoField)
	operate.Where(whereDeviceInfo, userId, deviceGroupId, 0, 0)
	if 0 != len(deviceStatus) {
		operate = operate.Where("device_info.status = ?", deviceStatus)
	}
	if 0 != len(productDesign) {
		operate = operate.Where("product_info.product_design = ?", productDesign)
	}
	operate.Preload("Channels", func(query *gorm.DB) *gorm.DB {
		return query.Select("id,device_id,channel_id,channel_name,channel_status,'channel' as level").
			Where("is_false != ?", 1).
			Where("delete_tag = ?", 0)
	}).Find(&result)
	if nil != operate.Error {
		log.Errorln("SelectDeviceChannelInfo method error:", operate.Error)
		return error_code.ErrDeviceTreeSelect, result
	}
	return error_code.Success, result
}

// DeviceGroupIsExist @description: 是否存在此设备分组
// @parameter deviceGroupName
// @parameter userGroupId
// @return code
// @return state
func DeviceGroupIsExist(deviceGroupName, userGroupId string) (code int, state int) {
	code = db.MyDb.CommonSelect(&state, comPojo.Field{
		TableName: "device_group",
		Pluck:     "1",
		Limit:     1,
		Where:     "device_group_name= ? and user_group_id= ? and delete_tag = ?"}, deviceGroupName, userGroupId, 0)
	return
}

// DeviceGroupNameIsExist @description: 是否存在此设备分组名称
// @parameter deviceGroupName
// @parameter userGroupId
// @return code
// @return state
func DeviceGroupNameIsExist(deviceGroupName string) (code int, state int) {
	code = db.MyDb.CommonSelect(&state, comPojo.Field{
		TableName: "device_group",
		Pluck:     "1",
		Limit:     1,
		Where:     "device_group_name= ? and delete_tag = ?"}, deviceGroupName, 0)
	return
}

// UserGroupIsExist @description: 是否存在此用户分组
// @parameter userGroupId
// @return code
// @return state
func UserGroupIsExist(userGroupId string) (code int, state int) {
	code = db.MyDb.CommonSelect(&state, comPojo.Field{
		TableName: "user_group",
		Pluck:     "1",
		Limit:     1,
		Where:     "user_group_id= ? and delete_tag = ?"}, userGroupId, 0)
	return
}

// DeviceGroupUpdate @description: 修改设备分组名称
// @parameter deviceGroupName
// @parameter deviceGroupId
// @parameter userGroupId
// @return code
func DeviceGroupUpdate(deviceGroupName, deviceGroupId, userGroupId string) (code int) {
	code = db.MyDb.CommonUpdate(comPojo.Field{
		TableName: "device_group",
		Where:     "device_group_id= ? and user_group_id = ? and delete_tag = ?"},
		map[string]interface{}{"device_group_name": deviceGroupName, "update_time": time.Now()}, deviceGroupId, userGroupId, 0)
	return
}

// DeviceGroupDelete @description: 删除用户分组id
// @parameter deviceGroupId
// @parameter userGroupId
// @return code
func DeviceGroupDelete(deviceGroupId, userGroupId string) (code int) {
	code = db.MyDb.CommonDelete(comPojo.Field{
		TableName: "device_group",
		Where:     "device_group_id= ? and user_group_id = ? and delete_tag = ?"}, deviceGroupId, userGroupId, 0)
	return
}

// UserAuthorizationUpdate @description: 变更用户设备授权分组
// @parameter old
// @parameter new
// @parameter userId
// @parameter userGroupId
// @return code
func UserAuthorizationUpdate(old, new, userId, userGroupId string) (code int) {
	code = db.MyDb.CommonUpdate(comPojo.Field{
		TableName: "user_authorization",
		Where:     "device_group_id= ? and delete_tag = ? and user_id = ? and user_group_id = ?"},
		map[string]interface{}{"device_group_id": new, "update_time": time.Now()}, old, 0, userId, userGroupId)
	return
}

// UserAuthorizationUpdate2 @description: 更新设备授权表 设备名称，设备分组id
// @parameter req
// @parameter userId
// @return code
func UserAuthorizationUpdate2(req pojo.UpdateDeviceInfoReq, userId, userGroupId string) (code int) {

	//查询设备授权信息
	code, userAuthorization := db.UserAuthorizationSelect(req.DeviceId, userId)
	if error_code.Success != code {
		return
	}

	mapInfo := map[string]interface{}{}

	if "" != req.DeviceName && userAuthorization.DeviceName != req.DeviceName {
		mapInfo["device_name"] = req.DeviceName
	}

	if "" != req.DeviceGroupId && userAuthorization.DeviceGroupId != req.DeviceGroupId {
		state := 0

		code = db.MyDb.CommonSelect(&state, comPojo.Field{
			TableName: "device_group",
			Pluck:     "1",
			Limit:     1,
			Where:     "device_group_id = ? and user_group_id= ? and delete_tag = ?"}, req.DeviceGroupId, userGroupId, 0)
		if error_code.Success != code {
			return
		}

		if 0 == state {
			return error_code.ErrIsNotExistDeviceGroup
		}

		mapInfo["device_group_id"] = req.DeviceGroupId
	}

	if 0 != len(mapInfo) {
		mapInfo["update_time"] = time.Now()
		code = db.MyDb.CommonUpdate(comPojo.Field{
			TableName: "user_authorization",
			Where:     "device_id = ? and user_id = ? and delete_tag = ?"}, mapInfo, req.DeviceId, userId, 0)
	}

	return
}

// DeviceGroupSelect @description: 分页查询设备分组信息
// @parameter userGroupId
// @return code
// @return result
func DeviceGroupSelect(req pojo.QueryDeviceGroupInfoReq, userGroupId string) (code int, result []pojo.DeviceGroupInfoResp) {
	result = make([]pojo.DeviceGroupInfoResp, 0)

	where := fmt.Sprintf("user_group_id = %s and delete_tag = %d", userGroupId, 0)
	if "" != req.DeviceGroupName {
		where = where + " and device_group_name like %" + req.DeviceGroupName + "%"
	}
	code = db.MyDb.CommonSelect(&result, comPojo.Field{
		TableName: "device_group",
		Pluck:     "device_group_id,device_group_name,created_time",
		Sort:      "created_time asc",
		Offset:    (req.CurrentPageId - 1) * req.PageCount,
		Limit:     req.PageCount,
		Where:     where})
	return
}

// DeviceGroupCountSelect @description: 查询设备分组总数
// @parameter req
// @parameter userGroupId
// @return code
// @return result
func DeviceGroupCountSelect(req pojo.QueryDeviceGroupInfoReq, userGroupId string) (code, result int) {

	where := fmt.Sprintf("user_group_id = %s and delete_tag = %d", userGroupId, 0)
	if "" != req.DeviceGroupName {
		where = where + " and device_group_name like %" + req.DeviceGroupName + "%"
	}
	code = db.MyDb.CommonSelect(&result, comPojo.Field{
		TableName: "device_group",
		Pluck:     "count(*)",
		Where:     where})
	return
}

// DeviceInfoSelect @description: 获取设备信息
// @parameter userId
// @parameter deviceId
// @return code
// @return result
func DeviceInfoSelect(userId, deviceId string) (code int, result comPojo.DeviceInfo) {
	result = comPojo.DeviceInfo{}

	code = db.MyDb.CommonSelect(&result, comPojo.Field{
		TableName: "device_info di",
		Pluck:     "di.device_sn,ua.device_name as name,di.status,di.product_type,dg.device_group_name as device_group_name,di.location,di.product_img_url,DATE_FORMAT(di.created_time,'%Y-%m-%d %H:%i:%s') as created_time,DATE_FORMAT(di.update_time,'%Y-%m-%d %H:%i:%s') as update_time,di.airia_ship",
		Joins:     "left join user_authorization ua on di.device_sn = ua.device_id left join device_group dg on ua.device_group_id = dg.device_group_id",
		Where:     "ua.device_id = ? and ua.user_id = ? and ua.delete_tag = ?"}, deviceId, userId, 0)
	return
}

// DeviceStateSelect @description: 获取设备状态分组信息
// @parameter deviceIds
// @return code
// @return result
func DeviceStateSelect(deviceIds []string) (code int, result []pojo.DeviceState) {
	code = db.MyDb.CommonSelect(&result, comPojo.Field{
		TableName: "device_info",
		Pluck:     "count(*) as count,status",
		Group:     "status",
		Where:     "device_sn in ? and delete_tag = ?"}, deviceIds, 0)
	return
}

// DeviceInfoUpdate @description: 更新设备表信息
// @parameter req
// @return code
func DeviceInfoUpdate(req pojo.UpdateDeviceInfoReq) (code int) {

	mapInfo := map[string]interface{}{}

	if "" != req.DeviceName {
		mapInfo["name"] = req.DeviceName
	}

	if "" != req.DeviceInstallAddress {
		mapInfo["location"] = req.DeviceInstallAddress
	}

	if 0 != len(mapInfo) {
		mapInfo["update_time"] = time.Now()
		code = db.MyDb.CommonUpdate(comPojo.Field{
			TableName: "device_info",
			Where:     "device_sn = ? and delete_tag = ?"}, mapInfo, req.DeviceId, 0)
	}
	return
}

// ChannelCountSelect @description: 查询用户授权设备的通道数量
// @parameter deviceIds
// @return code
// @return count
func ChannelCountSelect(deviceIds []string) (code, count int) {
	code = db.MyDb.CommonSelect(&count, comPojo.Field{
		TableName: "channel_info",
		Pluck:     "count(*)",
		Where:     "device_id in ? and delete_tag = ?"}, deviceIds, 0)
	return
}

// ChannelStateSelect @description: 获取通道状态分组信息
// @parameter deviceIds
// @return code
// @return result
func ChannelStateSelect(deviceIds []string) (code int, result []pojo.ChannelState) {
	code = db.MyDb.CommonSelect(&result, comPojo.Field{
		TableName: "channel_info",
		Pluck:     "count(*) as count,channel_status",
		Group:     "channel_status",
		Where:     "device_id in ? and delete_tag = ?"}, deviceIds, 0)
	return
}

// ChannelInfoSelect @description: 获取通道基本信息
// @parameter deviceId
// @parameter channelId
// @return code
// @return result
func ChannelInfoSelect(deviceId, channelId string) (code int, result comPojo.ChannelInfo) {
	result = comPojo.ChannelInfo{}

	code = db.MyDb.CommonSelect(&result, comPojo.Field{
		TableName: "channel_info",
		Pluck:     "device_id,device_name,channel_id,channel_name,channel_ip,channel_username,channel_access_type,take_way,channel_address,channel_url,channel_model,setup_time,channel_status,take_address,DATE_FORMAT(access_time,'%Y-%m-%d %H:%i:%s') as access_time,DATE_FORMAT(created_time,'%Y-%m-%d %H:%i:%s') as created_time",
		Where:     "device_id = ? and channel_id = ? and delete_tag = ?"}, deviceId, channelId, 0)
	return
}

func DeviceChannelCountSelect(deviceId, channelId, channelStatus string) (code, result int) {
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

	code = db.MyDb.CommonSelect(&result, comPojo.Field{
		TableName: "channel_info",
		Pluck:     "count(*)",
		Where:     where})

	return
}
