package mqConsumer

import (
	Dao "SeckillDesign/dao/mysql"
	"encoding/json"
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

//todo 这个consumer一直会在监听rabbitMQ

func OrderDeductConsumer() {

	queueName := "payDone"
	delivery := MqConsumerCommon(queueName)
	orderInfo := &OrderInfo{}
	deliveryMsg := delivery.(<-chan amqp.Delivery)
	//todo  没关系的  你的函数其实一直阻塞在这里
	//TODO 把这个程序开始在main函数中就运行出来   初始化 就一直等着呗

	for msg := range deliveryMsg {

		//TODO 把这些逻辑抽成一个函数吧
		message := string(msg.Body)
		//	再去unMarshal一下  unmarshal 到 orderInfo中
		err := json.Unmarshal([]byte(message), orderInfo)
		if err != nil {
			fmt.Println("json unmarshal orderInfo err,err=", err)
		}
		//这里不用再search orderinfo的那张表格了   你不是之前已经改变了吗？
		//if orderInfo.Status != 2{
		//
		//}
		activityId := orderInfo.ActivityId
		//todo  这里就是扣减库存成功
		//整成并发？

		err =Dao.DeductStock(activityId)

		if err!=nil{
			fmt.Println("Deduct Stock error,err=",err)
		}

	}

}
