package constant





type Order struct {
	OrderId      string
	UserId       string
	ActivityName string
	OrderPrice   string
	ActivityId   string
}

func (this *Order) SetOrderId(orderId string) {
	this.OrderId = orderId
}

func (this *Order) SetUserId(userId string) {
	this.UserId = userId
}

func (this *Order) SetActivityName(activityName string) {
	this.ActivityName = activityName
}

func (this *Order) SetOrderPrice(commodityPrice string) {
	this.OrderPrice = commodityPrice
}

func (this *Order) SetActivityId(activityId string) {
	this.ActivityId = activityId
}