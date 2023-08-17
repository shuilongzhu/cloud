package pojo

type ResponseTemplate struct {
	Code      int         `json:"code"`       //请求code码
	RequestId string      `json:"request_id"` //请求id
	Data      interface{} `json:"data"`       //响应数据
	Message   string      `json:"message"`    //响应信息
}

type ChannelInfoResp struct {
	TotalCount   int           `json:"total_count" example:"100"` //总记录条数
	PageNum      int           `json:"page_num" example:"10"`     //总页数
	ChannelInfos []ChannelInfo `json:"channel_infos"`             //通道具体信息
}
