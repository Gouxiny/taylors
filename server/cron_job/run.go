package cron_job

import (
	"github.com/robfig/cron/v3"
)

func Run() {
	c := cron.New()
	_, err := c.AddJob("30 23 * * ?", AllJob)
	if err != nil {
		panic(err)
	}
	c.Start()
}
