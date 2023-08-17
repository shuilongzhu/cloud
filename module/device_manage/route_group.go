package device_manage

import (
	"cloud/module/device_manage/controller"
	"github.com/gin-gonic/gin"
)

// DeviceAPIInit @description: 设备管理路由，baseUrl:/cloud/api/device
// @parameter route
func DeviceAPIInit(route *gin.RouterGroup) {
	route.GET("/tree", controller.QueryDeviceTree) //查询设备树目录结构

	route.POST("/add_group", controller.AddDeviceGroup) //添加设备分组

	route.POST("/update_group", controller.UpdateDeviceGroup) //编辑设备分组名称

	route.POST("/delete_group", controller.DeleteDeviceGroup) //删除设备分组

	route.POST("/group_info", controller.QueryDeviceGroupInfo) //查询设备组信息列表【分页】

	route.POST("/overview", controller.DeviceOverview) //设备总览，所有设备通道数量状态

	route.GET("/basic_info", controller.BasicInfo) //获取设备详细基本信息

	route.POST("/update", controller.DeviceUpdate) //编辑设备信息

	route.POST("/heartbeat", controller.DeviceHeartbeat) //定时上报设备心跳,接收设备最新资源数据

	route.GET("/resource_data", controller.QueryResourceData) //查询设备最新资源数据

	route.POST("/channel_info", controller.DeviceChannelInfo) //获取设备通道信息【分页】

	route.GET("/channel/basic_info", controller.ChannelBasicInfo) //获取通道详细基本信息

}
