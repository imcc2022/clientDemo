package ui

import (
	"clientDemo/service"
	"log"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
)

var loginGUIWidth, loginGUIHeight int

func init() {
	loginGUIWidth = 400
	loginGUIHeight = 100
}

func CreateGUI(logger *log.Logger, mw *walk.MainWindow, isLogin bool) {
	if !isLogin {
		LoginGUI(logger, mw)
	} else {
		MainGUI(logger)
	}
}

func LoginGUI(logger *log.Logger, mw *walk.MainWindow) {
	var username, password *walk.LineEdit
	mw, err := walk.NewMainWindow()
	if err != nil {
		logger.Fatalln("LoginUI create error")
		return
	}
	MainWindow{
		AssignTo: &mw,
		Title:    "登录",
		Size:     Size{Width: loginGUIWidth, Height: loginGUIHeight},
		Layout:   VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{
						Text: "                            用户名：",
					},
					LineEdit{
						AssignTo: &username,
						MaxSize:  Size{Width: 150, Height: 40},
					},
				},
			},
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{
						Text: "                            密    码：",
					},
					LineEdit{
						AssignTo:     &password,
						MaxSize:      Size{Width: 150, Height: 40},
						PasswordMode: true,
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					PushButton{
						Text:    "登录",
						MaxSize: Size{Width: 50, Height: 40},
						OnClicked: func() {
							flag := service.LoginService(username.Text(), password.Text())
							if flag {
								service.SaveLoginStatus(logger)
								mw.Close()
								MainGUI(logger)
							} else {
								walk.MsgBox(mw, "错误", "用户名或密码错误", walk.MsgBoxIconError)
							}
						},
					},
				},
			},
		},
	}.Create()
	setWindowMaxDisable(mw)
	setWindowSizeDisable(mw)
	setWindowIcon(logger, mw)
	setOriginLocation(mw, loginGUIWidth, loginGUIHeight)
	setSysTray(logger, mw)
	// setSysNotification(logger, mw, "syscloud", "正在连接中")
}

// 禁用窗口最大化
func setWindowMaxDisable(mw *walk.MainWindow) {
	hwnd := mw.Handle()
	currStyle := win.GetWindowLong(hwnd, win.GWL_STYLE)
	win.SetWindowLong(hwnd, win.GWL_STYLE, currStyle&^win.WS_MAXIMIZEBOX)
}

// 禁用窗口大小修改
func setWindowSizeDisable(mw *walk.MainWindow) {
	hwnd := mw.Handle()
	currStyle := win.GetWindowLong(hwnd, win.GWL_STYLE)
	win.SetWindowLong(hwnd, win.GWL_STYLE, currStyle&^win.WS_SIZEBOX)
}

func MainGUI(logger *log.Logger) {
	mw, err := walk.NewMainWindow()
	MainWindow{
		AssignTo: &mw,
		Title:    "主窗口",
		Size:     Size{Width: loginGUIWidth, Height: loginGUIHeight},
		Layout:   VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{
						Text: "                        主窗口信息：",
					},
					LineEdit{
						MaxSize:  Size{Width: 150, Height: 40},
						Text:     "这是主窗口",
						ReadOnly: true,
					},
				},
			},
		},
	}.Create()
	if err != nil {
		walk.MsgBox(mw, "错误", "程序打开失败", walk.MsgBoxIconError)
		walk.App().Exit(0)
		return
	}
	// mw.SetVisible(true)
	setSysTray(logger, mw)
	// setSysNotification(logger, mw, "syscloud", "正在连接中")
}

// 设置窗口初始位置
func setOriginLocation(mw *walk.MainWindow, width int, height int) {
	mw.SetX(int(win.GetSystemMetrics(win.SM_CXSCREEN))/2 - width/2)
	mw.SetY(int(win.GetSystemMetrics(win.SM_CYSCREEN))/2 - height/2)
}

// 设置icon
func setWindowIcon(logger *log.Logger, mw *walk.MainWindow) {
	icon := getIcon(logger)
	mw.SetIcon(icon)
}

// 获取icon
func getIcon(logger *log.Logger) walk.Image {
	icon, err := walk.Resources.Image("resource/icon.ico")
	if err != nil {
		logger.Println("ERROR:获取icon失败", err)
	}
	return icon
}

// 设置系统托盘
func setSysTray(logger *log.Logger, mw *walk.MainWindow) {
	notifyIcon, err := walk.NewNotifyIcon(mw)
	if nil != err {
		logger.Println("ERROR:设置提醒图标失败")
	}
	defer notifyIcon.Dispose()
	icon := getIcon(logger)
	err = notifyIcon.SetIcon(icon)
	if nil != err {
		logger.Println("ERROR:设置图标失败")
	}
	notifyIcon.SetVisible(true)
	if err := notifyIcon.SetToolTip("syscloud客户端"); err != nil { // 设置系统托盘悬浮信息
		logger.Println("ERROR:设置悬浮信息失败", err)
	}
	setConnectMenu(logger, notifyIcon)
	setDisconnectMenu(logger, notifyIcon)
	setExitMenu(logger, notifyIcon)
	notifyIcon.ShowInfo("syscloud", "正在连接中")
	// notifyIcon.ShowInfo("syscloud", "连接成功")
	mw.Run()
}

// 系统托盘右键菜单(退出)
func setExitMenu(logger *log.Logger, ni *walk.NotifyIcon) {
	exitAction := walk.NewAction()
	if err := exitAction.SetText("退出"); err != nil {
		logger.Println(err)
	}
	//Exit 实现的功能
	exitAction.Triggered().Attach(func() { walk.App().Exit(0) })
	if err := ni.ContextMenu().Actions().Add(exitAction); err != nil {
		logger.Println(err)
	}
}

// 系统托盘右键菜单(连接)
func setConnectMenu(logger *log.Logger, ni *walk.NotifyIcon) {
	connectAction := walk.NewAction()
	if err := connectAction.SetText("连接"); err != nil {
		logger.Println(err)
	}
	connectAction.Triggered().Attach(func() {
		// 连接功能
		service.ConnectService(logger)
	})
	if err := ni.ContextMenu().Actions().Add(connectAction); err != nil {
		logger.Println(err)
	}
}

// 系统托盘右键菜单(断开)
func setDisconnectMenu(logger *log.Logger, ni *walk.NotifyIcon) {
	disconnectAction := walk.NewAction()
	if err := disconnectAction.SetText("断开"); err != nil {
		logger.Println(err)
	}
	disconnectAction.Triggered().Attach(func() {
		// 断开功能
		service.DisconnectService(logger)
	})
	if err := ni.ContextMenu().Actions().Add(disconnectAction); err != nil {
		logger.Println(err)
	}
}
