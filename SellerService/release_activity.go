package SellerService

import (
	"SeckillDesign/constant"
	Dao "SeckillDesign/dao/mysql"
	//"database/sql"
	"fmt"
	"net/http"
	//"database/sql"
	//"time"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)



func ReleaseActivity(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}

	requestMap := r.Form
	activityName := requestMap["activityName"][0]
	commodityId := requestMap["commodityId"][0]
	originPrice := requestMap["originPrice"][0]
	price := requestMap["price"][0]
	stocks := requestMap["stocks"][0]
	availableStock := requestMap["availableStock"][0]
	activityStartTime := requestMap["activityStartTime"][0]
	//todo 这个endTime不用传过来
	activityEndTime := requestMap["activityEndTime"][0]

	activityInfo :=constant.ActivityInfo{
		ActivityName:activityName,
		CommodityId:commodityId,
		OriginPrice:originPrice,
		Price:price,
		Stocks:stocks,
		AvailableStock:availableStock,
		ActivityStartTime:activityStartTime,
		ActivityEndTime:activityEndTime,
	}

	err =Dao.InsertActivityInfo(activityInfo)
	if err != nil{
		fmt.Println("Insert Activity Info error ,err=",err)
		return
	}
	fmt.Fprintf(w, "上传活动信息正常！")
}








// TODO 这个就是一个handler  商家posetman上传一个活动   信息

//TODO  把这个活动保存到数据库里面呀

//TODO 先存到mysql 里面吧   后面到了缓存预热的时候  再将其放到redis里面
//todo  就是按照这种格式去insert到数据库里面

//todo 这个的话  前端就规定--- 格式吧

