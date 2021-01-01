package main

import (
	configs "aws_sqs/config"
	db2 "aws_sqs/db"
	"aws_sqs/handler"
	"aws_sqs/subscriber"
	"fmt"
)

func main() {

	config := configs.NewConfig()
	fmt.Println(config)
	db := db2.NewDB(config)
	fmt.Println(db)

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
