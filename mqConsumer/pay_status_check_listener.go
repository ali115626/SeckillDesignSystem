package mqConsumer

import (
	"SeckillDesign/SellerService/DownLogic"
	"SeckillDesign/constant"
)

func PayStatusCheckListener() {
	queueName := constant.DeadQueue
	delivery := DeadQueueConsumer(queueName)
	for msgs := range delivery {
		err := DownLogic.DoRevertStockLogic(msgs)
		if err == nil {
			msgs.Ack(false) //业务逻辑没问题的话  就手动确认
		} else {
			msgs.Reject(true)
		}
	}

}