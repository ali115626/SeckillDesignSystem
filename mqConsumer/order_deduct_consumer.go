package mqConsumer

import (
	"database/sql"
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
	db, err := sql.Open("mysql", "root:123456@/seckill_scheme?charset=utf8")
	if err != nil {
		fmt.Println("open database error,err=", err)
	}

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

		//_, err = db.Exec("insert into CommodityTable(commodityName,price,describeInfo,updateAt) values(?,?,?,?)", commodityName, price, describeInfo, updateAt)
		//if err != nil {
		//	fmt.Println("exec failed, err=", err)
		//	return
		//}
		//fmt.Fprintf(w, "商品上传成功！")
		//todo  这里就是库建库存成功
		_, err = db.Exec("update ActivityTable set locked_stock=locked_stock-1 WHERE activityId=?", activityId)
		if err != nil {
			fmt.Println("exec failed, err=", err)
		}

	}
}
