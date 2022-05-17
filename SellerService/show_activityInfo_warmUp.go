package SellerService

import (
	Dao "SeckillDesign/dao/mysql"
	"SeckillDesign/dao/redis"
	"encoding/json"
	"fmt"
	"net/http"
)

func ShowActivityInfo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	requestMap := r.Form
	activityId := requestMap["activityId"][0]
//	input  :activityId
	//redis.
	activityInfo,err :=redis.QueryActivityFromRedis(activityId)
	if err==nil{
		activityInfoMar,err:= json.Marshal(activityInfo)
		if err !=nil{
			fmt.Println("json marshal error,err=",err)
		}
		fmt.Println("预热成功！")
		fmt.Fprintf(w,string(activityInfoMar))
		return
	}

//	todo 否则就去数据库中去拿吧
	activityInfoDao,err :=Dao.QueryActivityFromDataBase(activityId)
	if err!=nil{
		fmt.Println(err)
		return
	}
	activityInfoDaoMar,err:=json.Marshal(activityInfoDao)
	if err!= nil{
		fmt.Println(err)
		return
	}
	fmt.Println("从数据库中获取成功")
	fmt.Fprintf(w,string(activityInfoDaoMar))


//	先去数据库里面去查  在去database 里面去查







}
