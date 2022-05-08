package SellerService

import (
	"SeckillDesign/constant"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

func BuildOrderProcess(orderInfo constant.OrderInfo)error{


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
	orderInfo.Status = 1
	//TODO 这个要做一下转化哈
	userId, err := strconv.Atoi(orderInfo.UserId)
	if err != nil {
		fmt.Println(err)
	}
	orderPrice, err := strconv.Atoi(orderInfo.OrderPrice)
	if err != nil {
		fmt.Println(err)

	}
	activityId, err := strconv.Atoi(orderInfo.ActivityId)
	if err != nil {
		fmt.Println(err)
	}
	status := strconv.Itoa(orderInfo.Status)
	insertResult, err := db.Exec("insert into OrderInfoTable(orderId,userId,activityName,orderPrice,activityId,status) values(?,?,?,?,?,?)", orderInfo.OrderId, userId, orderInfo.ActivityName, orderPrice, activityId, status)
	if err != nil {
		fmt.Println("exec failed, err=", err)
		return errors.New(fmt.Sprintf("exec failed, err=", err))
	}
	//todo  这样的话  insert进数据库就成功了
	fmt.Println(insertResult.RowsAffected())
	return nil

	//TODO  否则 insert到数据库中呗

	//	TODO  如果扣减失败了  直接return回去  扣减成功  修改status insert进数据库

	//		fmt.Println(err)
	//		fmt.Println("select  paper_content  error")
	//		return
	//
	//	}
	//
	//	if stocks > 0{
	//
	//	}

	//todo  所以 你那边还要去确认一下   有没有库存

}

//var stocks string
//todo  数据库这里 你就不能用一下乐观锁吗？  update ActivityTable set available_stock=available_stock-1,locked_stock=locked_stock+1 WHERE activityId= orderInfo.activityId and stocks

//todo 这个万一fail了 库存扣减失败  status = 0  这里的话 就直接return
//todo 这个如果成功了的话   库存扣减成功 status=1

//TODO	先去查库存  看看还有没有这个订单信息  如果有库存  把status改成 1  否则改成 status = 0
//	err = db.QueryRow("SELECT stocks FROM ActivityTable WHERE activityId=?", orderInfo.ActivityId).Scan(&stocks)
//	//发现即使db.QueryRow(）这里面ELECT delete_status FROM blog_info WHERE title_id SQL语句出问题了  也不会报错
//	if err != nil {
//	_, err = db.Exec("update ActivityTable set available_stock=available_stock-1,locked_stock=locked_stock+1 WHERE activityId= and available_stock > ) values(?,?)",orderInfo.ActivityId,0)
//	if err != nil {
//		fmt.Println("exec failed, err=", err)
//		//	todo 我就奇怪了   万一我在这边一直点下去怎么办   代码层校验是否重复上传   怎么办呢
//	}