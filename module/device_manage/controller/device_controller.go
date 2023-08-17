package controller

import (
	commonPojo "cloud/common/pojo"
	"cloud/common/utils/ctx"
	"cloud/error_code"
	"cloud/module/device_manage/pojo"
	"cloud/module/device_manage/service"
	"github.com/gin-gonic/gin"
)

// QueryDeviceTree One godoc
// @Summary      查询设备树目录结构
// @Description  查询此用户下的所有设备及通道
// @Tags         device_manage
// @Accept       json
// @Produce      plain
// @Param        Authorization-Token  header    string  true  "Authentication header"
// @Param        user_id    query      string  false  "user_id"
// @Param        device_status  query      string  false "device_status"
// @Param        product_type  query      string false  "product_type"
// @Success      200  {object}  pojo.DeviceGroup “successful”
// @Failure      400  {object}  pojo.ResponseTemplate
// @Router       /device/tree [get]
func QueryDeviceTree(c *gin.Context) {
	req := pojo.QueryDeviceTreeReq{}
	//1.参数校验解析
	if err := c.ShouldBind(&req); err != nil {
		ctx.ResponseJson(c, error_code.ErrInvalidJsonFormat, nil)
		return
	}

	claims, _ := c.Get("claims")

	//2.查询设备树
	code, data := service.SelectDeviceTree(req, claims.(commonPojo.Claims))
	//3.数据返回
	ctx.ResponseJson(c, code, data)
}

// AddDeviceGroup One godoc
// @Summary      添加设备分组
// @Description  添加设备分组
// @Tags         device_manage
// @Accept       json
// @Produce      plain
// @Param        Authorization-Token  header    string  true  "Authentication header"
// @Param        req    body      pojo.AddDeviceGroupReq  true  "AddDeviceGroupReq"
// @Success      200  {object}  pojo.ResponseTemplate “successful”
// @Failure      400  {object}  pojo.ResponseTemplate
// @Router       /device/add_group [post]
func AddDeviceGroup(c *gin.Context) {
	req := pojo.AddDeviceGroupReq{}

	//1.参数解析
	if err := c.BindJSON(&req); err != nil {
		ctx.ResponseJson(c, error_code.ErrInvalidJsonFormat, nil)
		return
	}
	//2.处理添加设备分组
	ctx.ResponseJson(c, service.AddDeviceGroup(req), nil)
}

// UpdateDeviceGroup One godoc
// @Summary      编辑设备分组
// @Description  编辑设备分组
// @Tags         device_manage
// @Accept       json
// @Produce      plain
// @Param        Authorization-Token  header    string  true  "Authentication header"
// @Param        req    body      pojo.DeviceGroupReq  true  "DeviceGroupReq"
// @Success      200  {object}  pojo.ResponseTemplate “successful”
// @Failure      400  {object}  pojo.ResponseTemplate
// @Router       /device/update_group [post]
func UpdateDeviceGroup(c *gin.Context) {
	req := pojo.DeviceGroupReq{}

	//1.参数解析
	if err := c.BindJSON(&req); err != nil {
		ctx.ResponseJson(c, error_code.ErrInvalidJsonFormat, nil)
		return
	}

	claims, _ := c.Get("claims")

	//2.处理编辑设备分组名称
	code := service.UpdateDeviceGroup(req, claims.(commonPojo.Claims))

	//3.数据返回
	ctx.ResponseJson(c, code, nil)
}

// DeleteDeviceGroup One godoc
// @Summary      删除设备分组
// @Description  删除设备分组
// @Tags         device_manage
// @Accept       json
// @Produce      plain
// @Param        Authorization-Token  header    string  true  "Authentication header"
// @Param        req    body      pojo.DeviceGroupReq  true  "DeviceGroupReq"
// @Success      200  {object}  pojo.ResponseTemplate “successful”
// @Failure      400  {object}  pojo.ResponseTemplate
// @Router       /device/delete_group [post]
func DeleteDeviceGroup(c *gin.Context) {
	req := pojo.DeviceGroupReq{}

	//1.参数解析
	if err := c.BindJSON(&req); err != nil {
		ctx.ResponseJson(c, error_code.ErrInvalidJsonFormat, nil)
		return
	}
	claims, _ := c.Get("claims")

	//2.处理删除设备分组
	code := service.DeleteDeviceGroup(req, claims.(commonPojo.Claims))
	//3.数据返回
	ctx.ResponseJson(c, code, nil)
}

// QueryDeviceGroupInfo One godoc
// @Summary      查询设备组信息列表 【分页】
// @Description  查询设备组信息列表 【分页】
// @Tags         device_manage
// @Accept       json
// @Produce      plain
// @Param        Authorization-Token  header    string  true  "Authentication header"
// @Param        req    body      pojo.QueryDeviceGroupInfoReq  true  "QueryDeviceGroupInfoReq"
// @Success      200  {object}  pojo.QueryDeviceGroupInfoResp “successful”
// @Failure      400  {object}  pojo.ResponseTemplate
// @Router       /device/group_info [post]
func QueryDeviceGroupInfo(c *gin.Context) {
	req := pojo.QueryDeviceGroupInfoReq{}

	if err := c.BindJSON(&req); err != nil {
		ctx.ResponseJson(c, error_code.ErrInvalidJsonFormat, nil)
		return
	}

	claims, _ := c.Get("claims")
	//处理查询设备组信息列表
	code, data := service.QueryDeviceGroupInfo(req, claims.(commonPojo.Claims))

	ctx.ResponseJson(c, code, data)
}

// DeviceOverview One godoc
// @Summary      设备总览，所有设备通道数量状态
// @Description  设备总览，所有设备通道数量状态
// @Tags         device_manage
// @Accept       json
// @Produce      plain
// @Param        Authorization-Token  header    string  true  "Authentication header"
// @Success      200  {object}  pojo.DeviceOverviewResp “successful”
// @Failure      400  {object}  pojo.ResponseTemplate
// @Router       /device/overview [post]
func DeviceOverview(c *gin.Context) {

	claims, _ := c.Get("claims")
	//获取设备总览信息
	code, data := service.DeviceOverview(claims.(commonPojo.Claims))

	ctx.ResponseJson(c, code, data)
}

// BasicInfo One godoc
// @Summary      设备详细基本信息
// @Description  设备详细基本信息
// @Tags         device_manage
// @Accept       json
// @Produce      plain
// @Param        Authorization-Token  header    string  true  "Authentication header"
// @Param        device_id    query      string  true  "device_id"
// @Success      200  {object}  pojo.DeviceInfo “successful”
// @Failure      400  {object}  pojo.ResponseTemplate
// @Router       /device/basic_info [get]
func BasicInfo(c *gin.Context) {

	claims, _ := c.Get("claims")

	//获取设备详细基本信息
	code, data := service.BasicInfo(c.Query("device_id"), claims.(commonPojo.Claims))

	ctx.ResponseJson(c, code, data)
}

// DeviceUpdate One godoc
// @Summary      编辑设备信息
// @Description  编辑设备信息
// @Tags         device_manage
// @Accept       json
// @Produce      plain
// @Param        Authorization-Token  header    string  true  "Authentication header"
// @Param        req    body      pojo.UpdateDeviceInfoReq  true  "UpdateDeviceInfoReq"
// @Success      200  {object}  pojo.ResponseTemplate “successful”
// @Failure      400  {object}  pojo.ResponseTemplate
// @Router       /device/update [post]
func DeviceUpdate(c *gin.Context) {

	req := pojo.UpdateDeviceInfoReq{}

	//1.参数校验解析
	if err := c.ShouldBind(&req); err != nil {
		ctx.ResponseJson(c, error_code.ErrInvalidJsonFormat, nil)
		return
	}

	claims, _ := c.Get("claims")

	//编辑设备信息
	code := service.DeviceUpdate(req, claims.(commonPojo.Claims))

	ctx.ResponseJson(c, code, nil)
}

// QueryResourceData One godoc
// @Summary      查询设备使用资源数据
// @Description  查询设备使用资源数据
// @Tags         device_manage
// @Accept       json
// @Produce      plain
// @Param        Authorization-Token  header    string  true  "Authentication header"
// @Param        device_id    query      string  true  "device_id"
// @Success      200  {object}  pojo.ResponseTemplate “successful”
// @Failure      400  {object}  pojo.ResponseTemplate
// @Router       /device/resource_data [get]
func QueryResourceData(c *gin.Context) {

	//处理查询设备使用资源数据
	code, data := service.QueryResourceData(c.Query("device_id"))

	//数据返回
	ctx.ResponseJson(c, code, data)
}

// DeviceHeartbeat One godoc
// @Summary      定时上报设备心跳,接收设备最新资源数据
// @Description  定时上报设备心跳,接收设备最新资源数据
// @Tags         device_manage
// @Accept       json
// @Produce      plain
// @Param        Authorization-Token  header    string  true  "Authentication header"
// @Param        req    body      pojo.QueryDeviceGroupInfoReq  true  "QueryDeviceGroupInfoReq"
// @Success      200  {object}  pojo.DeviceOverviewResp “successful”
// @Failure      400  {object}  pojo.ResponseTemplate
// @Router       /device/heartbeat [post]
func DeviceHeartbeat(c *gin.Context) {
	req := commonPojo.DeviceResourceReq{}

	//1.参数校验解析
	if err := c.ShouldBind(&req); err != nil {
		ctx.ResponseJson(c, error_code.ErrInvalidJsonFormat, nil)
		return
	}

	//2.上报心跳，接收设备最新使用资源数据
	code := service.DeviceHeartbeat(req)

	//3.数据返回
	ctx.ResponseJson(c, code, nil)
}

// DeviceChannelInfo One godoc
// @Summary      获取设备通道信息 【分页】
// @Description  获取设备通道信息 【分页】
// @Tags         device_manage
// @Accept       json
// @Produce      plain
// @Param        Authorization-Token  header    string  true  "Authentication header"
// @Param        req    body      pojo.QueryDeviceChannelInfoReq  true  "QueryDeviceChannelInfoReq"
// @Success      200  {object}  pojo.DeviceOverviewResp “successful”
// @Failure      400  {object}  pojo.ResponseTemplate
// @Router       /device/channel_info [post]
func DeviceChannelInfo(c *gin.Context) {

	req := commonPojo.QueryDeviceChannelInfoReq{}

	//1.参数校验解析
	if err := c.ShouldBind(&req); err != nil {
		ctx.ResponseJson(c, error_code.ErrInvalidJsonFormat, nil)
		return
	}

	claims, _ := c.Get("claims")

	//2.获取设备通道信息
	code, data := service.DeviceChannelInfo(req, claims.(commonPojo.Claims))

	//3.数据返回
	ctx.ResponseJson(c, code, data)
}

// ChannelBasicInfo One godoc
// @Summary      获取通道详细基本信息
// @Description  获取通道详细基本信息
// @Tags         device_manage
// @Accept       json
// @Produce      plain
// @Param        Authorization-Token  header    string  true  "Authentication header"
// @Param        device_id    query      string  true  "device_id"
// @Param        channel_id    query      string  true  "channel_id"
// @Success      200  {object}  pojo.DeviceInfo “successful”
// @Failure      400  {object}  pojo.ResponseTemplate
// @Router       /device/channel/basic_info [get]
func ChannelBasicInfo(c *gin.Context) {

	claims, _ := c.Get("claims")

	deviceId := c.Query("device_id")
	channelId := c.Query("channel_id")

	if "" == deviceId || "" == channelId {
		ctx.ResponseJson(c, error_code.ErrInvalidJsonFormat, nil)
		return
	}

	//获取设备详细基本信息
	code, data := service.ChannelBasicInfo(c.Query("device_id"), c.Query("channel_id"), claims.(commonPojo.Claims))

	ctx.ResponseJson(c, code, data)
}
