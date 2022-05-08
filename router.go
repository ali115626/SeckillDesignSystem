package main

import (
	"SeckillDesign/SellerService"
	"SeckillDesign/mqConsumer"

	//"SeckillDesign/orderSystem"
	"fmt"
	"net/http"
)

//func TestStringToTime(t *testing.T) {
//	str := "2021-01-03 15:23:11"
//	// 设置时区
//	loc, _ := time.LoadLocation("Asia/Shanghai")
//	d, _ := time.ParseInLocation("2006-01-02 15:04:05", str, loc)
//	fmt.Printf("time: %v\n", d)
//}

func main() {
	err := mqConsumer.BuildOrderConsumer()
	if err != nil {
		fmt.Println(err)
	}

	mqConsumer.PayStatusCheckListener()

	http.HandleFunc("/ReleaseActivity", SellerService.ReleaseActivity)

	//UploadCommodity
	http.HandleFunc("/UploadCommodity", SellerService.UploadCommodity)
	//PullActivityInfo

	http.HandleFunc("/PullActivityInfo", SellerService.PullActivityInfo)

	//GetCommodityInfo
	http.HandleFunc("/GetCommodityInfo", SellerService.GetCommodityInfo)

	//DoSecKill
	http.HandleFunc("/DoSecKill", SellerService.DoSecKill)

	//CreateOrder

	http.HandleFunc("/CreateOrder", SellerService.CreateOrder)

	//SearchOrderInfo
	http.HandleFunc("/SearchOrderInfo", SellerService.SearchOrderInfo)

	//BuildOrder

	//http.HandleFunc("/BuildOrder",orderSystem.BuildOrder)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}

}
