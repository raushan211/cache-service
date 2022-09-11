

gen:
	protoc --proto_path=proto proto/*.proto --go_out=server --go-grpc_out=server
	protoc --proto_path=proto proto/*.proto --go_out=client --go-grpc_out=client

clean:
	rm -rf server/pb/
	rm -rf client/pb/

.PHONY: run server
server:
	go run server/main.go redis

.PHONY: run client
client:
	go run client/main.go

install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
	brew install protobuf
	brew install clang-format
	brew install grpcurl
	export GO_PATH=~/go 
	export PATH=$PATH:/$GO_PATH/bin
	go mod vendor
path:
	export GO_PATH=~/go 
	export PATH=$PATH:/$GO_PATH/bin

test:
	rm -rf tmp && mkdir tmp
	go test -cover -race serializer/*.go