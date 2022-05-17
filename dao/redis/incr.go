package redis

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func RevertRedisActivityStock(commodityId string) error {
	RedisConnection, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		return errors.New(fmt.Sprintf("redis.Dial err=", err))
	}
	defer RedisConnection.Close()
	commodityStockKey := "commodity:stock:" + commodityId
	_, err = redis.Int(RedisConnection.Do("incr", commodityStockKey))
	if err != nil {
		return errors.New(fmt.Sprintf("redis incr activityStockKey error,err=", err))
	}
	return nil
}


