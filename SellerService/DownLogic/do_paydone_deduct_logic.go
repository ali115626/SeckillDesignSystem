package DownLogic

import (
	"SeckillDesign/constant"
	Dao "SeckillDesign/dao/mysql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/streadway/amqp"
)

func DoPayDoneDeductLogic(msg amqp.Delivery)error{
	//TODO 把这些逻辑抽成一个函数吧
	message := string(msg.Body)
	//	再去unMarshal一下  unmarshal 到 orderInfo中
	orderInfo :=constant.OrderInfo{}
	err := json.Unmarshal([]byte(message), orderInfo)
	if err != nil {
		return errors.New(fmt.Sprintf("json unmarshal orderInfo err,err=", err))
	}
	activityId := orderInfo.ActivityId
	//其实这个就是将锁定的库存给恢复一下
	err =Dao.DeductActivityTableStock(activityId)
	if err!=nil{
		//fmt.Println("Deduct Stock error,err=",err)
		return errors.New(fmt.Sprintf("Deduct Stock error,err=",err))
	}
	return nil
}




