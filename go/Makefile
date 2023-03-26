protoc_install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 && \
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 && \
	export PATH="$PATH:$(go env GOPATH)/bin"

protoc_generate:
	protoc --go_out=./tasks --go_opt=paths=source_relative \
    --go-grpc_out=./tasks --go-grpc_opt=paths=source_relative \
    tasks.proto

start:
	go run ./
