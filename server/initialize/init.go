package initialize

import "taylors/global"

func Init() {
	Log()
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		Mysql()
	}
	DBTables()

	// 程序结束前关闭数据库链接
	defer func() {
		_ = global.GVA_DB.Close()
	}()

	if global.GVA_CONFIG.System.UseMultipoint {
		// 初始化redis服务
		Redis()
	}

	Cron()
	Crawler()
}
