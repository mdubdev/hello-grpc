include version

generate:
	protoc hello.proto --go_out=plugins=grpc:/${HOME}/go/bin/

build: generate
	rm bin/*
	GOOS=linux go build -o bin/api main.go

run:
	export HELLOGRPC_HOST=:9090
	go run main.go

docker-build: build
	docker build --build-arg VERSION=${VERSION} -t mwilliamsdev/hello-grpc:${VERSION} .
	@echo "to push image execute:\ndocker push mwilliamsdev/hello-grpc:${VERSION}\n"

docker-run: docker-build
	docker stop hello-grpc || true && docker rm hello-grpc || true
	docker run --name hello-grpc -d -p 9090:9090 mwilliamsdev/hello-grpc

.PHONY: generate build run docker-build docker-run
