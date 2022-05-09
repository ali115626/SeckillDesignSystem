package mqConsumer

import (
	"SeckillDesign/SellerService"
	"fmt"
	"github.com/streadway/amqp"
)

type OrderInfo struct {
	OrderId      string
	UserId       string
	ActivityName string
	OrderPrice   string
	ActivityId   string
	Status       int
}


func OrderDeductConsumer() {
	queueName := "payDone"
	delivery := MqConsumerCommon(queueName)
	deliveryMsg := delivery.(<-chan amqp.Delivery)
	for msg := range deliveryMsg {
		err :=SellerService.DoPayDoneDeductLogic(msg)
		if err !=nil{
			err=msg.Reject(true)
			if err !=nil{
				fmt.Println("reject error,err=",err)
			}
		} else{
			err=msg.Ack(true)
			if err!=nil{
				fmt.Println("ack error,err=",err)
			}
		}
	}
}


//todo  没关系的  你的函数其实一直阻塞在这里
//TODO 把这个程序开始在main函数中就运行出来   初始化 就一直等着呗
//todo 这个consumer一直会在监听rabbitMQ
