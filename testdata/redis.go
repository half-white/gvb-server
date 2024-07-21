package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 100,
	})
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := rdb.Ping().Result()
	if err != nil {
		logrus.Error(err)
		return
	}
}

func main() {
	err := rdb.Set("xxxx", "value1", 10*time.Second).Err()
	fmt.Println(err)
	cmd := rdb.Keys("*")
	keys, err := cmd.Result()
	fmt.Println(keys, err)
}
