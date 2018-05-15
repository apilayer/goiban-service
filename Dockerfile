FROM golang:alpine

WORKDIR /go/src/goiban-service
COPY ./*.go ./

RUN apk add --no-cache git
RUN go get -d -v ./...
RUN go install -v ./...

CMD ["sh", "-c", "goiban-service -w"]