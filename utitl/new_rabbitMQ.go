package utitl

import (
	"errors"
	"fmt"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Conn *amqp.Connection
	Ch   *amqp.Channel
}

func NewRabbitMQ() *RabbitMQ {
	//todo 因为这个RabbitMQ只是一个类型，你需要将RabbitMQ实例化呀
	//rabbitMQ := RabbitMQ{}
	//var err error
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		errResult := fmt.Sprintf("connect to the rabbitMq failed %s", err)
		fmt.Println(errResult)
		err = errors.New(errResult)
		return nil
	}
	Channel, err := conn.Channel()
	return &RabbitMQ{
		Conn: conn,
		Ch:   Channel,
	}

}
