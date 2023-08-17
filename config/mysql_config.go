package config

import (
	"cloud/common/db"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

// ConnectMysql
// @description:连接mysql
func ConnectMysql() error {

	// 用户名:密码@tcp(IP:port)/数据库?charset=utf8mb4&parseTime=True&loc=Local
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		MysqlConfigData.UserName, MysqlConfigData.Password, MysqlConfigData.Host, MysqlConfigData.Port,
		MysqlConfigData.Database, MysqlConfigData.Charset, MysqlConfigData.ParseTime, MysqlConfigData.Loc)

	// 连接额外配置信息
	gormConfig := gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   MysqlConfigInstance.TablePre, //表前缀
			SingularTable: true, //使用单数表名，启用该选项时，`User` 的表名应该是 `user`而不是users
		},
	}
	// 打印SQL设置
	if MysqlConfigData.PrintSqlLog {
		slowSqlTime, err := time.ParseDuration(MysqlConfigData.SlowSqlTime)
		if nil != err {
			logrus.Errorln("打印SQL设置失败：", err)
			return err
		}
		loggerNew := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold: slowSqlTime, //慢SQL阈值
			LogLevel:      logger.Info,
			Colorful:      true, // 彩色打印开启
		})
		gormConfig.Logger = loggerNew
	}
	var err error
	// 建立连接
	db.MyDb.Db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dataSourceName, // DSN data source name
		DefaultStringSize:         256,            // string 类型字段的默认长度
		DisableDatetimePrecision:  true,           // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,           // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,           // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,          // 根据当前 MySQL 版本自动配置
	}), &gormConfig)
	if nil != err {
		logrus.Errorln("mySQL建立连接失败：", err)
		return err
	}
	// 设置连接池信息
	sqlDB, err2 := db.MyDb.Db.DB()
	if nil != err2 {
		logrus.Errorln("mySQL设置连接池信息失败：", err2)
		return err2
	}
	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(MysqlConfigData.MaxIdleConn)
	// 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(MysqlConfigData.MaxOpenConn)
	// 设置了连接可复用的最大时间
	duration, err3 := time.ParseDuration(MysqlConfigData.MaxLifeTime)
	if nil != err3 {
		logrus.Errorln("mySQL设置连接可复用的最大时间失败：", err3)
		return err3
	}
	sqlDB.SetConnMaxLifetime(duration)
	//打印SQL配置信息
	//marshal, _ := json.Marshal(db.Stats())
	//fmt.Printf("数据库配置: %s \n", marshal)
	return nil
}

// ConnectMysqlTest
// @description: 用于测试方法调用连接mysql
func ConnectMysqlTest() {
	if err := MainInit(DefaultConfigPath); nil != err {
		return
	}

	if err := ConnectMysql(); nil != err {
		logrus.Errorln("InitMysql ConnectMysql err :", err)
	}
}
