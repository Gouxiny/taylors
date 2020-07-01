package cron_job

import (
	"github.com/robfig/cron/v3"
	"taylors/global"
	"taylors/logger"
)

var cronObj *cron.Cron

func init() {
	cronObj = cron.New()
}

func Start() {
	_, err := cronObj.AddJob(global.GVA_CONFIG.Cron.SpecAll, MarketJob)
	if err != nil {
		logger.Error("添加定时任务失败:", err)
		return
	}
	cronObj.Start()
}

func Stop() {
	cronObj.Stop()
}
