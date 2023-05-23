package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/izzie-it/nixio/log"
	apmgoredis "go.elastic.co/apm/module/apmgoredisv8/v2"
)

func Connect(host string, db int) (*redis.Client, error) {
	var client *redis.Client
	var err error

	for i := 1; i <= 10; i++ {
		log.Infof("redis connection try %d\n", i)

		client = redis.NewClient(&redis.Options{
			Addr: host,
			DB:   db,
		})
		client.AddHook(apmgoredis.NewHook())

		_, err = client.Ping(context.Background()).Result()
		if err == nil {
			break
		}

		time.Sleep(time.Second)
	}

	log.Info("redis connected")

	return client, err
}
