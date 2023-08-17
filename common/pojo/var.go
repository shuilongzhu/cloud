package pojo

import (
	"sync"
	"time"
)

var (
	NIL               = *new(interface{})
	AdminUserId       string
	ClientIp          string
	RoleEnum          = map[int]string{}
	DeviceIdEnum      = map[string]int64{}
	ChannelIdEnum     = map[string]int64{}
	DeviceOnlineEnum  = sync.Map{}
	ChannelOnlineEnum = map[string]map[string]int64{}
	OffLineDeviceIds  = make([]string, 0)
)

func init() {
	RoleMapInit()
}

func RoleMapInit() {
	RoleEnum[0] = "普通用户"
	RoleEnum[1] = "系统超级管理员"
	RoleEnum[2] = "运维观察员"
	RoleEnum[3] = "模型管理员"
}

func SetCOE(key1, key2 string, value int64) {
	if _, ok := ChannelOnlineEnum[key1]; !ok {
		var tempMap = make(map[string]int64)
		ChannelOnlineEnum[key1] = tempMap
	}
	ChannelOnlineEnum[key1][key2] = value
}

func UpdateCOE(deviceId, channelId, status string) {
	if "1" == status {
		SetCOE(deviceId, channelId, time.Now().Unix())
		return
	}
	delete(ChannelOnlineEnum[deviceId], channelId)
}
