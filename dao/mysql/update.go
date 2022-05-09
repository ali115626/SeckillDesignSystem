package Dao

import (
	"SeckillDesign/constant"
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


func DeductActivityTableStock(activityId string) error {

	db, err := sql.Open("mysql", "root:123456@/seckill_scheme?charset=utf8")
	if err != nil {
		//fmt.Println("open database error,err=", err)
		return errors.New(fmt.Sprintf("open database error,err=", err))
	}

	_, err = db.Exec("update ActivityTable set locked_stock=locked_stock-1 WHERE activityId=?", activityId)
	if err != nil {
		fmt.Println("exec failed, err=", err)

		return errors.New(fmt.Sprintf("exec failed, err=", err))
	}
	return nil
}



func LockActivityTableStock(orderInfo constant.OrderInfo) error {
	db, err := sql.Open("mysql", "root:123456@/seckill_scheme?charset=utf8")
	if err != nil {
		return errors.New(fmt.Sprintf("open database error,err=", err))
	}

	result, err := db.Exec("update ActivityTable set available_stock=available_stock-1,locked_stock=locked_stock+1 WHERE activityId=? and available_stock > ?", orderInfo.ActivityId, 0)
	if err != nil {
		return errors.New(fmt.Sprintf("exec failed, err=", err))
		//	todo 我就奇怪了   万一我在这边一直点下去怎么办   代码层校验是否重复上传   怎么办呢
	}
	RowsAffected, err := result.RowsAffected()
	if RowsAffected == 0 {
		orderInfo.Status = 0
		//	TODO status :将其设为  库存扣减失败
		return errors.New(
			"库存扣减失败！") //前端发现你这个东西等于0了   然后就开始  todo 这样的话  也没有必要扣减数据库了
	}
	return nil


}











