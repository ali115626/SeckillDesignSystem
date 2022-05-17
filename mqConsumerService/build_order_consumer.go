package mqConsumerService

import (
	"SeckillDesign/mqConsumerService/DownLogic"
	//"SeckillDesign/SellerService/DownLogic"
	"SeckillDesign/utitl"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//TODO 这里面加一个用户是不是之前已经购买过的代码逻辑
//todo  从messageQueue中读取订单的信息
//todo  这些代码中的错误 你是需要打印到 log里面的
// error
func BuildOrderConsumer(){
	queueName := "orderMessage"
	//var delivery <-chan amqp.Delivery
	//queueName := MqConsumerCommon(queueName)

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
	if err != nil{
		fmt.Println(fmt.Sprintf("NewRabbitMQ().Ch.Consume",err))
		return
		//return errors.New(fmt.Sprintf("NewRabbitMQ().Ch.Consume",err))
	}
	for msg := range msgs {
		//TODO 把这些逻辑抽成一个函数吧
		fmt.Println("msg Body=-----",msg.Body)
		//-----------————————————————buildOrder()——————————————————————————————
		err :=DownLogic.BuildOrderProcessService(msg)
		if err == nil {
			//如果没有出现错误的话  业务逻辑可以正常走下去的话  就手动  确认  ack
			//fmt.Println("00000000000000000099990909099000000000000000009------------")

			err = msg.Ack(false)
			if err != nil {
				return
				//return err
			}
		} else {
			fmt.Println("11111111199990909099000000000000000009------------")
			err = msg.Reject(true)
			if err != nil {
				return
				//return err
			}
		}
	}
	return
}




//////////-_________——————————————————————————————————————————————————————————-
//	conn := utitl.NewRabbitMQ().Conn
//	defer conn.Close()
//	ch := utitl.NewRabbitMQ().Ch
//	defer ch.Close()
//	q, err := ch.QueueDeclare(
//		queueName, // name
//		false,     // durable
//		false,     // delete when unused
//		false,     // exclusive
//		false,     // no-wait
//		nil,       // arguments
//	)
//	if err != nil {
//		errResult := fmt.Sprintf("Failed to declare a queue %s", err)
//		fmt.Println(errResult)
//
//		err = errors.New(errResult)
//	}
//	msgs, err := ch.Consume(
//		q.Name, // queue
//		"",     // consumer true,   // auto-ack
//		false,  // exclusive
//		false,  // no-local
//		false,  // no-wait
//		false,
//		nil,
//	)
//	if err != nil {
//		return nil
//	}
//////////-_________——————————————————————————————————————————————————————————-