package SellerService

import (
	"SeckillDesign/constant"
	"SeckillDesign/producer"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

//type OrderInfo struct {
//	OrderId      string
//	UserId       string
//	ActivityName string
//	OrderPrice   string
//	ActivityId   string
//	Status       int
//}

func ProcessPayDoneOrder(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	requestMap := r.Form
	//如果isPay=0 就直接返回
	isPay := requestMap["isPay"][0]

	orderNo := requestMap["orderNo"][0]

	if isPay == "0" {
		fmt.Fprintf(w, "订单未付款")
		return
	}

	db, err := sql.Open("mysql", "root:123456@/seckill_scheme?charset=utf8")
	if err != nil {
		fmt.Println("open database error,err=", err)
	}
	//orderNo

	//select * from OrderInfoTable;
	//	+---------------------+--------+-------------------+------------+------------+--------+---------------------+
	//	| orderId             | userId | activityName      | orderPrice | activityId | status | createAt            |

	//todo status=2 代表付款成功了
	_, err = db.Exec("update OrderInfoTable set status=2 WHERE orderId=?", orderNo)
	if err != nil {
		fmt.Println("exec failed, err=", err)
		//	todo 我就奇怪了   万一我在这边一直点下去怎么办   代码层校验是否重复上传   怎么办呢
	}

	//todo 这里面应该还要search一下数据库

	orderinfo := constant.OrderInfo{}

	//body []byte,queueName string
	var userId string
	var activityName string
	var orderPrice string
	var status string
	var createAt string
	var activityId string
	err = db.QueryRow("SELECT userId,activityId,activityName,orderPrice,status,createAt FROM OrderInfoTable WHERE orderId=?", orderNo).Scan(&userId, &activityId, &activityName, &orderPrice, &status, &createAt)
	//发现即使db.QueryRow(）这里面ELECT delete_status FROM blog_info WHERE title_id SQL语句出问题了  也不会报错
	if err != nil {
		fmt.Println(err)
		fmt.Println("select  paper_content  error")
		return
	}
	orderinfo.OrderId = orderNo
	orderinfo.OrderPrice = orderPrice
	statusInt, err := strconv.Atoi(status)
	orderinfo.Status = statusInt
	orderinfo.ActivityId = activityId
	message := orderinfo
	messageResult, err := json.Marshal(message)

	queueName := "payDone"

	err = producer.SendMQCommon(messageResult, queueName)
	if err != nil {
		fmt.Println("sendMQ error,err=", err)
	}

	//	TODO  这里面直接写MQ 让那边 减库存   subtract 库存

}
