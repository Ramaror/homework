.PHONY: run
run:
	go run C:\Users\doksh\Downloads\homework\cmd\bot\main.go
build:
	go build -o bin/bot C:\Users\doksh\Downloads\homework\cmd\bot\main.go

LOCAL_BIN:="C:/Users/doksh/Downloads/homework/bin"
.PHONY: .deps
.deps:
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc