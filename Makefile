.PHONY: dev build docker

dev:
	go build
	./goiban-service 8080 root:root@/goiban?charset=utf8

build:
	./build.sh

docker:
	docker build -t fourcube/openiban:latest .