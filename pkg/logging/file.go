package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	LogSavePath = "runtime/logs/"
	LogSaveName = "log"
	LogFileExt  = "log"
	TimeFormat  = "20210713"
)

// 获取日志路径
func getLogFilePath() string {
	return LogSavePath
}

// 获取实际文件全路径
func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

// 打开log文件并且返回指针对象进行操作
func OpenLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission : %v", err)
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile:%v", err)
	}
	return handle
}

// 创建文件
func mkDir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)

	if err != nil {
		panic(err)
	}
}
