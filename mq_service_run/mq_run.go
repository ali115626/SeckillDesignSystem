package main

import (
	"SeckillDesign/mqConsumerService"
)

func main(){
	var intchan chan int

	intchan =make(chan int)
	//他的这里其实是阻塞的
	go mqConsumerService.BuildOrderConsumer()

	go mqConsumerService.PayStatusCheckListener()

	go mqConsumerService.OrderDeductConsumer()

	<-intchan

}
