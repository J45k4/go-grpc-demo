
install:
	go get google.golang.org/grpc

build-protocol:
	protoc --proto_path=proto/ --go_out=plugins=grpc:service proto/*.proto

build:
	build-protocol
	go build

