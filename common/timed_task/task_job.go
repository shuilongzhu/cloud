package timed_task

import (
	"cloud/common/db"
	"cloud/common/pojo"
	"cloud/config"
	"cloud/error_code"
	"github.com/go-co-op/gocron"
	log "github.com/sirupsen/logrus"
	"time"
)

func GoCron() {
	if !(pojo.LocalServiceIp == config.MysqlConfigData.Host || pojo.ClientIp == config.MysqlConfigData.Host) { //同一个数据库连接确保定时任务只会由一个服务去执行
		return
	}

	timezone, _ := time.LoadLocation("Asia/Shanghai")
	scheduler := gocron.NewScheduler(timezone)

	//每10秒 处理设备心跳
	deviceHeartbeatJob := scheduler.Every(10).Seconds()
	deviceHeartbeatJob.Do(handleDeviceHeartbeat)

	scheduler.StartBlocking()
}

// HandleDeviceHeartbeat @description: 检测设备是否有心跳，如果设备最近上一次心跳时间距离现在超过12秒，则该设备离线
func handleDeviceHeartbeat() {
	pojo.OffLineDeviceIds = make([]string, 0)

	//遍历DeviceOnlineEnum
	pojo.DeviceOnlineEnum.Range(walk)

	if 0 != len(pojo.OffLineDeviceIds) { //把离线设备通道状态数据更新为离线

		code := db.MyDb.CommonUpdate(pojo.Field{
			TableName: "device_info",
			Where:     "device_sn in ? and status = ? and delete_tag = ?"},
			map[string]interface{}{"status": 2, "update_time": time.Now()}, pojo.OffLineDeviceIds, 1, 0)
		if error_code.Success != code {
			log.Errorf("code:%v,handleDeviceHeartbeat() fail", code)
		}

		code = db.MyDb.CommonUpdate(pojo.Field{
			TableName: "channel_info",
			Where:     "device_id in ? and delete_tag = ?"},
			map[string]interface{}{"channel_status": 2}, pojo.OffLineDeviceIds, 0) //通道更新为离线
		if error_code.Success != code {
			log.Errorf("code:%v,handleDeviceHeartbeat() fail", code)
		}
	}
}

func walk(key, value interface{}) bool {
	if v, ok := value.(pojo.DeviceResourceReq); !ok || 15 > time.Now().Unix()-v.TimeStampC { //最近一次心跳超过15秒,就置为离线
		return true
	}
	var deviceId = key.(string)
	pojo.OffLineDeviceIds = append(pojo.OffLineDeviceIds, deviceId)
	pojo.DeviceOnlineEnum.Delete(deviceId)
	delete(pojo.ChannelOnlineEnum, deviceId)
	return true
}
