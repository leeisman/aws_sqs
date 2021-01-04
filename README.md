# Aws sqs receiver (email bounce)
- 建立 SNS 服務
- SES -> domains -> Notifications 選擇建立的SNS服務
- SQS 訂閱 SNS服務 (ses bounce會推送到sns再推給sqs)
