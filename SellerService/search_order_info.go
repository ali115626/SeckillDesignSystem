package SellerService

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

//TODO  返回的订单信息   type orderinfo struct  这样写费我时间

type OrderInfo struct {
	UserId       string `json:"user_id"`
	ActivityName string `json:"activity_name"`
	OrderPrice   string `json:"order_price"`
	Status       string `json:"status"`
	CreateAt     string `json:"create_at" `
}

func SearchOrderInfo(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		return
	}
	requestMap := r.Form
	orderId := requestMap["orderId"][0]
	db, err := sql.Open("mysql", "root:123456@/seckill_scheme?charset=utf8")
	if err != nil {
		fmt.Println("open database error,err=", err)
	}
	var userId string
	var activityName string
	var orderPrice string
	var status string
	var createAt string

	err = db.QueryRow("SELECT userId,activityName,orderPrice,status,createAt FROM OrderInfoTable WHERE orderId=?", orderId).Scan(&userId, &activityName, &orderPrice, &status, &createAt)
	//发现即使db.QueryRow(）这里面ELECT delete_status FROM blog_info WHERE title_id SQL语句出问题了  也不会报错
	if err != nil {
		fmt.Println(err)
		fmt.Println("select  paper_content  error")
		return
	}
	if status == "1" {
		status = "未付款"
	} else if status == "0" {
		status = "订单出错，创建订单失败"
	}
	OrderInfoResult := OrderInfo{}
	OrderInfoResult.UserId = userId
	OrderInfoResult.OrderPrice = orderPrice
	OrderInfoResult.Status = status
	OrderInfoResult.CreateAt = createAt
	//OrderInfoResult.
	marshalResult, err := json.Marshal(OrderInfoResult)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("marshalResult", string(marshalResult))
	fmt.Fprintf(w, string(marshalResult))
	//TODO 将查询到结果返回
}
