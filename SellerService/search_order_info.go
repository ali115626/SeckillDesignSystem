package SellerService

import (
	Dao "SeckillDesign/dao/mysql"
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
	OrderInfoResult,err :=Dao.ShowOrderInfo(orderId)
	if err != nil{
		fmt.Println("Show OrderInfo error,err=",err)//正常的话应该打印到log日志文件中吧
		return
	}
	marshalResult, err := json.Marshal(OrderInfoResult)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("marshalResult", string(marshalResult))
	fmt.Fprintf(w, string(marshalResult))
	//TODO 将查询到结果返回
}
