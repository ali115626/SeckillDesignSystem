package DownLogic

import (
	"SeckillDesign/constant"
	Dao "SeckillDesign/dao/mysql"
	//"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/streadway/amqp"
)


func BuildOrderProcessService(msg amqp.Delivery)error{
	message := string(msg.Body)
	//	再去unMarshal一下  unmarshal 到 orderInfo中
	fmt.Println("message=",message)
	orderInfo := constant.OrderInfo{}
	err := json.Unmarshal([]byte(message), &orderInfo)
	if err != nil {
		return errors.New(fmt.Sprintf("json unmarshal orderInfo err,err=", err))
	}
	//TODO 这里面会有一致性的问题吧
	//TODO 万一这里扣了两次库存咋办
	//TODO 会有那种宕机的情况
	//TODO 就是rabbitMQ  那里
	//TODO 就不怕它扣两次库存吗？
	//TODO  用CAS
	err =Dao.LockActivityTableStock(orderInfo)
	if err != nil {
		fmt.Println(fmt.Sprintf("lock activityTable stock error,err=", err))
		return errors.New(fmt.Sprintf("lock activityTable stock error,err=", err))
	}

	orderInfo.Status = 1
	err = Dao.InsertOrderInfoTable(orderInfo)

	if err!=nil{
		return err
	}
	return nil

}


//TODO 否则 insert到数据库中呗
//	TODO  如果扣减失败了  直接return回去  扣减成功  修改status insert进数据库
//todo  所以 你那边还要去确认一下   有没有库存
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

//	TODO 这个你要unMarshal一下吗？
//todo  检查一下  库存里面有没有
//TODO 有的话  给前端返回1 ：创建订单成功  2：已经支付了（支付成功了）

//	TODO 把订单信息insert到数据库