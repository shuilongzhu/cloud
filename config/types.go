package config

var (
	LogConfigData      = &LogConfig{}
	MysqlConfigData    = &MysqlConfig{}
	BaseInfoConfigData = &BaseInfoConfig{}
	DefaultConfigPath  = "D:\\AiRiA\\code\\video-compress-cloud\\conf\\config.toml"
)

// LogConfig
// @Description: 项目日志配置信息结构体
type LogConfig struct {
	LogPath          string `json:"project" ini:"log_Path"`
	LogFileName      string `json:"log_FileName" ini:"log_FileName"`
	LogMaxAge        int64  `json:"log_maxAge" ini:"log_maxAge"`             //日志最大保存时间（天）
	LogRotationTime  int64  `json:"log_rotationTime" ini:"log_rotationTime"` //日志切割时间间隔（小时）
	LogRotationCount int64  `json:"log_rotationCount" ini:"log_rotationCount"`
	LogRotationSize  int64  `json:"log_rotationSize" ini:"log_rotationSize"`
}

// MysqlConfig
// @Description: Mysql数据库相关的配置信息结构体
type MysqlConfig struct {
	Host        string `json:"mysql_hostname" ini:"mysql_hostname"`
	Port        string `json:"mysql_port" ini:"mysql_port"`
	UserName    string `json:"mysql_user" ini:"mysql_user"`
	Password    string `json:"mysql_pass" ini:"mysql_pass"`
	Database    string `json:"mysql_database_name" ini:"mysql_database_name"`
	Charset     string `json:"charset" ini:"charset"`
	MaxIdleConn int    `json:"max_idle_conn" ini:"max_idle_conn"`
	MaxOpenConn int    `json:"max_open_conn" ini:"max_open_conn"`
	ParseTime   string `json:"parse_time" ini:"parse_time"`
	Loc         string `json:"loc" ini:"loc"`
	Timeout     string `json:"timeout" ini:"timeout"`
	MaxLifeTime string `json:"max_life_time" ini:"max_life_time"`
	TablePre    string `json:"table_pre" ini:"table_pre"`
	SlowSqlTime string `json:"slow_sql_time" ini:"slow_sql_time"`
	PrintSqlLog bool   `json:"print_sql_log" ini:"print_sql_log"`
}

// BaseInfoConfig
// @Description: 项目基本信息结构体
type BaseInfoConfig struct {
	EnvName   string `json:"env_name" ini:"env_name"`
	IpAddress string `json:"ip_address" ini:"ip_address"`
}
