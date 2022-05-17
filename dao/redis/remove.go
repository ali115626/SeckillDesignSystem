package redis

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func RemoveFromActivityBuyUser(commodityId string,userId string) error{


	RedisConnection, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		return errors.New(fmt.Sprintf("redis.Dial err=", err))
	}
	defer RedisConnection.Close()



	commodityBuyUser := "commodity:BuyUser"+commodityId

	var resultt int

	fmt.Printf("type=%T,userId=%v",userId,userId)

	resultt, err = redis.Int(RedisConnection.Do("SREM", commodityBuyUser, userId)) //写
	if err != nil {
		return errors.New(fmt.Sprintf("SREM remove from commodityBuyUser error,err=",err))
		//fmt.Println("redis sadd failed", err.Error())
	}

	fmt.Println("resultt====",resultt)

	return nil


}
