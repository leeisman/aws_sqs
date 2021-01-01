FROM golang:1.14 as builder

WORKDIR /sqs
COPY . /sqs

ENV CGO_ENABLED=0
ENV GO111MODULE=on
#RUN go mod tidy && go mod download
RUN go build -mod=vendor -o sqs
FROM alpine:3

RUN apk add --no-cache ca-certificates

RUN mkdir -p /sqs
WORKDIR /sqs
COPY --from=builder /sqs/sqs /sqs
COPY --from=builder /sqs/app.yaml /sqs/app.yaml
ENV env=dev
# Run the web service on container startup.
RUN ls -l
CMD ["/sqs/sqs"]