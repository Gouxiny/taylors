package utils

import "time"

func NowUnix() int64 {
	now := time.Now()
	cc := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	return cc.Unix()
}
