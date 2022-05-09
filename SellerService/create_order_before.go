package SellerService

import (
	Dao "SeckillDesign/dao/mysql"
	"SeckillDesign/producer"
	"SeckillDesign/utitl"
	"encoding/json"
	"fmt"
	"github.com/godruoyi/go-snowflake"
	"net/http"
	"strconv"
)

//todo 秒杀成功之后  查看redis的zset里面有没有这个UserID  如果没有的话  直接return  如果有的话 才会进行下面的步骤

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	//	把activityId 和 userID 给传进来
	//	把order 写入Mq 中
	err := r.ParseForm()
	if err != nil {
		return
	}
	requestMap := r.Form
	activityId := requestMap["activityId"][0]
	//TODO 这里面加一个用户是不是之前已经购买过的代码逻辑
	userId := requestMap["userId"][0]
	//input:activityId、
	//output:activityName,price
	var activityName string
	var price string
	activityName,price,err=Dao.QueryPriceActivityTable(activityId)
	if err!=nil{
		fmt.Println("Query Price from ActivityTable error,err=",err)
	}
	//order := Order{}
	//  订单编号是用雪花算法生成的
	snowflake.SetMachineID(1)
	id := snowflake.ID()
	snowflakeId := strconv.FormatUint(id, 10)
	order := utitl.FillOrderInfo(snowflakeId,price,userId,activityName,activityId)
	orderByte, err := json.Marshal(order)
	err = producer.SendMQ(orderByte)
	if err != nil {
		fmt.Println("message send to rabbitMQ error", err)
	//	TODO  发送失败  好家伙 你就在这里return了
	//TODO 千万别return  不然你后面都不走了
	}
	err = producer.DelayProducer(order.OrderId)
	if err != nil {
		//todo 发送失败就发送失败呗  ！
		fmt.Println("发送延时消息失败，err=", err)
	}
	fmt.Fprintf(w, string(orderByte))

}


//TODO 1先暴力写一下吧 从Mq中取出订单的信息    数据库生成订单信息的表格   订单的状态  status  :已支付  未支付  ---》 锁定库存
//TODO 2.将mysql中的库存数更正一下   15分钟后就是查看一下   回复一下那个stocks

//	发送延时队列失败
//	不过你一直发送   是panic了 还是 retry

//这里面放上一个  加延时队列的信息
//	TODO 将订单信息写入MQ  发送延时队列的信息
//	给端上返回：正在生成订单中.......   这个之后 再去跳转到生成订单的那个页面
//todo  将这些整合到MQ中
//发现即使db.QueryRow(）这里面ELECT delete_status FROM blog_info WHERE title_id SQL语句出问题了  也不会报错

