package db

import (
	"cloud/common/pojo"
	"cloud/common/utils/logger"
	"cloud/error_code"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const skip = 3

var MyDb Gorm

type DbMethod interface {
	CommonCreate(tableName string, object interface{}) int
	CommonDelete(field pojo.Field, args ...interface{}) int
	CommonSelect(result interface{}, field pojo.Field, args ...interface{}) int
	CommonUpdate(field pojo.Field, mapInfo map[string]interface{}, args ...interface{}) int
}

type Gorm struct {
	Db       *gorm.DB `json:"db"`
	Ip       string   `json:"ip"`
	DataBase string   `json:"dataBase"`
}

// CommonCreate @description: gorm公共创建方法
// @receiver gorm
// @parameter tableName
// @parameter object
// @return int
func (gorm *Gorm) CommonCreate(tableName string, object interface{}) int {

	if "" == tableName {
		return error_code.ErrGormMethodParam
	}

	if err := gorm.Db.Table(tableName).Create(object).Error; nil != err {
		//打印错误日志信息
		errorInfo := logger.CallerInfo(skip)
		log.Errorln(errorInfo, fmt.Sprintf("Create tableName:%s  error:%s", tableName, err))
		return error_code.ErrGormMethodCreate
	}

	return error_code.Success
}

// CommonDelete @description: gorm公共删除方法
// @receiver gorm
// @parameter field
// @parameter args
// @return int
func (gorm *Gorm) CommonDelete(field pojo.Field, args ...interface{}) int {

	if !gormParamCheck(field) {
		return error_code.ErrGormMethodParam
	}

	operate := gorm.Db.Table(field.TableName)
	if "" != field.Where {
		operate = operate.Where(field.Where, args...)
	}

	operate = operate.Delete(nil)
	if nil != operate.Error {
		//打印错误日志信息
		errorInfo := logger.CallerInfo(skip)
		log.Errorln(errorInfo, fmt.Sprintf("Delete tableName:%s  error:%s", field.TableName, operate.Error))
		return error_code.ErrGormMethodDelete
	}

	return error_code.Success
}

// CommonSelect @description: gorm公共查询方法
// @receiver gorm
// @parameter result
// @parameter field
// @parameter args
// @return int
func (gorm *Gorm) CommonSelect(result interface{}, field pojo.Field, args ...interface{}) int {

	if !gormParamCheck(field) {
		return error_code.ErrGormMethodParam
	}

	if "" == field.Pluck {
		field.Pluck = "*"
	}
	operate := gorm.Db.Table(field.TableName).Select(field.Pluck)
	if "" != field.Joins {
		operate = operate.Joins(field.Joins)
	}

	//条件拼接
	if "" != field.Where {
		operate = operate.Where(field.Where, args...)
	}

	if "" != field.Group {
		operate = operate.Group(field.Group)
	}
	if "" != field.Sort {
		operate = operate.Order(field.Sort)
	}
	if 0 != field.Offset {
		operate = operate.Offset(field.Offset)
	}
	if 0 != field.Limit {
		operate = operate.Limit(field.Limit)
	}

	operate = operate.Scan(result)
	if nil != operate.Error {
		//打印错误日志信息
		errorInfo := logger.CallerInfo(skip)
		log.Errorln(errorInfo, fmt.Sprintf("Select tableName:%s  error:%s", field.TableName, operate.Error))
		return error_code.ErrGormMethodSelect
	}
	return error_code.Success
}

// CommonUpdate @description: gorm公共更新方法
// @receiver gorm
// @parameter field
// @parameter mapInfo
// @parameter args
// @return int
func (gorm *Gorm) CommonUpdate(field pojo.Field, mapInfo map[string]interface{}, args ...interface{}) int {

	if !gormParamCheck(field) {
		return error_code.ErrGormMethodParam
	}

	operate := gorm.Db.Table(field.TableName)
	if "" != field.Where {
		operate = operate.Where(field.Where, args...)
	}

	operate = operate.Updates(mapInfo)
	if nil != operate.Error {
		//打印错误日志信息
		errorInfo := logger.CallerInfo(skip)
		log.Errorln(errorInfo, fmt.Sprintf("Update tableName:%s  error:%s", field.TableName, operate.Error))
		return error_code.ErrGormMethodUpdate
	}

	return error_code.Success
}

// gormParamCheck @description: gorm方法入参校验
// @parameter field
// @return bool(是否合规) true:合规；false:不合规
func gormParamCheck(field pojo.Field) bool {
	if "" == field.TableName {
		return false
	}
	return true
}
