package model

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type MailHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type CommonHeaders struct {
	// Anytime <support@anytime.tw>
	From []string `json:"from"`
	// Wed, 30 Dec 2020 17:35:18 +0000
	Date string `json:"date"`
	// ["frankie.lee.job@gmailabc.comabc"]
	To []string `json:"to"`
	// <c5b579585f0b7e885dff24b66af95b04@anytime.test>
	MessageID string `json:"messageId"`
	// Anytime Verification Mail
	Subject string `json:"subject"`
}
type Mail struct {
	// 2020-12-30T17:35:21.855Z
	Timestamp string `json:"timestamp"`
	// support@anytime.tw
	Source string `json:"source"`
	// arn:aws:ses:us-west-2:590351187300:identity/anytime.tw
	SourceArn string `json:"sourceArn"`
	// 123.194.2.82
	SourceIp string `json:"sourceIp"`
	// ["frankie.lee.job@gmailabc.comabc"]
	Destination   []string      `json:"destination"`
	Headers       []MailHeader  `json:"headers"`
	CommonHeaders CommonHeaders `json:"commonHeaders"`
}

type Message struct {
	// "Bounce"
	NotificationType string `json:"notificationType"`
	Mail             Mail   `json:"mail"`
}

type Body struct {
	Type    string  `json:"type"`
	Message Message `json:"message"`
}

func (m *Message) UnmarshalJSON(data []byte) error {
	var mm struct {
		NotificationType string `json:"notificationType"`
		Mail             Mail   `json:"mail"`
	}

	fmt.Println(strconv.Unquote(string(data)))
	buf, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(buf), &mm)
	if err != nil {
		return err
	}
	*m = Message{
		NotificationType: mm.NotificationType,
		Mail:             mm.Mail,
	}
	return nil
}
