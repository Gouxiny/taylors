package main

import (
	"gin-vue-admin/core"
	"gin-vue-admin/cron_job"
	"gin-vue-admin/global"
	"gin-vue-admin/initialize"
	//"runtime"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		initialize.Mysql()
	}
	initialize.DBTables()
	// 程序结束前关闭数据库链接
	defer global.GVA_DB.Close()
	cron_job.Run()
	core.RunWindowsServer()
}
