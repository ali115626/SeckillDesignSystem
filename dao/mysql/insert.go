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