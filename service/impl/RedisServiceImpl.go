package impl

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func RedisConn(redisUrl string) *redis.Client {
	conn := redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return conn
}

func RedisPing(redisUrl string, args string) (any, error) {
	conn := RedisConn(redisUrl)
	res := conn.Ping(ctx)
	if res.String() != "ping: PONG" {
		err := errors.New(res.String())
		return nil, err
	} else {
		return "CONNECTED", nil
	}
}

func RedisSet(redisUrl string, key string, value string) (any, error) {
	conn := RedisConn(redisUrl)
	err := conn.Set(ctx, key, value, 0).Err()
	if err != nil {
		return nil, err
	} else {
		return key, nil
	}
}

//val, err := rdb.Get(ctx, "k1").Result()
//if err != nil {
//	panic(err)
//}
//fmt.Println("k1", val)

//val2, err := rdb.Get(ctx, "key2").Result()
//if err == redis.Nil {
//	fmt.Println("key2 does not exist")
//} else if err != nil {
//	panic(err)
//} else {
//	fmt.Println("key2", val2)
//}
// Output: key value
// key2 does not exist
