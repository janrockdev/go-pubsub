package impl

import (
	nats "github.com/nats-io/nats.go"
)

func NatsConn(natsUrl string) (*nats.Conn, error) {
	return nats.Connect(nats.DefaultURL)
}

func NatsPing(natsUrl string, filter string) (any, error) {
	conn, err := NatsConn(natsUrl)
	if err != nil {
		return nil, err
	}
	return conn.Status().String(), nil
}
