# Aws Sqs Receiver (Email Bounce) [參考](https://aws.amazon.com/tw/blogs/messaging-and-targeting/handling-bounces-and-complaints/)
- 建立 SNS 服務
- SES -> domains -> Notifications 選擇建立的SNS服務
- SQS 訂閱 SNS服務 (ses bounce會推送到sns再推給sqs)

# Aws Configuration and credential file settings [參考](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html) 

- ~/.aws/credentials
```cassandraql
[default]
aws_access_key_id=AKIAIOSFODNN7EXAMPLE
aws_secret_access_key=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
```
- ~/.aws/config
```cassandraql
[default]
region=us-west-2
output=json
```