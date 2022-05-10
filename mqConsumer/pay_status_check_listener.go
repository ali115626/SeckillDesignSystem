package mqConsumer

import (
	"SeckillDesign/SellerService/DownLogic"
	"SeckillDesign/constant"
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
	for msgs := range delivery {
		err := DownLogic.DoRevertStockLogic(msgs)
		if err == nil {
			msgs.Ack(false) //业务逻辑没问题的话  就手动确认
		} else {
			msgs.Reject(true)
		}
	}

}