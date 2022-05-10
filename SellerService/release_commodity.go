package SellerService

import (
	"SeckillDesign/constant"
	Dao "SeckillDesign/dao/mysql"
	"fmt"
	"net/http"
	"time"
)

//TODO 这个是商品的接口

//TODO  把商品放进去

//todo 先建立一个商品的列表
func ReleaseCommodity(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	requestMap := r.Form
	commodityName := requestMap["commodityName"][0]
	price := requestMap["price"][0]
	describeInfo := requestMap["describeInfo"][0]
	updateAt := time.Now().String()
	//todo 上传的时间   uploadTime
	commodityInfo:=constant.CommodityInfo{
		UpdateAt:      updateAt,
		CommodityName: commodityName,
		Price:         price,
		DescribeInfo:  describeInfo,
	}

	err =Dao.InsetIntoCommodityTable(commodityInfo)
	if err !=nil{
		//阻断它  别让其商品上传成功
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, "商品上传成功！")

}
