package redis

import (
	"SeckillDesign/constant"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
)


func QueryCommodityInfoFromRedis(commodityId string)(*constant.CommodityInfo,error){
	RedisConnection, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		//return nil,errors.New(fmt.Sprintf("redis.Dial err=", err))
	}
	defer RedisConnection.Close()
//warmUp:CommodityInfo:1


	commodityInfoKey := "warmUp:CommodityInfo:" + commodityId

	result, err :=redis.Bytes(RedisConnection.Do("get",commodityInfoKey))
	if err != nil {
		//return nil,errors.New(fmt.Sprintf("redis incr activityStockKey error,err=", err))
	}
	//todo 测试一下 这个里面没有key的时候该怎么办呢？

	if len(result)==0{
		//return 0,nil
		return nil,errors.New("缓存中没有此信息，此商品还未预热")
	}


	//activity:=constant.ActivityInfoForRedis{}
	commodityInfo:=constant.CommodityInfo{}

	err=json.Unmarshal(result,&commodityInfo)
	if err != nil{
		fmt.Println("result unmarshal into activity error,err=",err)
		return nil,err
	}
	return &commodityInfo,nil


}

func QueryActivityFromRedis(activityId string)(*constant.ActivityInfoForRedis,error){

	//redis.QueryActivityInfoFromRedis(activityId)
	//redis.QueryActivityInfoFromRedis(activityId)
	RedisConnection, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		return nil,errors.New(fmt.Sprintf("redis.Dial err=", err))
	}
	defer RedisConnection.Close()
	ActivityInfoKey := "warmUp:ActivityInfo:" + activityId

	result, err :=redis.Bytes(RedisConnection.Do("get",ActivityInfoKey))
	if err != nil {
		return nil,errors.New(fmt.Sprintf("redis incr activityStockKey error,err=", err))
	}
	//todo 测试一下 这个里面没有key的时候该怎么办呢？

	if len(result)==0{
		//return 0,nil
		return nil,errors.New("缓存中没有此信息，此商品还未预热")
	}


	activity:=constant.ActivityInfoForRedis{}
	err=json.Unmarshal(result,&activity)
	if err != nil{
		fmt.Println("result unmarshal into activity error,err=",err)
		return nil,err
	}
	return &activity,nil
}