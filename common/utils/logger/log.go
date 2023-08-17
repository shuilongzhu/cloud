package logger

import (
	"log"
	"os"
)

type LoggerFile struct {
	MyLogger *log.Logger
	File     *os.File
	FilePath string
}

// 记录器
func (L *LoggerFile) NewLogger() { // 返回记录器
	//fileName := "../log/xxx_debug.log"
	logFile, err := os.Create(L.FilePath)
	L.File = logFile
	if err != nil {
		log.Fatalln("open file error !")
	}
	//debugLog := log.New(logFile, "[Debug]", log.Llongfile)
	debugLog := log.New(logFile, "[Debug]", log.LstdFlags)
	//debugLog.Println("A debug message here")

	debugLog.SetPrefix("[Info]")
	//debugLog.Println("A Info Message here ")

	//debugLog.SetFlags(debugLog.Flags() | log.LstdFlags)
	//debugLog.Println("A different prefix")
	L.MyLogger = debugLog

}

func (L *LoggerFile) AddToLogger(logstrs []string) {
	if len(logstrs) == 0 {
		return
	}
	for _, logstr := range logstrs {
		L.MyLogger.Println(logstr)
	}
}

func (L *LoggerFile) SetPrefixInfo() {
	L.MyLogger.SetPrefix("[Info]")
}

func (L *LoggerFile) SetPrefixDebug() {
	L.MyLogger.SetPrefix("[Debug]")
	L.MyLogger.SetFlags(L.MyLogger.Flags() | log.LstdFlags)
}

func (L *LoggerFile) SetFlags() {
	L.MyLogger.SetFlags(L.MyLogger.Flags() | log.LstdFlags)
}

func (L *LoggerFile) LoggerClose() {
	_ = L.File.Close()
}
