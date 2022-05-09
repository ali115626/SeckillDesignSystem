package SellerService

import (
	Dao "SeckillDesign/dao/mysql"
	"encoding/json"
	"fmt"
	"net/http"
)


func PullActivityInfo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	requestMap := r.Form
	activityStartTime := requestMap["activityStartTime"][0]
	activityInfoList,err :=Dao.QueryActivityInfo(activityStartTime)
	if err !=nil{
		fmt.Println("Query ActivityInfo from database error,err=",err)
	}
	activityList, err := json.Marshal(activityInfoList)
	if err != nil {
		fmt.Println("json unmarshal error ")
	}
	fmt.Fprintf(w, string(activityList))
}




//todo 我觉得是为了和前端拉齐才会把json弄成interface{}

//todo 这里把ActivityId输进去

//TODO 获取json的信息进行渲染

//todo 这个activityId正好是主键     自己在磁盘上生成一课B+树

//todo 这里弄成一个关于struct的数组  再marshal一下  返回给客户端

//commidtyId := requestMap["commidtyId"][0]
//originPrice := requestMap["originPrice"][0]

//price := requestMap["price"][0]
//todo 这里弄成一个关于struct的数组  再marshal一下  返回给客户端


//--------query

