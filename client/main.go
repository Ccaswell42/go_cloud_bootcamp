package main

import (
	"context"
	"fmt"
	"gocloud/config"
	"gocloud/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const host = "0.0.0.0"

func Play(c api.PlaylistClient) {
	resp, err := c.Play(context.Background(), &api.Empty{})
	fmt.Println("PLAY:", resp, err)
}
func NextSong(c api.PlaylistClient) {
	resp, err := c.Next(context.Background(), &api.Empty{})
	fmt.Println("NEXT:", resp, err)
}

func Pause(c api.PlaylistClient) {
	resp, err := c.Pause(context.Background(), &api.Empty{})
	fmt.Println("Pause:", resp, err)
}

func PrevSong(c api.PlaylistClient) {
	resp, err := c.Prev(context.Background(), &api.Empty{})
	fmt.Println("PREV:", resp, err)
}

func AddSong(c api.PlaylistClient, song *api.Song) {
	resp, err := c.AddSong(context.Background(), song)
	fmt.Println("ADDSONG:", resp, err)
}

func DeleteSong(c api.PlaylistClient, song *api.Song) {
	resp, err := c.DeleteSong(context.Background(), song)
	fmt.Println("DELETE::", resp, err)
}

func GetCurrentSong(c api.PlaylistClient) {
	for {
		resp, err := c.GetCurrentSong(context.Background(), &api.Empty{})
		fmt.Println("CURRENT SONG:", resp, err)
		time.Sleep(1 * time.Second)
	}
}

func UpdateSong(c api.PlaylistClient, song *api.Song) {
	resp, err := c.UpdateNextSong(context.Background(), song)
	fmt.Println("UPDATE:", resp, err)

}

func main() {
	conf, err := config.GetConfig()
	if err != nil {
		log.Fatal("config", err)
	}

	grpcConn, err := grpc.Dial(
		host+conf.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("can't connect to grpc")
	}
	defer grpcConn.Close()
	c := api.NewPlaylistClient(grpcConn)
	Play(c)
	//Pause(c)
	//GetCurrentSong(c)
	//NextSong(c)
	//PrevSong(c)
	//UpdateSong(c, &api.Song{Name: "griby", Duration: 11})
	//
	//AddSong(c, &api.Song{Name: "griby", Duration: 11})
	//DeleteSong(c, &api.Song{Name: "oxxxymiron", Duration: 300})
}
