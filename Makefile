PROTO = api/proto/service.proto
genpb:
	protoc --go_out=. --go_opt=paths=import  --go-grpc_out=. --go-grpc_opt=paths=import $(PROTO)

all:
	go run cmd/main.go