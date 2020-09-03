package loggerUtil

import (
	"fmt"
	"os"
	"path"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"votePlatfom/pkg/common"
)

var logger *logrus.Logger
var once sync.Once

func init() {
	once.Do(func() {
		logger = logrus.New()
	})
}

func getFIleName() string {
	logFilePath := os.Getenv("Log_FILE_PATH")
	logFileName := os.Getenv("LOG_FILE_NAME")

	// 校验目录是否存在并创建
	err := common.CheckAndMkdirAll(logFilePath)
	if err != nil {
		panic(err)
	}

	logFileName = logFileName + "-" + time.Now().Format("2006-01-02") + ".log"

	//日志文件
	fileName := path.Join(logFilePath, logFileName)

	// 校验文件是否存在并创建
	err = common.CheckAndCreateFile(fileName)
	if err != nil {
		panic(err)
	}

	return fileName
}

func setLoggerParam() {
	// 设置文件
	fileName := getFIleName()

	//打开文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("logger err: ", err)
	}

	//设置输出
	logger.Out = src

	//设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

func Info(args ...interface{}) {
	setLoggerParam()
	logger.Info(args)
}

func Error(args ...interface{}) {
	setLoggerParam()
	logger.Error(args)
}

func Warn(args ...interface{}) {
	setLoggerParam()
	logger.Warn(args)
}
