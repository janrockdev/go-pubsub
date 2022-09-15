package service

import (
	"github.com/janrockdev/go-pubsub/service/impl"
	"github.com/janrockdev/go-pubsub/utils"
)

func NatsService(natsUrl string, args string) error {
	//example
	res, err := impl.NatsPing(natsUrl, args)
	if err != nil {
		return err
	}
	utils.Logr.Info(res)
	return nil
}
