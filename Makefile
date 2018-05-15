.PHONY: dev build

dev:
	go build
	./goiban-service 8080 root:root@/goiban?charset=utf8

build:
	./build.sh