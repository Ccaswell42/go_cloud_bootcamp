PROTO = api/proto/service.proto


build:
	docker-compose build

up:
	docker-compose up

genpb:
	protoc --go_out=. --go_opt=paths=import  --go-grpc_out=. --go-grpc_opt=paths=import $(PROTO)

cli:
	go run client/main.go

all:
	go run cmd/main.go
