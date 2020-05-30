package cron_job

import (
	"github.com/robfig/cron/v3"
	"taylors/global"
)

var cronObj *cron.Cron

func init() {
	cronObj = cron.New()
}

func Start() {
	_, err := cronObj.AddJob(global.GVA_CONFIG.Cron.SpecAll, MarketJob)
	if err != nil {
		global.GVA_LOG.Error("添加定时任务失败:", err)
		return
	}
	cronObj.Start()
}

func Stop() {
	cronObj.Stop()
}
