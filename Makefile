proto:
protoc -I=./pkg/proto --go_out=./pkg/proto --go_opt=paths=source_relative --go-grpc_out=./pkg/proto --go-grpc_opt=paths=source_relative ./pkg/proto/auth.proto

server:
	go run cmd/main.go