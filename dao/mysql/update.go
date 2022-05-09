package Dao

import (
	"database/sql"
	"errors"
	"fmt"
)

//todo  这里面如果是要扣减库存的话  别整成负数   校验一下  if  n  >= 0{}

func RevertActivityTableStock(activityId string) error {

	db, err := sql.Open("mysql", "root:123456@/seckill_scheme?charset=utf8")
	if err != nil {
		//fmt.Println("open database error,err=", err)
		return errors.New(fmt.Sprintf("open database error,err=", err))
	}

	_, err = db.Exec("update ActivityTable set available_stock=available_stock+1,locked_stock=locked_stock-1 WHERE activityId=?", activityId)
	if err != nil {
		fmt.Println("exec failed, err=", err)

		return errors.New(fmt.Sprintf("exec failed, err=", err))
	}
	return nil


}
