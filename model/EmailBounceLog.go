package model

import (
	"time"
)

type EmailBounceLog struct {
	ID               uint      `gorm:"primarykey"`
	MessageID        string    `json:"messageId",gorm:"message_id"`
	SourceIP         string    `json:"sourceIp",gorm:"source_ip"`
	Subject          string    `json:"subject",gorm:"subject"`
	From             string    `json:"from",gorm:"from"`
	To               string    `json:"to",gorm:"to"`
	MailSendDateTime time.Time `json:"mailSendDateTime",gorm:"mail_send_datetime"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (EmailBounceLog) TableName() string {
	return "email_bounce_logs"
}
