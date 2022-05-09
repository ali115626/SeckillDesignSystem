package Dao

import (
	"database/sql"
	"fmt"
)

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
