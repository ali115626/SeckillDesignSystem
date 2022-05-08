package SellerService

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

//TODO 这个是商品的接口

//TODO  把商品放进去

//todo 先建立一个商品的列表
//| commodityId   | int(11)     | NO   | PRI | NULL    | auto_increment |
//| commodityName | varchar(20) | YES  |     | NULL    |                |
//| price         | int(11)     | YES  |     | NULL    |                |
//| describeInfo  | text

func UploadCommodity(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		return
	}
	requestMap := r.Form
	commodityName := requestMap["commodityName"][0]
	price := requestMap["price"][0]
	describeInfo := requestMap["describeInfo"][0]

	updateAt := time.Now()

	//todo 上传的时间   uploadTime

	//activityStartTime := requestMap["activityStartTime"][0]
	//activityEndTime := requestMap["activityEndTime"][0]
	db, err := sql.Open("mysql", "root:123456@/seckill_scheme?charset=utf8")
	if err != nil {
		fmt.Println("open database error,err=", err)
	}
	_, err = db.Exec("insert into CommodityTable(commodityName,price,describeInfo,updateAt) values(?,?,?,?)", commodityName, price, describeInfo, updateAt)
	if err != nil {
		fmt.Println("exec failed, err=", err)
		return
	}
	fmt.Fprintf(w, "商品上传成功！")

}
