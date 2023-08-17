package service

import (
	commonDb "cloud/common/db"
	comPojo "cloud/common/pojo"
	"cloud/common/utils/data_convert"
	"cloud/error_code"
	"cloud/module/device_manage/dao"
	"cloud/module/device_manage/pojo"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math"
	"strconv"
	"time"
)

// SelectDeviceTree @description: 处理查询设备树
// @parameter dm_types.QueryDeviceTreeReq
// @return int
// @return []dto.DeviceGroup
func SelectDeviceTree(req pojo.QueryDeviceTreeReq, claims comPojo.Claims) (code int, result []pojo.DeviceGroup) {
	//根据user_id查询此用户组下所有的设备组
	code, result = dao.DeviceGroupByUserGroupIdSelect(claims.UserGroupId)
	if error_code.Success != code || 0 == len(result) {
		return
	}

	for i := 0; i < len(result); i++ {
		code, result[i].Devices = dao.DeviceChannelInfoSelect(result[i].DeviceGroupId, req.DeviceStatus, req.ProductType, claims.UserId)
		if error_code.Success != code {
			return code, []pojo.DeviceGroup{}
		}
	}

	return
}

// AddDeviceGroup @description: 添加设备分组
// @parameter req
// @return int
func AddDeviceGroup(req pojo.AddDeviceGroupReq) (code int) {
	//校验用户组id是否存在
	code, state := dao.UserGroupIsExist(req.UserGroupId)
	if error_code.Success != code {
		return
	}

	if 0 == state {
		return error_code.ErrIsNotExistUserGroup
	}

	//查询该用户组下此设备组名称是否已存在
	code, state = dao.DeviceGroupIsExist(req.DeviceGroupName, req.UserGroupId)
	if error_code.Success != code {
		return
	}

	if 1 == state {
		return error_code.ErrIsExistDeviceGroup
	}

	deviceGroup := comPojo.DeviceGroup{
		Id:              data_convert.NextId(),
		DeviceGroupId:   comPojo.IntPrefix + strconv.FormatInt(data_convert.NextId(), 10),
		UserGroupId:     req.UserGroupId,
		DeviceGroupName: req.DeviceGroupName,
		CreatedTime:     time.Now(),
		UpdateTime:      time.Now(),
	}
	//添加设备分组
	code = commonDb.DeviceGroupCreate(deviceGroup)

	return
}

// UpdateDeviceGroup @description: 编辑设备分组名称
// @parameter req
// @return code
func UpdateDeviceGroup(req pojo.DeviceGroupReq, claims comPojo.Claims) int {

	//查询设备组id是否已存在
	code, deviceGroup := commonDb.DeviceGroupSelect(req.DeviceGroupId)
	if error_code.Success != code {
		return code
	}

	if 0 == deviceGroup.Id {
		return error_code.ErrIsNotExistDeviceGroup
	}

	if "未分组设备" == deviceGroup.DeviceGroupName {
		log.Warnln("未分组不允许修改名称")
		return error_code.ErrUngroupNotAllowOperate
	}

	//查询是否存在此设备分组名称
	code, state := dao.DeviceGroupNameIsExist(req.DeviceGroupName)
	if error_code.Success != code {
		return code
	}

	if 1 == state {
		return error_code.ErrIsExistDeviceGroupName
	}

	//更新分组名称
	return dao.DeviceGroupUpdate(req.DeviceGroupName, req.DeviceGroupId, claims.UserGroupId)
}

// DeleteDeviceGroup @description: 删除设备分组
// @parameter req
// @parameter userId
// @parameter userGroupId
// @return int
func DeleteDeviceGroup(req pojo.DeviceGroupReq, claims comPojo.Claims) int {
	code, deviceGroup := commonDb.DeviceGroupSelect(req.DeviceGroupId)
	if error_code.Success != code || 0 == deviceGroup.Id {
		return code
	}

	if "未分组设备" == deviceGroup.DeviceGroupName {
		log.Warnln("该未分组设备不允许删除")
		return error_code.ErrUngroupNotAllowOperate
	}

	//获取该用户组下未分组设备的设备分组id
	code, deviceGroupId := dao.DeviceGroupUngroupSelect(claims.UserGroupId)
	if error_code.Success != code {
		return code
	}

	//把目标分组下的所有设备更新到未分组下
	code = dao.UserAuthorizationUpdate(req.DeviceGroupId, deviceGroupId, claims.UserId, claims.UserGroupId)
	if error_code.Success != code {
		return code
	}

	//删除设备分组
	return dao.DeviceGroupDelete(req.DeviceGroupId, claims.UserGroupId)
}

// QueryDeviceGroupInfo @description: 获取设备组信息列表
// @parameter req
// @parameter userGroupId
// @parameter userId
// @return code
// @return result
func QueryDeviceGroupInfo(req pojo.QueryDeviceGroupInfoReq, claims comPojo.Claims) (code int, result pojo.QueryDeviceGroupInfoResp) {
	//根据user_group_id查询此用户组下所有的设备组
	code, result.DeviceGroupInfo = dao.DeviceGroupSelect(req, claims.UserGroupId)
	if error_code.Success != code || 0 == len(result.DeviceGroupInfo) {
		return
	}

	for i := 0; i < len(result.DeviceGroupInfo); i++ {

		code, deviceIds := commonDb.AuthDeviceIdSelect(claims.UserId, result.DeviceGroupInfo[i].DeviceGroupId, "")
		if error_code.Success != code {
			return code, pojo.QueryDeviceGroupInfoResp{}
		}

		result.DeviceGroupInfo[i].DeviceCount = len(deviceIds)

		if 0 == result.DeviceGroupInfo[i].DeviceCount {
			continue
		}

		code, result.DeviceGroupInfo[i].ChannelCount = dao.ChannelCountSelect(deviceIds)
		if error_code.Success != code {
			return code, pojo.QueryDeviceGroupInfoResp{}
		}
	}

	//获取设备分组总数
	code, result.TotalCount = dao.DeviceGroupCountSelect(req, claims.UserGroupId)
	if error_code.Success != code || 0 == result.TotalCount {
		return
	}

	result.PageNum = int(math.Ceil(float64(result.TotalCount) / float64(req.PageCount))) //page总数

	return
}

// DeviceOverview @description: 获取设备通道总数，在线数，在线率
// @parameter claims
// @return code
// @return result
func DeviceOverview(claims comPojo.Claims) (code int, result pojo.DeviceOverviewResp) {
	result = pojo.DeviceOverviewResp{}

	//获取用户的所有授权设备id
	code, deviceIds := commonDb.AuthDeviceIdSelect(claims.UserId, "", "")
	if error_code.Success != code || 0 == len(deviceIds) {
		return
	}

	//获取设备状态分组信息
	code, deviceStates := dao.DeviceStateSelect(deviceIds)
	if error_code.Success != code {
		return
	}

	for _, deviceState := range deviceStates {
		result.DeviceCount += deviceState.Count

		if 1 == deviceState.State { //1:在线
			result.OnlineDeviceCount = deviceState.Count
		}
	}

	result.OnlineDeviceRatio, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", (float64(result.OnlineDeviceCount)/float64(result.DeviceCount))*100), 64) // 保留2位小数

	//获取通道状态分组信息
	code, channelStates := dao.ChannelStateSelect(deviceIds)
	if error_code.Success != code {
		return
	}

	for _, channelState := range channelStates {
		result.ChannelCount += channelState.Count

		if 1 == channelState.ChannelStatus { //1:在线
			result.OnlineChannelCount = channelState.Count
		}
	}

	result.OnlineChannelRatio, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", (float64(result.OnlineChannelCount)/float64(result.ChannelCount))*100), 64) // 保留2位小数

	return
}

// BasicInfo @description: 获取设备详细基本信息
// @parameter deviceId
// @parameter claims
// @return code
// @return result
func BasicInfo(deviceId string, claims comPojo.Claims) (code int, result comPojo.DeviceInfo) {

	//该设备是否授权
	code = commonDb.DeviceIsAuth(deviceId, claims.UserId)
	if error_code.Success != code {
		return
	}

	//获取设备详细基本信息
	code, result = dao.DeviceInfoSelect(claims.UserId, deviceId)

	return
}

// DeviceUpdate @description: 编辑设备信息
// @parameter req
// @return code
func DeviceUpdate(req pojo.UpdateDeviceInfoReq, claims comPojo.Claims) (code int) {

	//该设备是否授权
	code = commonDb.DeviceIsAuth(req.DeviceId, claims.UserId)
	if error_code.Success != code {
		return
	}

	//更新设备授权表信息
	code = dao.UserAuthorizationUpdate2(req, claims.UserId, claims.UserGroupId)
	if error_code.Success != code {
		return
	}

	//更新设备表
	return dao.DeviceInfoUpdate(req)
}

// QueryResourceData @description: 查询设备使用资源数据
// @parameter deviceId
// @return int
// @return error
// @return *dto.QueryResourceDataResp
func QueryResourceData(deviceId string) (int, pojo.QueryResourceDataResp) {
	result := pojo.QueryResourceDataResp{}

	if v, ok := comPojo.DeviceOnlineEnum.Load(deviceId); ok {
		var deviceResource, ok2 = v.(comPojo.DeviceResourceReq)
		if !ok2 {
			return error_code.ErrInterfaceToObject, result
		}

		result.DeviceId = deviceResource.DeviceId
		result.CpuUsage = deviceResource.CpuUsage
		var freeMemory, _ = strconv.ParseFloat(deviceResource.MemoryFree, 32)
		var totalMemory, _ = strconv.ParseFloat(deviceResource.TotalMemory, 32)
		result.MemoryUsage = fmt.Sprintf("%.2f", (totalMemory-freeMemory)*100/totalMemory)
		var freeDisk, _ = strconv.ParseFloat(deviceResource.DiskFree, 32)
		var totalDisk, _ = strconv.ParseFloat(deviceResource.TotalDisk, 32)
		result.DiskUsage = fmt.Sprintf("%.2f", (totalDisk-freeDisk)*100/totalDisk)
		result.UpdateTime = data_convert.TimestampToFormat(deviceResource.TimeStamp)

	}
	return error_code.Success, result
}

// DeviceHeartbeat @description: 上报心跳，接收设备最新使用资源数据
// @parameter req
// @return code
func DeviceHeartbeat(req comPojo.DeviceResourceReq) int {

	if 0 == comPojo.DeviceIdEnum[req.DeviceId] {
		return error_code.ErrDeviceNotRegister
	}

	if _, ok := comPojo.DeviceOnlineEnum.Load(req.DeviceId); !ok { //之前不在线
		//设备状态变更为在线
		code := commonDb.MyDb.CommonUpdate(comPojo.Field{
			TableName: "device_info",
			Where:     "device_sn = ? and status in ('2','3') and delete_tag = ?"},
			map[string]interface{}{"status": 1, "update_time": time.Now()}, req.DeviceId, 0)
		if error_code.Success != code {
			return code
		}
	}

	//通道状态变更
	for _, chanInfo := range req.ChanInfo {

		_, ok := comPojo.ChannelOnlineEnum[req.DeviceId][strconv.Itoa(chanInfo.ChanId)]
		status := chanInfo.ChanStatus + 1

		if (!ok && 1 == status) || (ok && 1 != status) { //之前不在线,现在在线，变更为在线;之前在线,现在不在线，变更为不在线
			code := commonDb.MyDb.CommonUpdate(comPojo.Field{
				TableName: "channel_info",
				Where:     "device_id = ? and channel_id = ? and channel_status != ? and delete_tag = ?"},
				map[string]interface{}{"channel_status": status}, req.DeviceId, chanInfo.ChanId, status, 0)
			if error_code.Success != code {
				return code
			}
		}

		comPojo.UpdateCOE(req.DeviceId, strconv.Itoa(chanInfo.ChanId), strconv.Itoa(status))
	}

	req.TimeStampC = time.Now().Unix()                //记录当前时间戳
	comPojo.DeviceOnlineEnum.Store(req.DeviceId, req) //在线设备心跳记录(最新状态)

	return error_code.Success
}

// DeviceChannelInfo @description: 获取设备通道信息逻辑处理
// @parameter deviceId
// @parameter claims
// @return code
// @return result
func DeviceChannelInfo(req comPojo.QueryDeviceChannelInfoReq, claims comPojo.Claims) (code int, result comPojo.ChannelInfoResp) {

	//该设备是否授权
	code = commonDb.DeviceIsAuth(req.DeviceId, claims.UserId)
	if error_code.Success != code {
		return
	}

	//获取设备通道信息
	code, result.ChannelInfos = commonDb.ChannelInfoSelect(req.DeviceId, "", "", req.CurrentPageId, req.PageCount)
	if error_code.Success != code || 0 == len(result.ChannelInfos) {
		return
	}

	//获取设备通道总数
	code, result.TotalCount = dao.DeviceChannelCountSelect(req.DeviceId, "", "")
	if error_code.Success != code || 0 == result.TotalCount {
		return
	}

	result.PageNum = int(math.Ceil(float64(result.TotalCount) / float64(req.PageCount))) //page总数

	return
}

// ChannelBasicInfo @description: 获取通道基本信息
// @parameter deviceId
// @parameter channelId
// @parameter claims
// @return code
// @return result
func ChannelBasicInfo(deviceId, channelId string, claims comPojo.Claims) (code int, result comPojo.ChannelInfo) {

	//该设备是否授权
	code = commonDb.DeviceIsAuth(deviceId, claims.UserId)
	if error_code.Success != code {
		return
	}

	//获取通道详细基本信息
	code, result = dao.ChannelInfoSelect(deviceId, channelId)

	return
}
