package service

import (
	"github.com/janrockdev/go-pubsub/service/impl"
	"github.com/janrockdev/go-pubsub/utils"
)

func RedisService(redisUrl string, args string) error {
	//example
	res, err := impl.RedisPing(redisUrl, args)
	if err != nil {
		return err
	}
	utils.Logr.Info(res)
	return nil
}
