package ctx

import (
	"cloud/common/pojo"
	"cloud/common/utils/logger"
	"cloud/error_code"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// ResponseJson @description: 统一响应信息
// @parameter c
// @parameter er
// @parameter data
func ResponseJson(c *gin.Context, er int, data interface{}) {
	msg := error_code.StatusCode[er]
	c.Set("responseCode", er)

	requestId, _ := c.Get("request_id")

	c.JSON(http.StatusOK, pojo.ResponseTemplate{
		Code:      er,
		RequestId: strconv.FormatInt(requestId.(int64), 10),
		Data:      data,
		Message:   msg,
	})

	if 200 == er {
		return
	}
	//请求统一错误日志输出
	errorInfo := logger.CallerInfo(3)
	log.Errorf("url:%s requestId:%v errCode:%d errMessage:%s location:%v", c.Request.URL.Path, requestId, er, msg, errorInfo)
}
