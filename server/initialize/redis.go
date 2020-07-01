package initialize

import (
	"github.com/go-redis/redis"
	"taylors/global"
	"taylors/logger"
)

func Redis() {
	redisCfg := global.GVA_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		logger.Error(err)
	} else {
		logger.Info("redis connect ping response:", pong)
		global.GVA_REDIS = client
	}
}
