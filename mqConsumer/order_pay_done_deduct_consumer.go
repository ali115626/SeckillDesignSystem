package mqConsumer

import (
	"SeckillDesign/SellerService/DownLogic"
	"SeckillDesign/utitl"
	"fmt"
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
	qName := MqConsumerCommonNew(queueName)
	msgs, err := utitl.NewRabbitMQ().Ch.Consume(
		qName, // queue
		"",     // consumer true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		false,
		nil,
	)
	if err != nil {
		return
	}
	//deliveryMsg := delivery.(<-chan amqp.Delivery)
	for msg := range msgs {
		//做业务逻辑的时候 这个里面可不可以  用一下那个 go程  这个消息不小心拒收  你总不能阻塞而不能处理其他消息吧
		err :=DownLogic.DoPayDoneDeductLogic(msg)
		if err != nil{
			fmt.Println(err)
			err=msg.Reject(true)
			//如果主要的业务逻辑出现错误  那你的  几乎一直会reject   太吓人了  还是不要这么做了
			if err !=nil{
				fmt.Println("reject error,err=",err)
			}
		} else{
			err=msg.Ack(false)
			if err!=nil{
				fmt.Println("ack error,err=",err)
			}
		}
	}
}


//todo  没关系的  你的函数其实一直阻塞在这里
//TODO 把这个程序开始在main函数中就运行出来   初始化 就一直等着呗
//todo 这个consumer一直会在监听rabbitMQ
