package main

import (
	configs "sqs/config"
	db2 "sqs/db"
	"sqs/handler"
	subscriber "sqs/subscriber"
)

func main() {

	config := configs.NewConfig()
	db := db2.NewDB(config)

	sqsSubscribeProcess := subscriber.NewSubscriber(config)
	Handlers := handler.NewSubHandler(config, db)

	// 非同步取得bounce
	go func(sqsSubscribeProcess *subscriber.Subscriber) {
		sqsSubscribeProcess.Subscribe()
	}(sqsSubscribeProcess)

	// loop 取得 bounce
	for {
		body := <-sqsSubscribeProcess.BounceReceive
		Handlers.HandleBounce(body)
	}
}
