include .env
include version

generate:
	-rm bin/*.pb.go
	-rm -r /tmp/hello-grpc
	-mkdir /tmp/hello-grpc

	protoc hello.proto --go_out=plugins=grpc:/tmp/hello-grpc/
	cp /tmp/hello-grpc/hello.pb.go bin/

.PHONY generate

build: generate
	-rm -r bin
	mkdir -p bin

	GOOS=linux CGO_ENABLED=0 go build -o bin/api main.go

.PHONY build

run:
	HELLOGRPC_HOST=:9090
	go run main.go

.PHONY run
