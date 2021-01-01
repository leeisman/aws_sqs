package subscriber

import (
	"fmt"
	"sqs/handler"
	"testing"
)

func TestSubscriber_Subscribe(t *testing.T) {
	sqsSubscribeProcess := NewSubscriber()
	sqsSubscribeProcess.Subscribe()
	Handlers := &handler.SubHandler{}
	for {
		body := <-sqsSubscribeProcess.BounceReceive
		Handlers.HandleBounce(body)
	}
}

func TestSubscriber_formatToBody(t *testing.T) {
	subscriber := NewSubscriber()

	jsonString := `{
	 "Type" : "Notification",
	 "MessageId" : "c2c0be36-cc60-5b8c-b097-e65ec2c5e056",
	 "TopicArn" : "arn:aws:sns:us-west-2:590351187300:ses-bounces-topic",
	 "Message" : "{\"notificationType\":\"Bounce\",\"bounce\":{\"feedbackId\":\"01010176b9e06b0b-6a019094-6673-43cb-a5b3-9d6f649ba00d-000000\",\"bounceType\":\"Transient\",\"bounceSubType\":\"General\",\"bouncedRecipients\":[{\"emailAddress\":\"frankie.lee.job@gmailabc.com\",\"action\":\"failed\",\"status\":\"4.4.7\",\"diagnosticCode\":\"smtp; 554 4.4.7 Message expired: unable to deliver in 840 minutes.<421 4.4.2 Connection closed unexpectedly>\"}],\"timestamp\":\"2020-12-31T17:37:41.000Z\",\"remoteMtaIp\":\"34.102.136.180\",\"reportingMTA\":\"dsn; a27-50.smtp-out.us-west-2.amazonses.com\"},\"mail\":{\"timestamp\":\"2020-12-31T03:18:50.598Z\",\"source\":\"support@anytime.tw\",\"sourceArn\":\"arn:aws:ses:us-west-2:590351187300:identity/anytime.tw\",\"sourceIp\":\"211.21.120.39\",\"sendingAccountId\":\"590351187300\",\"messageId\":\"01010176b6ce1fe6-7ce4d2f3-e075-4386-8574-f91a811fdd89-000000\",\"destination\":[\"frankie.lee.job@gmailabc.com\"],\"headersTruncated\":false,\"headers\":[{\"name\":\"Received\",\"value\":\"from anytime.test (211-21-120-39.HINET-IP.hinet.net [211.21.120.39]) by email-smtp.amazonaws.com with SMTP (SimpleEmailService-d-IH1GZADI7) id gXegksLzx2Pp4mYBnBqH for frankie.lee.job@gmailabc.com; Thu, 31 Dec 2020 03:18:50 +0000 (UTC)\"},{\"name\":\"Message-ID\",\"value\":\"<83e4571828d3c5750c63e6887f1d5e42@anytime.test>\"},{\"name\":\"Date\",\"value\":\"Thu, 31 Dec 2020 03:18:47 +0000\"},{\"name\":\"Subject\",\"value\":\"Anytime Verification Mail\"},{\"name\":\"From\",\"value\":\"Anytime <support@anytime.tw>\"},{\"name\":\"To\",\"value\":\"frankie.lee.job@gmailabc.com\"},{\"name\":\"MIME-Version\",\"value\":\"1.0\"},{\"name\":\"Content-Type\",\"value\":\"text/html; charset=utf-8\"},{\"name\":\"Content-Transfer-Encoding\",\"value\":\"quoted-printable\"}],\"commonHeaders\":{\"from\":[\"Anytime <support@anytime.tw>\"],\"date\":\"Thu, 31 Dec 2020 03:18:47 +0000\",\"to\":[\"frankie.lee.job@gmailabc.com\"],\"messageId\":\"<83e4571828d3c5750c63e6887f1d5e42@anytime.test>\",\"subject\":\"Anytime Verification Mail\"}}}",
	 "Timestamp" : "2020-12-31T17:37:41.383Z",
	 "SignatureVersion" : "1",
	 "Signature" : "qo/VKQ/GeLnKb9CCONdMbp6vL2dtdCm6HfHFyMyBfNQju4YHJVGlopoHyHSBqCQAHQFH/TptHE1hhXU0n4KJZVntydGzhOA3g4tg0eKJ3cNuW3ENHW3sA1048i7ZfAF3IJ/zfSemiifDuxIHnvH9fVWoIIsMzRoQrh/UwOa76EaTw7kALc02IqeHpae8oyo3GurJXFjGiq/fzrfP6kbacDLlU9EhuXHFqaTQEt8lAx2hqYlOZ+QHF0LC3/PJPUM9KKOC/nr/Gj+vsx9OnsOgv9KNBkfgsIZg+CWlRAHwYJPS4rEAJfRzwivb036hGcG5P1wDWJjCaI4htHd7ExpxiA==",
	 "SigningCertURL" : "https://sns.us-west-2.amazonaws.com/SimpleNotificationService-010a507c1833636cd94bdb98bd93083a.pem",
	 "UnsubscribeURL" : "https://sns.us-west-2.amazonaws.com/?Action=Unsubscribe&SubscriptionArn=arn:aws:sns:us-west-2:590351187300:ses-bounces-topic:00ce26b5-a1f0-4fd0-90c0-98768fa36379"
	}`
	body := subscriber.formatToBody(jsonString)
	fmt.Println("show")
	fmt.Println(body.Type)
	fmt.Println(body.Message.Mail)

}
