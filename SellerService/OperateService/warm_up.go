package OperateService

import (
	Dao "SeckillDesign/dao/mysql"
	"SeckillDesign/dao/redis"
	"fmt"
	"net/http"
)

func WarmUp(w http.ResponseWriter, r *http.Request) {
	//	把activityId 和 userID 给传进来
	//	把order 写入Mq 中
	err := r.ParseForm()
	if err != nil {
		return
	}
	requestMap := r.Form
	//还有就是你要预热哪一件商品呀！
	activityId := requestMap["activityId"][0]
	//这里可以根据activityId 去获取commodityId   查一次数据库其实也没关系
	isWarmUp := requestMap["isWarmUp"][0]
	if isWarmUp != "yes"{
		return
	}
	// QueryCommodityIdFromActivityTable(activityId string) (string,error){
	//commodityId
	commodityId,err :=Dao.QueryCommodityIdFromActivityTable(activityId)
	if err != nil{
		fmt.Println("Query CommodityId From ActivityTable error,err=",err)
		return
	}
	//timeInterval := requestMap["interval"][0]
	//TODO  这里面应该给其设置一个过期时间  选做
	//activity Info 预热成功
	activityInfo,err := Dao.QueryActivityAllTableFromActivityId(activityId)
	if err != nil{
		fmt.Println("Query	ActivityAllTable From ActivityId error,err=",err)
		return
	}
	//todo 这里面应该弄一个错误的码
	err=redis.SaveActivityInfo(activityId,*activityInfo)
	if err!= nil{
		fmt.Println("Save ActivityInfo error||",err)
		return
	}
	//commodity table的预热
	CommodityInfo,err:=Dao.SearchCommodityDetailFromTable(commodityId)
	if err!=nil{
		fmt.Println("Search Commodity Detail From Table error,err=",err)
		return
	}
	//(*constant.CommodityInfo,error) {
	err=redis.SaveCommodityInfoRedis(commodityId,*CommodityInfo)
	if err!= nil{
		fmt.Println("Save Commodity Info Redis error,err=",err)
		return
	}

	fmt.Fprintf(w,"预热成功！")

	//Redis.SaveCommodityTable






}
