package subscriber

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	configs "sqs/config"
	"sqs/model"
)

type Subscriber struct {
	BounceReceive chan model.Body
	Config        *configs.Config
}

func NewSubscriber(config *configs.Config) *Subscriber {
	sub := &Subscriber{
		Config: config,
	}
	sub.BounceReceive = make(chan model.Body, 0)
	return sub
}

func (s *Subscriber) Subscribe() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := sqs.New(sess)
	qURL := s.Config.SQS.QURL
	for {
		result, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
			AttributeNames: []*string{
				aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
			},
			MessageAttributeNames: []*string{
				aws.String(sqs.QueueAttributeNameAll),
			},
			QueueUrl:            &qURL,
			MaxNumberOfMessages: aws.Int64(10),
			VisibilityTimeout:   aws.Int64(60), // 60 seconds
			WaitTimeSeconds:     aws.Int64(0),
		})
		if err != nil {
			fmt.Println("Error", err)
			return
		}
		if len(result.Messages) == 0 {
			fmt.Println("Received no messages")
			continue
		}
		messge := result.Messages[0]
		body := s.formatToBody(*messge.Body)
		if body.Message.NotificationType == "Bounce" {
			s.BounceReceive <- body
		}
	}
}

func (s *Subscriber) formatToBody(jsonString string) model.Body {
	body := model.Body{}
	err := json.Unmarshal([]byte(jsonString), &body)
	if err != nil {
		fmt.Println("err ", err)
		return body
	}
	return body
}
