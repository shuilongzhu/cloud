package error_code

const (
	ErrInvalidJsonFormat = CommonErrorCodeBase + 1 + iota
	ErrTokenIssue
	ErrTokenTimeOut
	ErrTokenNothing
	ErrTokenOnlyOne
	ErrTokenSet
	ErrGormMethodParam
	ErrGormMethodCreate
	ErrGormMethodDelete
	ErrGormMethodSelect
	ErrGormMethodUpdate
	ErrInterfaceToObject
)

func init() {
	StatusCode[ErrInvalidJsonFormat] = "请求参数反序列化错误或者必填参数未填错误"
	StatusCode[ErrTokenIssue] = "Token校验失败"
	StatusCode[ErrTokenTimeOut] = "Token超时"
	StatusCode[ErrTokenNothing] = "Token为空"
	StatusCode[ErrTokenSet] = "Token生成失败"
	StatusCode[ErrTokenOnlyOne] = "该用户已重新登陆，您已被下线"
	StatusCode[ErrGormMethodParam] = "gorm method parameter error"
	StatusCode[ErrGormMethodCreate] = "gorm公共创建方法出错"
	StatusCode[ErrGormMethodDelete] = "gorm公共删除方法出错"
	StatusCode[ErrGormMethodSelect] = "gorm公共查询方法出错"
	StatusCode[ErrGormMethodUpdate] = "gorm公共更新方法出错"
	StatusCode[ErrInterfaceToObject] = "泛型转结构体失败"

}
