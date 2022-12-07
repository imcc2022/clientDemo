package main

import (
	"log"

	"clientDemo/service"
	"clientDemo/ui"

	"github.com/lxn/walk"
)

var logger *log.Logger
var mainWindow *walk.MainWindow
var isLogin bool

func init() {
	logger = service.CreateLogService()
	isLogin = service.CheckLoginStatus(logger)
}

func main() {
	ui.CreateGUI(logger, mainWindow, isLogin)
}
