FROM golang:alpine

WORKDIR /go/src/goiban-service
COPY ./*.go ./

RUN adduser -D openiban \
&& apk add --no-cache git \
&& go get -d -v ./... \
&& go install -v ./...

USER openiban

CMD ["/go/bin/goiban-service"]