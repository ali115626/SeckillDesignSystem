package SellerService

import (
	Dao "SeckillDesign/dao/mysql"
	"SeckillDesign/dao/redis"
	"encoding/json"
	"fmt"
	"net/http"
)

//跟之前的一样把   input  : commodityId  稍微改改输入输出就好

func ShowCommodityInfo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	requestMap := r.Form

	//	input  : commodityId

	//	先去数据库里面去查  在去database 里面去查

	commodityId := requestMap["commodityId"][0]
	//	input  :activityId
	//redis.

	commodityInfo,err :=redis.QueryCommodityInfoFromRedis(commodityId)
	if err==nil{
		commodityInfoMar,err:= json.Marshal(commodityInfo)
		if err !=nil{
			fmt.Println("json marshal error,err=",err)
		}
		fmt.Println("预热成功！")
		fmt.Fprintf(w,string(commodityInfoMar))
		return

	}

	//	todo 否则就去数据库中去拿吧
	commodityInfoDao,err :=Dao.QueryActivityFromDataBase(commodityId)
	if err!=nil{
		fmt.Println(err)
		return
	}
	commodityInfoDaoMar,err:=json.Marshal(commodityInfoDao)
	if err!= nil{
		fmt.Println(err)
		return
	}
	fmt.Println("从数据库中获取成功")
	fmt.Fprintf(w,string(commodityInfoDaoMar))
}