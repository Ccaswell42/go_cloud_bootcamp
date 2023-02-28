package main

import (
	"context"
	"fmt"
	"gocloud/pkg/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	obj := Obj{}
	server := grpc.NewServer()
	api.RegisterPlaylistServer(server, obj)
	fmt.Println("starting server at :8081")

	lis, err := net.Listen("tcp", "0.0.0.0:8081")
	if err != nil {
		log.Fatalln("cant listen port", err)
	}
	go server.Serve(lis)
	fmt.Scanln()
}

type Obj struct {
	api.UnimplementedPlaylistServer
}

func (o Obj) Play(ctx context.Context, empt *api.Empty) (*api.PlayMessage, error) {
	testo := api.PlayMessage{SongName: "diss"}
	return &testo, nil
}
