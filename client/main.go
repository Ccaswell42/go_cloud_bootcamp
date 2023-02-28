package main

import (
	"context"
	"fmt"
	"gocloud/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func RequestTest() {
	grcpConn, err := grpc.Dial(
		"0.0.0.0:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("can't connect to grpc")
	}
	defer grcpConn.Close()
	c := api.NewPlaylistClient(grcpConn)
	resp, _ := c.Play(context.Background(), &api.Empty{})
	fmt.Println(resp.SongName)
}

func main() {
	RequestTest()

}
