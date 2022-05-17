package DownLogic

import (
	//"SeckillDesign/constant"
	Dao "SeckillDesign/dao/mysql"
	"SeckillDesign/dao/redis"
	//"encoding/json"
	"errors"
	"fmt"
	"github.com/streadway/amqp"
)

func DoRevertStockLogic(msgs amqp.Delivery) error {
	//todo mq这里面传userId和commodityId  or 根据 orderNo  去查 userId 去查 activityId---> commodityId
	orderNo := string(msgs.Body)


	var status string
	var activityId string

	activityId,status,userId,err :=Dao.QueryOrderInfoStatus(orderNo)
	if status != "2" {
		activityID := activityId
		//todo 这里面还是加一个版本控制吧。就给mysql加就行了  redis不用加
		err =Dao.RevertActivityTableStock(activityID)
		if err!=nil{
			fmt.Println(err)
			return errors.New(fmt.Sprintf("RevertActivityTableStock",err))
		}
		//todo 这里面获取一下commodityId   ---从activityId获取commodityId

		commodityId,err :=Dao.GetCommodityIdFromActivityTable(activityID)

		if err != nil{
			//todo 如果在这里面return 了  也说明 redis的库存没有 revert成功  所以这里还是要 给mysql加上乐观锁呗
			return errors.New(fmt.Sprintf("Get CommodityId From ActivityTable error,err=",err))
		}


		//todo N你这个用commodityId不行吗？
		err = redis.RevertRedisActivityStock(commodityId)
		if err != nil{
			return errors.New(fmt.Sprintf("RevertActivityTableStock",err))
		}
		//	TODO  订单未支付   你肯定要revert一下redis吧
		//err = redis.RevertRedisActivityStock(activityID)
		//if err != nil{
		//	fmt.Println(errors.New(fmt.Sprintf("Redis revert stock error,err=",err)))
		//}
		//	todo 从限购名单中移除

		 //commodityBuyUser := "commodity:BuyUser"+commodityId

		err =redis.RemoveFromActivityBuyUser(commodityId,userId)

		//RemoveFromActivityBuyUser()


	//TODO	这个即使没有移除成功也没关系  差你这一个用户又能咋滴



	}


	return nil
}
