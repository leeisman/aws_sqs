build-linux:
	env GOOS=linux GOARCH=amd64  go build -mod=vendor -o sqs

daemon-run:
	./sqs >/dev/null 2>&1 < /dev/null &