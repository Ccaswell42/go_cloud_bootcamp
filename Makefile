PROTO = api/proto/service.proto


all:
	go run cmd/main.go

genpb:
	protoc --go_out=. --go_opt=paths=import  --go-grpc_out=. --go-grpc_opt=paths=import $(PROTO)

client:
	go run client/main.go

doc:  compose
	docker-compose up

compose:
	docker-compose build