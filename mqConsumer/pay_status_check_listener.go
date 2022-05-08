package mqConsumer

import (
	"SeckillDesign/SellerService"
	"SeckillDesign/constant"
	"SeckillDesign/utitl"
	"database/sql"
	"fmt"
)

func PayStatusCheckListener() {
	//	todo 消费延时队列的消息
	conn := utitl.NewRabbitMQ().Conn
	defer conn.Close()
	ch := utitl.NewRabbitMQ().Ch
	defer ch.Close()
	//for {
	delivery, err := ch.Consume(constant.DeadQueue, "", false, false, false, false, nil)
	if err != nil {
		fmt.Println("consume error,err=", err)
	}

	forever := make(chan bool)

	db, err := sql.Open("mysql", "root:123456@/seckill_scheme?charset=utf8")
	if err != nil {
		fmt.Println("open database error,err=", err)
	}

	go func() {
		for msgs := range delivery {
			//15min钟之后  一定会收到这条消息的
			//生产者发一个orderNo

			//consumer  : select activityId,status from activityTable where orderNo

			//再根据activityId去revert 库存

			//这个时候要根据订单的orderNo  查看订单的状态    如果 订单的状态为  status != 2 {未付款 你就把那个avilable_stock=avilable_stock+1 (where activity_id=?),  }
			orderNo := string(msgs.Body)
			//	再去unMarshal一下  unmarshal 到 orderInfo中

			//activityId | status

			var status interface{}
			var activityId string
			err = db.QueryRow("SELECT activityId,status FROM OrderInfoTable WHERE orderId=?", orderNo).Scan(&activityId, &status)
			//发现即使db.QueryRow(）这里面ELECT delete_status FROM blog_info WHERE title_id SQL语句出问题了  也不会报错
			if err != nil {
				fmt.Println(err)
				fmt.Println("select  paper_content  error")
				return
			}
			statusStr, ok := status.(string)
			if !ok {
				fmt.Println("status断言失败！")
			}
			if statusStr != "2" {
				//	订单未支付
				//	avilable_stock=avilable_stock+1 (where activity_id=?)
				//	MySQL

				_, err = db.Exec("update ActivityTable set available_stock=available_stock+1,locked_stock=locked_stock-1 WHERE activityId=?", activityId)
				if err != nil {
					fmt.Println("exec failed, err=", err)
				}
				//	Redis

				//	从限购名单中删除

			}

			//fmt.Println(string(msgs.Body))
			//从这个message中检查出来 orderNo

			//根据这个orderNo 去search数据库   查看这个订单是否付款的status

			//if status=2{
			//如果没有付款    就根据 orderNo去 revert database
			//
			// revert redis  } else{
			//		 }

			err := SellerService.DoRevertStock()

			if err != nil {

			} else {
				msgs.Ack(false)

			}

			//TODO 他应该是在这里面去写业务逻辑  revertStock   如果业务逻辑处理成功  就手动去确认 失败  就 返回 NACK  但是现在还是不知道
			//TODO 如何去 合理地  不断监听消息  缺一个 if {} else{}

			//todo   先看看它能不能work
			//TODO 手动确认当前消息
		}

	}()
	<-forever
	//
	//
	//
	//message :=<-delivery
	//fmt.Println(string(message.Body))
	//
	//message.Ack(true)

	////todo 也就是说这些通道里面全部都是 struct
	////todo 重要  重要  他这个是只能去loop一条消息  就是循环地去找   这玩意真傻 走着走着 就卡住了  只要有值  就不会堵在这里？
	//for  msg := range delivery{
	//
	//	fmt.Println(string(msg.Body))
	//
	//}
	//}

	//fmt.Println("1"+string(aa.Body))
	//就是消费它当前的消息

	//这个消息一旦确认之后  这个消息就从队列中删除了
	//aa.Ack(false)

	//没确认的话 你一直可以重复消费

	//fmt.Println(delivery)

	//	TODO 将剩余库存回滚

}
