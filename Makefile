


run:
	go run main.go

server:
	go run main.go -server grpc

gen:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    grpc/template.proto