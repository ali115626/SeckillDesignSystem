package redis

import (
	"SeckillDesign/constant"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
)


func SaveCommodityInfoRedis(commodityId string,CommodityInfo constant.CommodityInfo) error{
	RedisConnection, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		//return errors.New(fmt.Sprintf("redis.Dial err=", err))
	}
	defer RedisConnection.Close()

	CommodityInfoKey := "warmUp:CommodityInfo:" + commodityId

	CommodityInfoMar,err:=json.Marshal(CommodityInfo)
	if err != nil{
		return err
	}
	_, err = RedisConnection.Do("set",CommodityInfoKey,CommodityInfoMar)

	if err != nil {
		return err

		//return errors.New(fmt.Sprintf("redis incr activityStockKey error,err=", err))
	}
	return nil


}



func SaveActivityInfo(activityId string,activityInfo constant.ActivityInfoForRedis)error{

	RedisConnection, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		return errors.New(fmt.Sprintf("redis.Dial err=", err))
	}
	defer RedisConnection.Close()
	ActivityInfoKey := "warmUp:ActivityInfo:" + activityId
	activityInfoMar,err:=json.Marshal(activityInfo)
	if err != nil{
		return err
	}
	_, err =RedisConnection.Do("set",ActivityInfoKey,activityInfoMar)
	if err != nil {
		return errors.New(fmt.Sprintf("redis incr activityStockKey error,err=", err))
	}
	//if n==1{
	//	fmt.Println("redis保存成功！")
	//}
	return nil
}
