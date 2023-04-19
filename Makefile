proto:
	protoc pkg/pb/*.proto --go_out=. --go-grpc_out=. $(find . -name '*.proto')

server:
	go run cmd/main.go