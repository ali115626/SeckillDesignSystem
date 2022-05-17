package Dao

import (
	"SeckillDesign/constant"
	"database/sql"

	//"SeckillDesign/constant"
	"errors"
	"fmt"
	"strconv"
)

//todo  这样的话  insert进数据库就成功了


func InsetIntoCommodityTable(commodityInfo constant.CommodityInfo)error {
	db, err := sql.Open("mysql", "root:123456@/seckill_scheme?charset=utf8")
	if err != nil {
		return errors.New(fmt.Sprintf("open database error,err=", err))
	}
	_, err = db.Exec("insert into CommodityTable(commodityName,price,describeInfo,updateAt) values(?,?,?,?)", commodityInfo.CommodityName,commodityInfo.Price,commodityInfo.DescribeInfo, commodityInfo.UpdateAt)
	if err != nil {
		return errors.New(fmt.Sprintf("exec failed, err=", err))
	}
	return nil
}

func InsertOrderInfoTable(orderInfo constant.OrderInfo) error{
	db, err := sql.Open("mysql", "root:123456@/seckill_scheme?charset=utf8")
	if err != nil {
		return errors.New(fmt.Sprintf("open database error,err=", err))
	}
	userId, err := strconv.Atoi(orderInfo.UserId)
	if err != nil {
		fmt.Println(err)
		return err
	}
	orderPrice, err := strconv.Atoi(orderInfo.OrderPrice)
	if err != nil {
		fmt.Println(err)
		return err//这个你总要标识 是那个程序出的问题吧
	}
	activityId, err := strconv.Atoi(orderInfo.ActivityId)
	if err != nil {
		fmt.Println(err)
		return err
	}
	status := strconv.Itoa(orderInfo.Status)

	//TODO  这里面先要检查一下  orderTable里面有没有

	_, err = db.Exec("insert into OrderInfoTable(orderId,userId,activityName,orderPrice,activityId,status) values(?,?,?,?,?,?)", orderInfo.OrderId, userId, orderInfo.ActivityName, orderPrice, activityId, status)
	if err != nil {
		fmt.Println("exec failed, err=", err)
		return errors.New(fmt.Sprintf("exec failed, err=", err))
	}
	return nil
}


func InsertActivityInfo(activityInfo constant.ActivityInfo) error{
	db, err := sql.Open("mysql", "root:123456@/seckill_scheme?charset=utf8")
	if err != nil {

		return errors.New(fmt.Sprintf("open database error,err=", err))
		//TODO  iphone双11抢购  起的名字总不能是一样的吧
		//	TODO 这里面用hash值去校验用户两次的内容是否一样
		//TODO 其实只要仅校验一下名字就行了   把名字放上索引吧 名字太长了  这里需要数据库的知识
	}
	_, err = db.Exec("insert into ActivityTable(activityName,commidtyId,originPrice,price,stocks,activityStartTime,activityEndTime,available_stock) values(?,?,?,?,?,?,?,?)", activityInfo.ActivityName, activityInfo.CommodityId, activityInfo.OriginPrice, activityInfo.Price,activityInfo.Stocks,activityInfo.ActivityStartTime,activityInfo.ActivityEndTime, activityInfo.AvailableStock)
	if err != nil {
		return errors.New(fmt.Sprintf("exec failed, err=", err))
		//	todo 我就奇怪了   万一我在这边一直点下去怎么办   代码层校验是否重复上传   怎么办呢
	}
	return nil
}