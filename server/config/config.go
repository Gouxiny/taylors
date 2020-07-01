package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"taylors/cron_job"
	"taylors/global"
)

const defaultConfigFile = "config.yaml"

func Init() (err error) {
	v := viper.New()
	v.SetConfigFile(defaultConfigFile)
	err = v.ReadInConfig()
	if err != nil {
		return
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		err = v.Unmarshal(&global.GVA_CONFIG)
		if err != nil {
			return
		}

		//监听配置文件
		cron_job.Stop()
		cron_job.Start()
	})
	err = v.Unmarshal(&global.GVA_CONFIG)
	if err != nil {
		return
	}
	global.GVA_VP = v
	return
}
