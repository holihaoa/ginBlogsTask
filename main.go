package main

import (
	"ginBlogsTask/server/core"
	"ginBlogsTask/server/global"
	"ginBlogsTask/server/initialize"

	"go.uber.org/zap"
)

func main() {
	initializeSystem()
	core.RunServer()
}

func initializeSystem() {
	global.GVA_VP = core.Viper() // 初始化Viper
	global.GVA_LOG = core.Zap()  // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	if global.GVA_DB != nil {
		initialize.RegisterTables() // 初始化表
	}
}
