package initialize

import "taylors/global"

func Init() {
	Log()
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		Mysql()
	}

	DBTables()

	if global.GVA_CONFIG.System.UseMultipoint {
		// 初始化redis服务
		Redis()
	}

	Cron()
	Crawler()
}
