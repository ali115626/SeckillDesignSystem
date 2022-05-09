package mqConsumer

import (
	"SeckillDesign/SellerService"
	"SeckillDesign/constant"
	Dao "SeckillDesign/dao/mysql"
	redis "SeckillDesign/dao/redis"
	"errors"
	"fmt"
	//"github.com/gomodule/redigo/redis"

)

func PayStatusCheckListener() {
	queueName := constant.DeadQueue
	delivery := DeadQueueConsumer(queueName)
	for msgs := range delivery {
		//TODO 处理业务逻辑 成功就ack  不成功就reject
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
		err = SellerService.DoRevertStock()

		if err != nil {

		} else {
			msgs.Ack(false)

		}

		//TODO 他应该是在这里面去写业务逻辑  revertStock   如果业务逻辑处理成功  就手动去确认 失败  就 返回 NACK  但是现在还是不知道
		//TODO 如何去 合理地  不断监听消息  缺一个 if {} else{}

		//todo   先看看它能不能work
		//TODO 手动确认当前消息
	}

	//}()
	//<-forever
	//
	//
	//
	//message :=<-delivery
	//fmt.Println(string(message.Body))
	//
	//message.Ack(true)

	////todo 也就是说这些通道里面全部都是 struct
	////todo 重要  重要  他这个是只能去loop一条消息  就是循环地去找   这玩意真傻 走着走着 就卡住了  只要有值  就不会堵在这里？
	//for  msg := range delivery{
	//
	//	fmt.Println(string(msg.Body))
	//
	//}
	//}

	//fmt.Println("1"+string(aa.Body))
	//就是消费它当前的消息

	//这个消息一旦确认之后  这个消息就从队列中删除了
	//aa.Ack(false)

	//没确认的话 你一直可以重复消费

	//fmt.Println(delivery)

	//	TODO 将剩余库存回滚

}



//15min钟之后  一定会收到这条消息的
//生产者发一个orderNo

//consumer  : select activityId,status from activityTable where orderNo

//再根据activityId去revert 库存

//这个时候要根据订单的orderNo  查看订单的状态    如果 订单的状态为  status != 2 {未付款 你就把那个avilable_stock=avilable_stock+1 (where activity_id=?),  }
