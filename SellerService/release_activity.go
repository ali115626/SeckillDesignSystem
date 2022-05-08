package SellerService

import (
	"database/sql"
	"fmt"
	"net/http"
	//"database/sql"
	//"time"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// TODO 这个就是一个handler  商家posetman上传一个活动   信息

//TODO  把这个活动保存到数据库里面呀

//TODO 先存到mysql 里面吧   后面到了缓存预热的时候  再将其放到redis里面

//
//activityId        | int(11)     | NO   | PRI | NULL    |       |
//| activityName      | varchar(20) | YES  |     | NULL    |       |
//| commidtyId        | int(11)     | YES  |     | NULL    |       |
//| originPrice       | int(11)     | YES  |     | NULL    |       |
//| price             | int(11)     | YES  |     | NULL    |       |
//| stocks            | int(11)     | YES  |     | NULL    |       |
//| activityStartTime | datetime    | YES  |     | NULL    |       |
//| activityEndTime   | datetime    | YES  |     | NULL
//todo  就是按照这种格式去insert到数据库里面
//str := "2021-01-03 15:23:11"
//loc, _ := time.LoadLocation("Asia/Shanghai")
//d, _ := time.ParseInLocation("2006-01-02 15:04:05", str, loc)
//fmt.Printf("time: %v\n", d)

//todo 这个的话  前端就规定--- 格式吧
func ReleaseActivity(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		return
	}
	requestMap := r.Form
	activityName := requestMap["activityName"][0]

	commidtyId := requestMap["commidtyId"][0]
	originPrice := requestMap["originPrice"][0]

	price := requestMap["price"][0]

	stocks := requestMap["stocks"][0]
	availableStock := requestMap["availableStock"][0]

	activityStartTime := requestMap["activityStartTime"][0]
	activityEndTime := requestMap["activityEndTime"][0]
	db, err := sql.Open("mysql", "root:123456@/seckill_scheme?charset=utf8")
	if err != nil {
		fmt.Println("open database error,err=", err)

		//panic(err)
		//TODO  iphone双11抢购  起的名字总不能是一样的吧

		//	TODO 这里面用hash值去校验用户两次的内容是否一样

		//TODO 其实只要仅校验一下名字就行了   把名字放上索引吧 名字太长了  这里需要数据库的知识

	}
	_, err = db.Exec("insert into ActivityTable(activityName,commidtyId,originPrice,price,stocks,activityStartTime,activityEndTime,available_stock) values(?,?,?,?,?,?,?,?)", activityName, commidtyId, originPrice, price, stocks, activityStartTime, activityEndTime, availableStock)
	if err != nil {
		fmt.Println("exec failed, err=", err)
		//	todo 我就奇怪了   万一我在这边一直点下去怎么办   代码层校验是否重复上传   怎么办呢
	}
	fmt.Fprintf(w, "上传活动信息正常！")

}
