package error

import (
	"cloud/common/pojo"
	"fmt"
	log "github.com/sirupsen/logrus"
)

// CatchGoError @description: 常驻Goruntine 需要捕获panic
// @parameter name
// @parameter sendMessage(错误是否发送到钉钉)
func CatchGoError(name string, sendMessage bool) {
	if err := recover(); pojo.NIL != err {
		log.Errorln(fmt.Sprintf("%s 协程遇到错误，err：%v", name, err))
		if sendMessage {
			//DingHookTextAlarmGroup(fmt.Sprintf("【%s】%s 协程遇到错误,这可能是一个常驻协程，请及时关注！error：%s", config.EnvName, goruntineName, err))
		}
	}
}
