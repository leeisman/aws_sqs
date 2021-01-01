package model

import (
	"time"
)

type EmailBounceLog struct {
	ID               uint       `gorm:"primarykey"`
	MessageID        string     `json:"messageId"`
	SourceIP         string     `json:"sourceIp"`
	Subject          string     `json:"subject"`
	From             string     `json:"from"`
	To               string     `json:"to"`
	MailSendDatetime *time.Time `json:"mailSendDatetime" gorm:"column:mail_send_datetime"`
	CreatedAt        *time.Time
	UpdatedAt        *time.Time
}

func (EmailBounceLog) TableName() string {
	return "email_bounce_logs"
}
