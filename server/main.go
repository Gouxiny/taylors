package main

import (
	"taylors/core"
	"taylors/global"
	"taylors/initialize"
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
	initialize.InitModule()
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		initialize.Mysql()
	}
	initialize.DBTables()

	// 程序结束前关闭数据库链接
	defer func() {
		_ = global.GVA_DB.Close()
	}()

	core.RunWindowsServer()
}
