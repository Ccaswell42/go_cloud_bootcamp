package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"gocloud/config"
	"gocloud/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"time"
)

const host = "0.0.0.0:"

func Play(c api.PlaylistClient) {
	resp, err := c.Play(context.Background(), &api.Empty{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("PLAY:", resp.Message)
	}

}
func NextSong(c api.PlaylistClient) {
	resp, err := c.Next(context.Background(), &api.Empty{})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("NEXT:", resp.Message)
	}
}

func Pause(c api.PlaylistClient) {
	resp, err := c.Pause(context.Background(), &api.Empty{})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Pause:", resp.Message)
	}

}

func PrevSong(c api.PlaylistClient) {
	resp, err := c.Prev(context.Background(), &api.Empty{})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("PREV:", resp.Message)
	}

}

func AddSong(c api.PlaylistClient, song *api.Song) {
	resp, err := c.AddSong(context.Background(), song)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("ADD SONG:", resp.Message)
	}

}

func DeleteSong(c api.PlaylistClient, song *api.Song) {
	resp, err := c.DeleteSong(context.Background(), song)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("DELETE:", resp.Message)
	}
}

func GetCurrentSong(c api.PlaylistClient) {

	resp, err := c.GetCurrentSong(context.Background(), &api.Empty{})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("CURRENT SONG:", resp.Message)
	}

}

func UpdateSong(c api.PlaylistClient, song *api.Song) {
	resp, err := c.UpdateNextSong(context.Background(), song)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("UPDATE:", resp.Message)
	}
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
	ReadlLine(c)
}

func ReadlLine(c api.PlaylistClient) {
	fmt.Println("client starts")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "exit" {
			break
		}
		switch scanner.Text() {
		case "play":
			Play(c)
		case "pause":
			Pause(c)
		case "next":
			NextSong(c)
		case "prev":
			PrevSong(c)
		case "get":
			GetCurrentSong(c)
		case "add":
			song, err := ParseSong(scanner)
			if err == nil {
				AddSong(c, song)
			}
		case "delete":
			song, err := ParseSong(scanner)
			if err == nil {
				DeleteSong(c, song)
			}
		case "update":
			song, err := ParseSong(scanner)
			if err == nil {
				UpdateSong(c, song)
			}

		}
	}
}

func ParseSong(sc *bufio.Scanner) (*api.Song, error) {
	fmt.Println("enter song name")
	var name, duration string
	sc.Scan()
	name = sc.Text()
	fmt.Println("enter song duration in format:", "1h10m10s")
	sc.Scan()
	duration = sc.Text()
	durTime, err := time.ParseDuration(duration)
	if err != nil || durTime < 0 {
		fmt.Println("CLI: failed to read duration")
		return nil, errors.New("failed read duration")
	}
	return &api.Song{Name: name, Duration: uint64(durTime.Seconds())}, nil
}
