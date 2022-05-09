package constant

const NormalQueue string = "normal_queue"
const NormalExchange string = "normal_exchange"
const DeadQueue string = "dead_queue"
const DeadExchange string = "dead_exchange"

const DeadRoutingKey string = "dead_routing_key"

const NormalRoutingKey string = "normal_routing_key"

type OrderInfo struct {
	OrderId      string
	UserId       string
	ActivityName string
	OrderPrice   string
	ActivityId   string
	Status       int
}


type ActivityInfo struct {
	ActivityName      string
	CommodityId       string
	OriginPrice       string
	Price             string
	Stocks            string
	ActivityStartTime string
}



type CommodityInfo struct {
	UpdateAt      string `json:"updateAt"`
	CommodityName string `json:"commodityName"`
	Price         string `json:"price"`
	DescribeInfo  string `json:"describeInfo"`
}
