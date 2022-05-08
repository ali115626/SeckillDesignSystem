package SellerService

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

//todo 我觉得是为了和前端拉齐才会把json弄成interface{}

//todo 这里把ActivityId输进去

//TODO 获取json的信息进行渲染

//todo 这个activityId正好是主键     自己在磁盘上生成一课B+树

//activityName,commidtyId,originPrice,price,stocks,activityStartTime
type ActivityInfo struct {
	ActivityName      string
	CommodityId       string
	OriginPrice       string
	Price             string
	Stocks            string
	ActivityStartTime string
}

func PullActivityInfo(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		return
	}
	requestMap := r.Form
	activityStartTime := requestMap["activityStartTime"][0]

	//commidtyId := requestMap["commidtyId"][0]
	//originPrice := requestMap["originPrice"][0]

	//price := requestMap["price"][0]

	//--------query
	db, err := sql.Open("mysql", "root:123456@/seckill_scheme?charset=utf8")
	if err != nil {
		fmt.Println("open database error,err=", err)
	}

	rows, err := db.Query("select activityName,commidtyId,originPrice,price,stocks,activityStartTime from ActivityTable where activityStartTime=?", activityStartTime)
	if err != nil {
		fmt.Printf("select data error: %v\n", err)
		return
	}
	//todo 这里弄成一个关于struct的数组  再marshal一下  返回给客户端

	var activityInfoList []ActivityInfo

	//var paper string
	//activityInfo := ActivityInfo{}
	var activityInfo ActivityInfo
	for rows.Next() {
		err := rows.Scan(&activityInfo.ActivityName, &activityInfo.CommodityId, &activityInfo.OriginPrice, &activityInfo.Price, &activityInfo.Stocks, &activityInfo.ActivityStartTime)
		//activityInfo.ActivityName=
		//activityInfo.activityStartTime=
		//activityInfo.commodityId=
		////activityInfo.commodityId
		//activityInfo.price=
		//activityInfo.originPrice=
		//activityInfo.stocks=
		fmt.Println(activityInfo)

		activityInfoList = append(activityInfoList, activityInfo)
		if err != nil {

			fmt.Println(err)
			return
		}
	}

	fmt.Println(activityInfoList)

	//------渲染

	activityList, err := json.Marshal(activityInfoList)
	if err != nil {

		fmt.Println("json unmarshal error ")

	}
	fmt.Println(string(activityList))

	fmt.Fprintf(w, string(activityList))

}
