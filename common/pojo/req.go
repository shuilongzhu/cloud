package pojo

type DeviceResourceReq struct {
	DeviceId    string       `json:"device_id" binding:"required"`    //设备ID
	CpuUsage    string       `json:"cpuload_percent"`                 //'CPU利用率',
	MemoryFree  string       `json:"memory_free"`                     //'内存未使用大小',
	MemoryUnit  string       `json:"memory_unit"`                     //内存单位 MB
	TotalMemory string       `json:"memory_total" binding:"required"` // '总内存大小',
	DiskUnit    string       `json:"disk_unit"`                       //'硬盘单位',
	DiskFree    string       `json:"disk_free"`                       //'硬盘未使用大小',
	TotalDisk   string       `json:"disk_total" binding:"required"`   //'总硬盘大小',
	TimeStamp   int64        `json:"time_stamp" binding:"required"`   //时间戳
	TimeStampC  int64        `json:"time_stamp_c"`                    //当前时间戳，因为数据在传输过程中会有延迟，到达cloud平台的时间戳
	ChanInfo    []ChanStatus `json:"chan_info"`
	Longitude   string       `json:"longitude"` //经度
	Latitude    string       `json:"latitude"`  //纬度
}

type ChanStatus struct {
	ChanId     int `json:"chan_id"`
	ChanStatus int `json:"chan_status"`
}

type QueryDeviceChannelInfoReq struct {
	DeviceId      string `json:"device_id" binding:"required"`                   //设备id
	CurrentPageId int    `json:"current_page_id" binding:"required" example:"1"` //当前页面id
	PageCount     int    `json:"page_count" binding:"required" example:"10"`     //每页条数
}
