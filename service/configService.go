package service

import (
	"encoding/json"
	"log"
	"os"
)

type Configs struct {
	IsLogin       bool   `json:"islogin"`
	TestAttribute string `json:"testAttribute"`
}

// 检查配置文件是否存在
func CheckConfig(logger *log.Logger) (dirExist bool, fileExist bool) {
	dirExist = checkConfigDir()
	if dirExist {
		fileExist = checkConfigFile()
		if !fileExist {
			createConfigFile(logger)
		}
		return dirExist, fileExist
	} else {
		if createConfigDir(logger) {
			createConfigFile(logger)
		}
		return dirExist, false
	}
}

func checkConfigDir() bool {
	_, err := os.Stat("config")
	if err == nil { // 文件夹存在
		return true
	} else {
		if os.IsExist(err) {
			return true
		} else {
			return false
		}
	}
}

func checkConfigFile() bool {
	_, err := os.Stat("config/config.txt")
	if err == nil { // 文件存在
		return true
	} else {
		if os.IsExist(err) {
			return true
		} else {
			return false
		}
	}
}

// 检查登录状态
func CheckLoginStatus(logger *log.Logger) bool {
	dirExist, fileExist := CheckConfig(logger)
	if dirExist && fileExist {
		configs := readConfig(logger)
		return configs.IsLogin
	}
	return false
}

// 保存登录状态
func SaveLoginStatus(logger *log.Logger) {
	config := readConfig(logger)
	config.IsLogin = true
	data, err := json.Marshal(&config)
	if err != nil {
		logger.Println("ERROR:json格式转换失败")
		return
	}
	err = os.WriteFile("config/config.txt", data, os.ModePerm)
	if nil != err {
		logger.Println("ERROR:配置文件写入失败")
	}
}

func readConfig(logger *log.Logger) Configs {
	var config Configs
	file, err := os.ReadFile("config/config.txt")
	if nil != err {
		return config
	}
	json.Unmarshal(file, &config)
	return config
}

func createConfigDir(logger *log.Logger) bool {
	err := os.Mkdir("config", os.ModePerm)
	if nil != err {
		logger.Println("ERROR:创建config文件夹失败")
		return false
	}
	return true
}

func createConfigFile(logger *log.Logger) {
	file, err := os.Create("config/config.txt")
	if nil != err {
		logger.Println("ERROR:创建config文件夹失败")
	}
	file.Close()
}
