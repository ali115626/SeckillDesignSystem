package SellerService

import (
	Dao "SeckillDesign/dao/mysql"
	"SeckillDesign/dao/redis"
	"errors"
	"fmt"
)

func DoRevertStock(msgs ) error {
	orderNo := string(msgs.Body)
	var status *string
	var activityId *string
	activityId,status,err :=Dao.QueryOrderInfoStatus(orderNo)
	if *status != "2" {
		activityID := *activityId
		err =Dao.RevertActivityTableStock(activityID)
		if err!=nil{
			fmt.Println()
		}
		//	TODO  订单未支付   你肯定要revert一下redis吧
		err = redis.RevertRedisActivityStock(activityID)
		if err != nil{
			fmt.Println(errors.New(fmt.Sprintf("Redis revert stock error,err=",err)))
		}
		//	todo 从限购名单中移除
	}


	return nil
}
