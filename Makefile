build-linux:
	env GOOS=linux GOARCH=amd64  go build -mod=vendor -o aws_sqs_receiver

daemon-run:
	./aws_sqs_receiver >/dev/null 2>&1 < /dev/null &