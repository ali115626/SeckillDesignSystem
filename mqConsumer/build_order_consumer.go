package mqConsumer

import (
	"SeckillDesign/SellerService"
	"SeckillDesign/constant"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/streadway/amqp"
)

//func CreateOrder(){
////	todo 消费MQ
//	message :=consumer.ReceiveMessage()
////	TODO 这个你要unMarshal一下吗？
//	fmt.Println(message)
////	TODO 把订单信息insert到数据库
//}

//func BuildOrder(w http.ResponseWriter, r *http.Request){
//	//	把activityId 和 userID 给传进来
//	//	把order 写入Mq 中
//	err := r.ParseForm()
//	if err != nil {
//		return
//	}
//	requestMap := r.Form
//	activityId := requestMap["activityId"][0]
//	//TODO 这里面加一个用户是不是之前已经购买过的代码逻辑
//	//fmt.Println("userID=",commidtyId)
//	//userId其实是这样给传过来的 不是从请求中获取得到的
//	userId := requestMap["userId"][0]

//	todo 消费MQ

//todo  从messageQUeue中读取订单的信息

//todo  这些代码中的错误 你是需要打印到 log里面的

func BuildOrderConsumer() error {

	queueName := "orderMessage"
	delivery := MqConsumerCommon(queueName)
	deliveryMsg := delivery.(<-chan amqp.Delivery)
	for msg := range deliveryMsg {
		//TODO 把这些逻辑抽成一个函数吧
		message := string(msg.Body)
		//	再去unMarshal一下  unmarshal 到 orderInfo中
		orderInfo := constant.OrderInfo{}
		err := json.Unmarshal([]byte(message), orderInfo)
		//	TODO 这个你要unMarshal一下吗？
		//todo  检查一下  库存里面有没有
		//TODO 有的话  给前端返回1 ：创建订单成功  2：已经支付了（支付成功了）

		//	TODO 把订单信息insert到数据库
		if err != nil {
			return errors.New(fmt.Sprintf("json unmarshal orderInfo err,err=", err))
		}

		//-----------————————————————buildOrder()——————————————————————————————
		err = SellerService.BuildOrderProcess(orderInfo)
		if err == nil {
			//如果没有出现错误的话  业务逻辑可以正常走下去的话  就手动  确认  ack
			err = msg.Ack(false)
			if err != nil {
				return err
			}
			//return nil
			//return err
		} else {
			//	如果出现问题的话  就不要去确认这条消息
			err = msg.Reject(true)
			if err != nil {
				return err
			}
			//msg.Nack
		}

	}
	return nil


}