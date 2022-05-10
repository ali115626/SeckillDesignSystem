package producer

import (
	"errors"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

//TODO  这里面  你再去封装一下  队列名再改一下
//todo  到时候就调用一下这个函数吧
func SendMQCommon(body []byte, queueName string) error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		errResult := fmt.Sprintf("connect to the rabbitMq failed %s", err)
		fmt.Println(errResult)
		err = errors.New(errResult)
		return err
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		errResult := fmt.Sprintf("Failed to open a channel %s", err)
		fmt.Println(errResult)
		err = errors.New(errResult)
		return err
	}
	//failOnError(err, "Failed to open a channel")
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
		return err
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			//ContentType: "text/plain",
			Body: body,
		})
	if err !=nil{
		log.Printf("ch.Publish error!!!!!!")

	}
	log.Printf(" [x] Sent %s", body)
	//failOnError(err, "Failed to publish a message")

	if err != nil {
		errResult := fmt.Sprintf("Failed to publish a message %s", err)
		fmt.Println(errResult)

		err = errors.New(errResult)
		return err
	}
	return nil

}
