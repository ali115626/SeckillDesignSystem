package mqConsumer

import (
	"SeckillDesign/SellerService/DownLogic"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/streadway/amqp"
)

//TODO 这里面加一个用户是不是之前已经购买过的代码逻辑
//todo  从messageQueue中读取订单的信息
//todo  这些代码中的错误 你是需要打印到 log里面的

func BuildOrderConsumer() error {
	queueName := "orderMessage"
	delivery := MqConsumerCommon(queueName)
	deliveryMsg := delivery.(<-chan amqp.Delivery)
	for msg := range deliveryMsg {
		//TODO 把这些逻辑抽成一个函数吧
		//-----------————————————————buildOrder()——————————————————————————————
		err := DownLogic.BuildOrderProcessService(msg)
		if err == nil {
			//如果没有出现错误的话  业务逻辑可以正常走下去的话  就手动  确认  ack
			err = msg.Ack(false)
			if err != nil {
				return err
			}
		} else {
			err = msg.Reject(true)
			if err != nil {
				return err
			}
		}
	}
	return nil
}