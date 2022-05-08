package producer

import (
	"SeckillDesign/constant"
	"SeckillDesign/utitl"
	"errors"
	"fmt"
	"github.com/streadway/amqp"
)

//
//func CheckErr(err error){
//	if err != nil{
//		errStr:=fmt.Sprintf("declare NormalQueue error,err=",err)
//		return errors.New(errStr)
//	}
//}

func DelayProducer(message string) error {

	//TODO 他这个不需要你传入什么  直接生成一个链接
	//就是获取这个连接呀
	conn := utitl.NewRabbitMQ().Conn

	defer conn.Close()

	mqCh := utitl.NewRabbitMQ().Ch

	defer mqCh.Close()

	var err error

	_, err = mqCh.QueueDeclare(constant.NormalQueue, true, false, false, false, amqp.Table{
		"x-message-ttl":             2,                       //20000,//todo  你这个是声明队列的过期时间还是消息的过期时间   这个时间是可以改的吧
		"x-dead-letter-exchange":    constant.DeadExchange,   //声明死信交换机
		"x-dead-letter-routing-key": constant.DeadRoutingKey, //声明死信队列 -----这个原因吗？
	})
	if err != nil {
		errStr := fmt.Sprintf("declare NormalQueue error,err=", err)
		return errors.New(errStr)
	}

	err = mqCh.ExchangeDeclare(constant.NormalExchange, amqp.ExchangeDirect, true, false, false, false, nil)
	if err != nil {
		errStr := fmt.Sprintf("NormalExchange exchange declare false,err=", err)
		return errors.New(errStr)
	}

	err = mqCh.QueueBind(constant.NormalQueue, constant.NormalRoutingKey, constant.NormalExchange, false, nil)
	if err != nil {
		err = errors.New(fmt.Sprintf("normal：队列、交换机、routing-key 绑定失败", err))
		fmt.Println("normal：队列、交换机、routing-key 绑定失败", err)
		return err

	}
	//	声明死信队列
	_, err = mqCh.QueueDeclare(constant.DeadQueue, true, false, false, false, nil)
	if err != nil {

		err = errors.New(fmt.Sprintf("deadQueue声明失败,err=", err))
		return err

	}

	err = mqCh.ExchangeDeclare(constant.DeadExchange, amqp.ExchangeDirect, true, false, false, false, nil)
	if err != nil {
		err = errors.New(fmt.Sprintf("deadExchange声明失败,err=", err))
		return err
	}
	err = mqCh.QueueBind(constant.DeadQueue, constant.DeadRoutingKey, constant.DeadExchange, false, nil)
	if err != nil {
		err = errors.New(fmt.Sprintf("死信交换机和死信队列绑定失败，err=", err))
		return err
	}

	err = mqCh.Publish(constant.NormalExchange, constant.NormalRoutingKey, false, false, amqp.Publishing{ContentType: "text/plain",
		Body: []byte(message)})

	if err != nil {
		return errors.New(fmt.Sprintf("消息发布失败，err=", err))
	}
	return nil
}
