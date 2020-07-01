package initialize

import (
	"taylors/global"
	"taylors/logger"
)

func Log() {
	logger.Init(global.GVA_CONFIG.Log.FileName, global.GVA_CONFIG.Log.Level)
}
