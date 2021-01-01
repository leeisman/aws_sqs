package handler

import (
	"encoding/json"
	"fmt"
	configs "sqs/config"
	"sqs/db"
	"sqs/model"
)

type SubHandler struct {
	Config *configs.Config
	DB     *db.DB
}

func NewSubHandler(config *configs.Config, db *db.DB) *SubHandler {
	return &SubHandler{
		Config: config,
		DB:     db,
	}
}
func (s *SubHandler) HandleBounce(body model.Body) {
	emailBounceLog := model.EmailBounceLog{
		MessageID: body.Message.Mail.CommonHeaders.MessageID,
		SourceIP:  body.Message.Mail.SourceIp,
		Subject:   body.Message.Mail.CommonHeaders.Subject,
		From:      body.Message.Mail.CommonHeaders.From[0],
		To:        body.Message.Mail.CommonHeaders.To[0],
	}

	err := s.DB.Writer.Find(&model.EmailBounceLog{}).Error
	if err != nil {
		fmt.Print("find err: ", err.Error())
	}
	fmt.Println("get!!!!!!!")
	jsonMarshal, _ := json.Marshal(emailBounceLog)
	fmt.Println(string(jsonMarshal))

	s.DB.Writer.Create(emailBounceLog)

}
