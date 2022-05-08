package SellerService

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

//todo  这里面输入activityId
//TODO 由activityId得到商品ID  再由商品id 获取商品详情信息   这个外键你还没弄呢  select * from ActivityTable;
//activityId        | int(11)     | NO   | PRI | NULL    | auto_increment |
//| activityName      | varchar(20) | YES  |     | NULL    |                |
//| commidtyId        | int(11)     | YES  |     | NULL    |                |
//| originPrice       | int(11)     | YES  |     | NULL    |                |
//| price             | int(11)     | YES  |     | NULL    |                |
//| stocks            | int(11)     | YES  |     | NULL    |                |
//| activityStartTime | datetime    | YES  | MUL | NULL    |                |
//| activityEndTime   | datetime    | YES  |     | NULL    |                |

type CommodityInfo struct {
	UpdateAt      string `json:"updateAt"`
	CommodityName string `json:"commodityName"`
	Price         string `json:"price"`
	DescribeInfo  string `json:"describeInfo"`
}

//var updateAt string
//var commodityName string
//var price string
//var describeInfo string

func GetCommodityInfo(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		return
	}
	requestMap := r.Form
	activityId := requestMap["activityId"][0]

	db, err := sql.Open("mysql", "root:123456@/seckill_scheme?charset=utf8")
	if err != nil {
		fmt.Println("open database error,err=", err)
	}

	var commidtyId string
	err = db.QueryRow("SELECT commidtyId FROM ActivityTable WHERE activityId=?", activityId).Scan(&commidtyId)
	//发现即使db.QueryRow(）这里面ELECT delete_status FROM blog_info WHERE title_id SQL语句出问题了  也不会报错
	if err != nil {
		fmt.Println(err)
		fmt.Println("select  paper_content  error")
		return

	}

	//desc CommodityTable;
	//+---------------+-------------+------+-----+---------+----------------+
	//| Field         | Type        | Null | Key | Default | Extra          |
	//+---------------+-------------+------+-----+---------+----------------+
	//| commodityId   | int(11)     | NO   | PRI | NULL    | auto_increment |
	//| commodityName | varchar(20) | YES  |     | NULL    |                |
	//| price         | int(11)     | YES  |     | NULL    |                |
	//| describeInfo  | text        | YES  |     | NULL    |                |
	//| updateAt      | datetime    | YES  |     | NULL    |                |

	var updateAt string
	var commodityName string
	var price string
	var describeInfo string

	var commodityInfo CommodityInfo

	err = db.QueryRow("SELECT commodityName,price,describeInfo,updateAt FROM CommodityTable WHERE commodityId=?", commidtyId).Scan(&commodityName, &price, &describeInfo, &updateAt)
	//发现即使db.QueryRow(）这里面ELECT delete_status FROM blog_info WHERE title_id SQL语句出问题了  也不会报错
	if err != nil {
		fmt.Println(err)
		fmt.Println("select  paper_content  error")
		return

	}

	commodityInfo.CommodityName = commodityName
	commodityInfo.DescribeInfo = describeInfo
	commodityInfo.UpdateAt = updateAt
	commodityInfo.Price = price

	commodityContent, err := json.Marshal(commodityInfo)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, string(commodityContent))

}
