.PHONY: proto

proto:
	protoc --proto_path=./pkg/auth/pb --go_out=. --go-grpc_out=. ./pkg/auth/pb/*.proto
	protoc --proto_path=./pkg/menu/pb --go_out=. --go-grpc_out=. ./pkg/menu/pb/*.proto
	protoc --proto_path=./pkg/order/pb --go_out=. --go-grpc_out=. ./pkg/order/pb/*.proto

server:
	go run cmd/main.go