package middleware

import (
	"bytes"
	"cloud/common/db"
	"cloud/common/pojo"
	"cloud/common/utils/ctx"
	"cloud/common/utils/data_convert"
	"cloud/error_code"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Cors @description: 跨域资源共享,以下为cors实现
// @return gin.HandlerFunc
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*") // 这是允许访问所有域

			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma,Authorization-Token,AuthorizationToken")
			//              允许跨域设置                                                                                                      可以返回其他子段

			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析

			c.Header("Access-Control-Max-Age", "172800")          // 缓存请求信息 单位为秒[48小时]
			c.Header("Access-Control-Allow-Credentials", "false") //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")             // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
	}
}

// Auth @description: 用户token认证
// @return gin.HandlerFunc
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("request_id", data_convert.NextId())

		if !whiteListUrl[c.Request.URL.Path] { //token认证
			code := ctx.GetTokenInfo(c)
			if error_code.Success != code {
				ctx.ResponseJson(c, code, nil)
				return
			}
		}
	}
}

// ApiRequestRecord @description: 服务api请求记录
// @return gin.HandlerFunc
func ApiRequestRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		//开始时间
		startTime := time.Now()
		//获取请求入参
		bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		//处理请求
		c.Next()
		if true {
			go RequestRecord(c, startTime, bodyBytes)
		}
	}
}

// RequestRecord @description: 记录请求操作日志
// @parameter c
// @parameter bodyBytes
// @parameter startTime
// @parameter endTime
func RequestRecord(c *gin.Context, startTime time.Time, bodyBytes []byte) {
	requestId, _ := c.Get("request_id")
	result := pojo.RequestRecord{Id: requestId.(int64), CreateTime: time.Now()}

	//防止入参过大，导致数据入库失败
	if 4000 > len(bodyBytes) {
		result.ReqParam = string(bodyBytes)
	}
	if reqParam, ok := c.Get("reqParam"); ok {
		result.ReqParam = reqParam.(string)
	}

	result.ReqMethod = c.Request.Method
	result.ReqUrl = c.Request.RequestURI

	code, _ := c.Get("responseCode")
	claims, _ := c.Get("claims")
	if nil != code {
		result.Code = code.(int)
	} else {
		result.Code = c.Writer.Status()
	}
	if nil != claims {
		tokenData := claims.(pojo.Claims)
		result.UserName = tokenData.UserName
	}

	result.ClientIp = c.Request.Header.Get("X-Real-IP")
	result.StartTime = startTime
	//执行时间(ms)
	result.ExecutionTime = time.Now().Sub(startTime).Milliseconds()
	if error_code.Success != result.Code {
		//utils.SendToRobotMaiCroCloud(utils.RequestRecordToStr(*result))
	}
	db.RequestRecordCreate(result)
}
