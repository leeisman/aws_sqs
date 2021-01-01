package main

import (
	"sqs/handler"
	subscriber "sqs/subscriber"
)

func main() {
	sqsSubscribeProcess := subscriber.NewSubscriber()
	sqsSubscribeProcess.Subscribe()
	Handlers := &handler.SubHandler{}
	for {
		body := <-sqsSubscribeProcess.BounceReceive
		Handlers.HandleBounce(body)
	}
}
