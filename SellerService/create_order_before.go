package SellerService

import (
	"SeckillDesign/producer"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/godruoyi/go-snowflake"
	"net/http"
	"strconv"
)

type Order struct {
	OrderId      string
	UserId       string
	ActivityName string
	OrderPrice   string
	ActivityId   string
}

func (this *Order) SetOrderId(orderId string) {
	this.OrderId = orderId
}

func (this *Order) SetUserId(userId string) {
	this.UserId = userId
}

func (this *Order) SetActivityName(activityName string) {
	this.ActivityName = activityName
}

func (this *Order) SetOrderPrice(commodityPrice string) {
	this.OrderPrice = commodityPrice
}

func (this *Order) SetActivityId(activityId string) {
	this.ActivityId = activityId
}

//秒杀成功之后  查看redis的zset里面有没有这个UserID  如果没有的话  直接return  如果有的话 才会进行下面的步骤

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
	//fmt.Println("userID=",commidtyId)
	//userId其实是这样给传过来的 不是从请求中获取得到的
	userId := requestMap["userId"][0]
	//	TODO  根据activityId去search一下 activityName 和活动的商品的价格
	db, err := sql.Open("mysql", "root:123456@/seckill_scheme?charset=utf8")
	if err != nil {
		fmt.Println("open database error,err=", err)
	}
	var activityName string

	var price string

	//todo 这个时候去search这个表格没有关系  因为商品活动的信息你已经  预热到 redis里面了  但是先去数据库里面查吧
	err = db.QueryRow("SELECT activityName,price FROM ActivityTable WHERE activityId=?", activityId).Scan(&activityName, &price)
	//发现即使db.QueryRow(）这里面ELECT delete_status FROM blog_info WHERE title_id SQL语句出问题了  也不会报错
	if err != nil {
		fmt.Println(err)
		fmt.Println("select  paper_content  error")
		return

	}

	order := Order{}
	//todo  订单编号是用雪花算法生成的

	snowflake.SetMachineID(1)

	// Or set private ip to machineid, testing...
	// snowflake.SetMachineID(snowflake.PrivateIPToMachineID())

	id := snowflake.ID()
	snowflakeId := strconv.FormatUint(id, 10)

	//fmt.Println("The length of snowflakeId,snowflakeId",len(snowflakeId))

	//id:=strconv.FormatUint(id,10)

	order.SetOrderId(snowflakeId)

	order.SetOrderPrice(price)
	order.SetUserId(userId)
	order.SetActivityName(activityName)
	order.SetActivityId(activityId)

	//todo  将这些整合到MQ中
	//marshal一下
	fmt.Println(order)

	orderByte, err := json.Marshal(order)
	err = producer.SendMQ(orderByte)
	if err != nil {
		fmt.Println("message send to rabbitMQ error", err)
	}

	fmt.Printf("%+v", order)

	//	TODO 将订单信息写入MQ  发送延时队列的信息

	//	给端上返回：正在生成订单中.......   这个之后 再去跳转到生成订单的那个页面

	orderInfo, _ := json.Marshal(order)

	fmt.Fprintf(w, string(orderInfo))

	//这里面放上一个  加延时队列的信息

	err = producer.DelayProducer(order.OrderId)

	if err != nil {
		//	发送延时队列失败
		//	不过你一直发送   是panic了 还是 retry
		fmt.Println("发送延时消息失败，err=", err)
	}

	//TODO 1先暴力写一下吧 从Mq中取出订单的信息    数据库生成订单信息的表格   订单的状态  status  :已支付  未支付  ---》 锁定库存

	//TODO 2.将mysql中的库存数更正一下   15分钟后就是查看一下   回复一下那个stocks

}
