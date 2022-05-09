package redis

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func RevertRedisActivityStock(activityId string) error {
	RedisConnection, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		return errors.New(fmt.Sprintf("redis.Dial err=", err))
	}
	defer RedisConnection.Close()
	activityStockKey := "commodity:stock:" + activityId
	_, err = redis.Int(RedisConnection.Do("incr", activityStockKey))
	if err != nil {
		return errors.New(fmt.Sprintf("redis incr activityStockKey error,err=", err))
	}
	return nil
}
