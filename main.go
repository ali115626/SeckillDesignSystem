package main

import (
	"SeckillDesign/SellerService"
	OperateService "SeckillDesign/SellerService/OperateService"

	//OperateService "SeckillDesign/SellerService/OperateService"
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
	//todo  其实这些服务不应该放到一起  吧
	//mqConsumerService.BuildOrderConsumer()
	// go mqConsumerService.BuildOrderConsumer()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//////todo 感觉这个有点问题  这个是不应该放在一起的吧
	//这个是单独运行的东西  你启动后 直接监听 listen就行了
	//go mqConsumerService.PayStatusCheckListener()
	//go mqConsumerService.OrderDeductConsumer()
	//todo 你这个检不检查错误有啥用呢
	//todo 这两个接口是商家侧
	//http.HandleFunc("/hello", SellerService.ProcessHello)
	http.HandleFunc("/ReleaseActivity", SellerService.ReleaseActivity)
	//UploadCommodity
	http.HandleFunc("/UploadCommodity", SellerService.ReleaseCommodity)
	//PullActivityInfo
	http.HandleFunc("/PullActivityInfo", SellerService.PullActivityInfo)
	//GetCommodityInfo
	http.HandleFunc("/GetCommodityInfo", SellerService.GetCommodityInfo)
	//DoSecKill
	http.HandleFunc("/DoSecKill", SellerService.DoSecKill)//这个还是待压测

	http.HandleFunc("/CreateOrder", SellerService.CreateOrder)

	http.HandleFunc("/SearchOrderInfo", SellerService.SearchOrderInfo)
	//CreateOrder
	//SearchOrderInfo
	http.HandleFunc("/ProcessPayDoneOrder", SellerService.ProcessPayDoneOrder)
	//ProcessPayDoneOrder


	//WarmUp

	http.HandleFunc("/WarmUp", OperateService.WarmUp)


	//QueryActivityFromRedis

	//ShowActivityInfo

	http.HandleFunc("/ShowActivityInfo", SellerService.ShowActivityInfo)

	//ShowCommodityInfo
	http.HandleFunc("/ShowCommodityInfo", SellerService.ShowCommodityInfo)

	//显示倒计时CutDownWarmUp
	http.HandleFunc("/CutDownWarmUp", OperateService.CutDownWarmUp)

	//err := http.ListenAndServe(":8080", nil)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}

}
