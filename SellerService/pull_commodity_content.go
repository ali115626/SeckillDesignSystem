package SellerService

import (
	Dao "SeckillDesign/dao/mysql"
	"encoding/json"
	"fmt"
	"net/http"
)

//todo  这里面输入activityId
//TODO 由activityId得到商品ID  再由商品id 获取商品详情信息   这个外键你还没弄呢  select * from ActivityTable;


func GetCommodityInfo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	requestMap := r.Form
	activityId := requestMap["activityId"][0]
	var commodityId string
	commodityId,err =Dao.QueryCommodityIdFromActivityTable(activityId)
	if err != nil{
		fmt.Println(err)
	}
	commodityInfo,err :=Dao.SearchCommodityDetailFromTable(commodityId)
	if err != nil{
		fmt.Println(err)
	//	查询数据库失败  那你直接就return 呗  还继续走下去吗
		return
	}
	commodityContent, err := json.Marshal(commodityInfo)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, string(commodityContent))

}
