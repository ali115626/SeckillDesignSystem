package SellerService

import (
	"SeckillDesign/constant"
	Dao "SeckillDesign/dao/mysql"
	"SeckillDesign/producer"
	"encoding/json"
	"fmt"
	"net/http"
)



func ProcessPayDoneOrder(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	requestMap := r.Form
	isPay := requestMap["isPay"][0]
	orderNo := requestMap["orderNo"][0]
	if isPay == "0" {
		fmt.Fprintf(w, "订单未付款")
		return
	}
	err=Dao.ChangeOrderTablePayStatus(orderNo)
	if err !=nil{
		fmt.Println("修改订单状态失败，err=",err)
	}
	//todo 这个MQ要分出来是库存系统吗
	//orderInfo := &constant.OrderInfo{}
	var orderInfo *constant.OrderInfo


	orderInfo,err=Dao.QueryDetailsFromOrderInfoTable(orderNo)
	if err != nil{
		fmt.Println("query Detail info FromOrderInfoTable error,err=",err)
		return //这里还是return吧  因为你如果  没有得到  OrderInfoTable的detail信息 走下去也没什么用
	}
	//todo  你这里慎用return   你这里一return   就会导致出错之后 后面的代码就不能跑了   这个完全要看   你能否忍受这个错误

	message := *orderInfo

	messageResult, err := json.Marshal(message)
	queueName := "payDone"
	//TODO 这个确实是可以同步调用吧
	err = producer.SendMQCommon(messageResult, queueName)
	if err != nil {
		fmt.Println("sendMQ error,err=", err)
	}
}
