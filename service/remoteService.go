package service

import "log"

func LoginService(username string, password string) bool {
	if username == "root" && password == "root" {
		return true
	}
	return false
}

func ConnectService(logger *log.Logger) {
	logger.Println("连接成功")
}

func DisconnectService(logger *log.Logger) {
	logger.Println("断开成功")
}
