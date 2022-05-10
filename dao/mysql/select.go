package Dao

import (
	"SeckillDesign/constant"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)


func SearchCommodityDetailFromTable(commodityId string) (*constant.CommodityInfo,error) {
	db, err := sql.Open("mysql", "root:123456@/seckill_scheme?charset=utf8")
	if err != nil {
		fmt.Println("open database error,err=", err)
	}

	var updateAt string
	var commodityName string
	var price string
	var describeInfo string
	var commodityInfo constant.CommodityInfo
	err = db.QueryRow("SELECT commodityName,price,describeInfo,updateAt FROM CommodityTable WHERE commodityId=?", commodityId).Scan(&commodityName, &price, &describeInfo, &updateAt)
	//发现即使db.QueryRow(）这里面ELECT delete_status FROM blog_info WHERE title_id SQL语句出问题了  也不会报错
	if err != nil {
		//fmt.Println(err)
		//fmt.Println("select  paper_content  error")
		return nil,errors.New(fmt.Sprintf("select paper_content error,err="))
	}
	commodityInfo.CommodityName = commodityName
	commodityInfo.DescribeInfo = describeInfo
	commodityInfo.UpdateAt = updateAt
	commodityInfo.Price = price
	return &commodityInfo,nil


}

func QueryCommodityIdFromActivityTable(activityId string) (string,error){
	var commodityId string
	db, err := sql.Open("mysql", "root:123456@/seckill_scheme?charset=utf8")
	if err != nil {
		fmt.Println("open database error,err=", err)
	}
	err = db.QueryRow("SELECT commodityId FROM ActivityTable WHERE activityId=?", activityId).Scan(&commodityId)
	//发现即使db.QueryRow(）这里面ELECT delete_status FROM blog_info WHERE title_id SQL语句出问题了  也不会报错
	if err != nil {
		return "",errors.New(fmt.Sprintf("select paper content error,err=",err))
	}
	return commodityId,nil
}




func QueryActivityInfo(activityStartTime string)(*[]constant.ActivityInfo,error) {
	db, err := sql.Open("mysql", "root:123456@/seckill_scheme?charset=utf8")
	if err != nil {
		return nil,errors.New(fmt.Sprintf("open database error,err=", err))
	}
	rows, err := db.Query("select activityName,commidtyId,originPrice,price,stocks,activityStartTime from ActivityTable where activityStartTime=?", activityStartTime)
	if err != nil {
		return nil,errors.New(fmt.Sprintf("select data error: %v\n", err))
	}
	var activityInfoList []constant.ActivityInfo
	var activityInfo constant.ActivityInfo
	for rows.Next() {
		err := rows.Scan(&activityInfo.ActivityName, &activityInfo.CommodityId, &activityInfo.OriginPrice, &activityInfo.Price, &activityInfo.Stocks, &activityInfo.ActivityStartTime)

		activityInfoList = append(activityInfoList, activityInfo)
		if err != nil {
			//fmt.Println(err)
			return nil,errors.New(fmt.Sprintf("row scan error,err=",err))
		}
	}
	return &activityInfoList,nil


}



func QueryDetailsFromOrderInfoTable(orderNo string)(*constant.OrderInfo,error){
	var userId string
	var activityName string
	var orderPrice string
	var status string
	var createAt string
	var activityId string
	db, err := sql.Open("mysql", "root:123456@/seckill_scheme?charset=utf8")
	if err != nil {
		//fmt.Println("open database error,err=", err)
		//TODO  iphone双11抢购  起的名字总不能是一样的吧
		//	TODO 这里面用hash值去校验用户两次的内容是否一样
		//TODO 其实只要仅校验一下名字就行了   把名字放上索引吧 名字太长了  这里需要数据库的知识
		return nil,errors.New(fmt.Sprintf("open database error,err=", err))
	}
	err = db.QueryRow("SELECT userId,activityId,activityName,orderPrice,status,createAt FROM OrderInfoTable WHERE orderId=?", orderNo).Scan(&userId, &activityId, &activityName, &orderPrice, &status, &createAt)
	//发现即使db.QueryRow(）这里面ELECT delete_status FROM blog_info WHERE title_id SQL语句出问题了  也不会报错
	if err != nil {
		fmt.Println(err)
		fmt.Println("select  paper_content  error")
		return nil,err
	}
	orderInfo := constant.OrderInfo{}
	orderInfo.OrderId = orderNo
	orderInfo.OrderPrice = orderPrice
	statusInt, err := strconv.Atoi(status)
	orderInfo.Status = statusInt
	orderInfo.ActivityId = activityId
	return &orderInfo,nil
}



func QueryOrderInfoStatus(orderNo string)(*string,*string,error){
	db, err := sql.Open("mysql", "root:123456@/seckill_scheme?charset=utf8")
	if err != nil {
		fmt.Println("open database error,err=", err)
		//panic(err)
		//TODO  iphone双11抢购  起的名字总不能是一样的吧

		//	TODO 这里面用hash值去校验用户两次的内容是否一样

		//TODO 其实只要仅校验一下名字就行了   把名字放上索引吧 名字太长了  这里需要数据库的知识
	}
	var activityId string
	var status string
	err = db.QueryRow("SELECT activityId,status FROM OrderInfoTable WHERE orderId=?", orderNo).Scan(&activityId, &status)
	//发现即使db.QueryRow(）这里面ELECT delete_status FROM blog_info WHERE title_id SQL语句出问题了  也不会报错
	if err != nil {
		fmt.Println(err)
		fmt.Println("select  paper_content  error")
		return nil,nil,err
	}
	return &activityId,&status,nil
}

func QueryPriceActivityTable(activityId string)(string,string,error){
	//	TODO  根据activityId去search一下 activityName 和活动的商品的价格
	db, err := sql.Open("mysql", "root:123456@/seckill_scheme?charset=utf8")
	if err != nil {
		return "","",errors.New(fmt.Sprintf("open database error,err=", err))
	}
	var activityName string
	var price string
	//todo 这个时候去search这个表格没有关系  因为商品活动的信息你已经  预热到 redis里面了  但是先去数据库里面查吧
	err = db.QueryRow("SELECT activityName,price FROM ActivityTable WHERE activityId=?", activityId).Scan(&activityName, &price)
	if err != nil {
		return "","",errors.New(fmt.Sprintf("select paper_content error,err=",err))
	}
	return activityName,price,nil
}



func ShowOrderInfo(orderId string)(*constant.OrderInfoShow,error){
	db, err := sql.Open("mysql", "root:123456@/seckill_scheme?charset=utf8")
	if err != nil {
		fmt.Println("open database error,err=", err)
	}
	var status string
	var userId string
	var activityName string
	var orderPrice string
	var createAt string
	err = db.QueryRow("SELECT userId,activityName,orderPrice,status,createAt FROM OrderInfoTable WHERE orderId=?", orderId).Scan(&userId, &activityName, &orderPrice, &status, &createAt)
	//发现即使db.QueryRow(）这里面ELECT delete_status FROM blog_info WHERE title_id SQL语句出问题了  也不会报错
	if err != nil {
		return nil,errors.New(fmt.Sprintf("select paper_content error,err=",err))
	}


	if status == "1" {
		status = "未付款"
	} else if status == "0" {
		status = "订单出错，创建订单失败"
	}

	OrderInfoResult := constant.OrderInfoShow{}
	OrderInfoResult.UserId = userId
	OrderInfoResult.OrderPrice = orderPrice
	OrderInfoResult.Status = status

	OrderInfoResult.CreateAt = createAt

	return &OrderInfoResult,nil
}