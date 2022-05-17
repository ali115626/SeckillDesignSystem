package mqConsumerService

import (
	//"SeckillDesign/SellerService/DownLogic"
	"SeckillDesign/constant"
	"SeckillDesign/mqConsumerService/DownLogic"
	"SeckillDesign/utitl"
	"fmt"
)

func PayStatusCheckListener() {

	queueName := constant.DeadQueue

	qName := DeadQueueConsumer(queueName)
	delivery, err := utitl.NewRabbitMQ().Ch.Consume(qName, "", false, false, false, false, nil)

	if err != nil {
		fmt.Println("consume error,err=", err)
	}
	//TODO 这里面缺一个 那个版本控制
	for msgs := range delivery {

		err := DownLogic.DoRevertStockLogic(msgs)
		//DorevertRedis(msgs)

		if err == nil {
			msgs.Ack(false) //业务逻辑没问题的话  就手动确认
		} else {
			msgs.Reject(true)
		}

	}

}