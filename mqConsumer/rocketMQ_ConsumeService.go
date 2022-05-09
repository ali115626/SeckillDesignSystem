package mqConsumer

import (
	"SeckillDesign/utitl"
	"errors"
	"fmt"
	"github.com/streadway/amqp"
)
//todo 这个就是一些公共的函数
//   还是写死了      它只能消费创建订单的信息  这个应该将其弄成common的吧

//一个common的消费者
//	queueName :="orderMessage"
func ReceiveMessageNormalConsumer(queueName string) string {
	//TODO  要将这个改成一个函数吧
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		errResult := fmt.Sprintf("connect to the rabbitMq failed %s", err)
		fmt.Println(errResult)
		err = errors.New(errResult)
		//return err
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		errResult := fmt.Sprintf("Failed to open a channel %s", err)
		fmt.Println(errResult)

		err = errors.New(errResult)
		//return err
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		queueName, // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		errResult := fmt.Sprintf("Failed to declare a queue %s", err)
		fmt.Println(errResult)

		err = errors.New(errResult)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",    // consumer true,   // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		false,
		nil,
	)
	d := <-msgs
	return string(d.Body)
}

func MqConsumerCommon(queueName string) interface{} {
	conn := utitl.NewRabbitMQ().Conn
	defer conn.Close()
	ch := utitl.NewRabbitMQ().Ch
	defer ch.Close()
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		errResult := fmt.Sprintf("Failed to declare a queue %s", err)
		fmt.Println(errResult)

		err = errors.New(errResult)
	}
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		false,
		nil,
	)
	if err != nil {
		return err
	}
	return msgs
}

func DeadQueueConsumer(queueName string) <-chan amqp.Delivery{
	conn := utitl.NewRabbitMQ().Conn
	defer conn.Close()
	ch := utitl.NewRabbitMQ().Ch
	defer ch.Close()
	delivery, err := ch.Consume(queueName, "", false, false, false, false, nil)
	if err != nil {
		fmt.Println("consume error,err=", err)
	}
	return delivery
}

