
install:
	go get google.golang.org/grpc

build-protocol:
	protoc --proto_path=proto/ --go_out=plugins=grpc:service proto/*.proto
	grpc_tools_node_protoc --proto_path=proto/ --js_out=import_style=commonjs,binary:node-client --grpc_out=node-client --plugin=grpc_tools_node_protoc_plugin proto/GrpcDemoService.proto

build:
	build-protocol
	go build

