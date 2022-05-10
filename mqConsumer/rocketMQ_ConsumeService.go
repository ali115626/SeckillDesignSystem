package mqConsumer

import (
	"SeckillDesign/constant"
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

func MqConsumerCommon(queueName string) <-chan amqp.Delivery{
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

	//return
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
		return nil
	}
	return msgs
}


//最好把这个什么队列的声明放到consumer里面吧

func DeadQueueConsumer(queueName string) string{
	conn := utitl.NewRabbitMQ().Conn
	defer conn.Close()
	mqCh := utitl.NewRabbitMQ().Ch
	defer mqCh.Close()
	//ch := utitl.NewRabbitMQ().Ch
	//defer ch.Close()

	var err error
	_, err = mqCh.QueueDeclare(constant.NormalQueue, true, false, false, false, amqp.Table{
		"x-message-ttl":             120000,                       //1000*60*4,//todo  你这个是声明队列的过期时间还是消息的过期时间   这个时间是可以改的吧
		"x-dead-letter-exchange":    constant.DeadExchange,   //声明死信交换机
		"x-dead-letter-routing-key": constant.DeadRoutingKey, //声明死信队列 -----这个原因吗？
	})
	if err != nil {
		errStr := fmt.Sprintf("declare NormalQueue error,err=", err)
		fmt.Println(errStr)
		//return errors.New(errStr)
	}
	err = mqCh.ExchangeDeclare(constant.NormalExchange, amqp.ExchangeDirect, true, false, false, false, nil)
	if err != nil {
		errStr := fmt.Sprintf("NormalExchange exchange declare false,err=", err)
		//return errors.New(errStr)
		fmt.Println(errStr)
	}
	err = mqCh.QueueBind(constant.NormalQueue, constant.NormalRoutingKey, constant.NormalExchange, false, nil)
	if err != nil {
		err = errors.New(fmt.Sprintf("normal：队列、交换机、routing-key 绑定失败", err))
		fmt.Println("normal：队列、交换机、routing-key 绑定失败", err)
		//return err
	}
	//声明死信队列
	_, err = mqCh.QueueDeclare(constant.DeadQueue, true, false, false, false, nil)
	if err != nil {
		err = errors.New(fmt.Sprintf("deadQueue声明失败,err=", err))
		fmt.Println(err)
		//return err
	}
	err = mqCh.ExchangeDeclare(constant.DeadExchange, amqp.ExchangeDirect, true, false, false, false, nil)
	if err != nil {
		err = errors.New(fmt.Sprintf("deadExchange声明失败,err=", err))
		fmt.Println(err)
		//return err
	}
	err = mqCh.QueueBind(constant.DeadQueue, constant.DeadRoutingKey, constant.DeadExchange, false, nil)
	if err != nil {
		err = errors.New(fmt.Sprintf("死信交换机和死信队列绑定失败，err=", err))
		fmt.Println(err)
		//return err
	}
	return queueName

	//return delivery
}




func MqConsumerCommonNew(queueName string) string{
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

	return q.Name

}

