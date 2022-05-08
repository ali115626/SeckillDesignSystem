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
