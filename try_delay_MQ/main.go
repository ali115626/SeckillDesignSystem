package main

import "SeckillDesign/producer"

func main() {

	message := "hello"
	array := make([]int, 10)

	for _ = range array {

		//
		producer.DelayProducer(message)
	}

}
