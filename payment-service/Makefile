build-proto:
	protoc --go_out=./app/ --go-grpc_out=./app/ proto/payment.proto
	protoc --go_out=./app/ --go-grpc_out=./app/ proto/echo.proto
build:
	make build-proto
	go build -o bin/app cmd/main.go
run:
	go run cmd/main.go

run-client:
	go run client/client.go