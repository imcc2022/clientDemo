package service

import (
	"log"
	"os"
)

// 创建日志服务
func CreateLogService() *log.Logger {
	file := "log.txt"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm) // 以读写方式打开文件，如果没有该名称文件则自动创建，以追加的方式写入文件
	if err != nil {
		return nil
	}
	logger := log.New(logFile, "[logInfo:]", log.LstdFlags|log.Lshortfile|log.LUTC) // 将文件设置为loger作为输出
	return logger
}
