package utitl

import (
	"SeckillDesign/constant"
)

//把数据填充到orderinfo 这个结构体之中
//func FillOrderInfo(snowflakeId string,price string,userId string,activityName string,activityId string)([]byte, error) {
//
//	//orderByte, err := json.Marshal(order)
//	order := constant.Order{}
//
//	order.SetOrderId(snowflakeId)
//	order.SetOrderPrice(price)
//	order.SetUserId(userId)
//	order.SetActivityName(activityName)
//	order.SetActivityId(activityId)
//	//orderByte, err := json.Marshal(order)
//	if
//
//}
	//return

//把数据填充到orderinfo 这个结构体之中
func FillOrderInfo(snowflakeId string,price string,userId string,activityName string,activityId string) constant.Order {
	order := constant.Order{}
	order.SetOrderId(snowflakeId)
	order.SetOrderPrice(price)
	order.SetUserId(userId)
	order.SetActivityName(activityName)
	order.SetActivityId(activityId)
	return order
}




