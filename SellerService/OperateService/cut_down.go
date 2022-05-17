package OperateService

import (
	"SeckillDesign/dao/redis"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func CutDownWarmUp(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	requestMap := r.Form
	activityId := requestMap["activityId"][0]

	activityStartTime,err:=redis.QueryTimeFromActivityTable(activityId)

	if err!=nil{
		fmt.Println(err)
		return
	}


	Now := time.Now().Unix()

	//futureTime := activityStartTime

	//-------------
	toBeCharge := activityStartTime                            //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "2006-01-02 15:04:05"                             //转化所需模板
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, err := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	if err !=nil{
		fmt.Println(err)
		return
	}
	futureTime := theTime.Unix() //转化为时间戳 类型是int64

	fmt.Println("futureTime=",futureTime)
	fmt.Println("Now=",Now)

	interval :=strconv.FormatInt(futureTime-Now,10)

	fmt.Fprintf(w,fmt.Sprintf("距离活动开始还有"+interval+"秒"))

}





//	先根据activityId 去查 startTime
//	return startTime-time.Now()



