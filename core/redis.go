package core

import (
	"context"
	"gvb_server/global"
	"time"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

func ConnectRedis() *redis.Client {
	return ConnectDB(0)
}

func ConnectDB(db int) *redis.Client {
	redisConf := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr(),
		Password: redisConf.Password,
		DB:       0,
		PoolSize: redisConf.PoolSize,
	})
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := rdb.Ping().Result()
	if err != nil {
		logrus.Error("redis连接失败")
	}
	return rdb
}
