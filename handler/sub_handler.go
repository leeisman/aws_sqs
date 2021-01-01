package handler

import (
	"fmt"
	"sqs/model"
)

type SubHandler struct {
}

func (s *SubHandler) HandleBounce(body model.Body) {
	fmt.Println("get!!!!!!!")
	fmt.Println(body.Type)
	fmt.Println(body.Message.Mail.CommonHeaders.Subject)
	fmt.Println(body.Message.Mail.CommonHeaders.From)
}
