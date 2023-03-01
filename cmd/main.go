package main

import (
	"fmt"
	"gocloud/config"
	"gocloud/pkg/playlist"
	"gocloud/pkg/storage"
	"log"
	"time"
)

func main() {
	conf, err := config.GetConfig()
	if err != nil {
		log.Fatal("config", err)
	}
	repo, err := storage.ConnectToDb(conf)
	if err != nil {
		log.Fatal("connect", err)
	}

	//song1 := playlist.Song{Name: "basta", Duration: 14 * time.Second, IsPlaying: false}
	//song2 := playlist.Song{Name: "atl", Duration: 12 * time.Second, IsPlaying: false}
	//song3 := playlist.Song{Name: "srcip", Duration: 17 * time.Second, IsPlaying: false}
	//
	//repo.Set(song3)
	//repo.Set(song2)
	//repo.Set(song1)

	list, err := repo.GetAll()
	if err != nil {
		log.Println(err)
	}

	controller := playlist.Playlist
	controller.FirstTrack = list
	for e := controller.FirstTrack.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	err = controller.Play()
	if err != nil {
		log.Println("play", err)
	}
	fmt.Println("---------------------")

	go func() {
		for {
			song, err := controller.CurrentSong()
			if err != nil {
				log.Println(err)
			}
			fmt.Println(song)
			time.Sleep(1 * time.Second)
		}

	}()
	//time.Sleep(2 * time.Second)
	go func() {
		err := controller.PauseSong()
		if err != nil {
			log.Println("pause err", err)
		}
		//err = controller.Play()
		//if err != nil {
		//	log.Println("pause err", err)
		//}
		//time.Sleep(100 * time.Millisecond)
		err = controller.PauseSong()
		if err != nil {
			log.Println("pause err", err)
		}

	}()
	//go func() {
	//	time.Sleep(2 * time.Second)
	//	controller.NextSong()
	//	time.Sleep(2 * time.Second)
	//	controller.NextSong()
	//	time.Sleep(2 * time.Second)
	//	controller.NextSong()
	//	time.Sleep(2 * time.Second)
	//	controller.NextSong()
	//	time.Sleep(2 * time.Second)
	//	controller.NextSong()
	//	time.Sleep(2 * time.Second)
	//	controller.NextSong()
	//	time.Sleep(2 * time.Second)
	//	controller.NextSong()
	//	time.Sleep(2 * time.Second)
	//	controller.NextSong()
	//	time.Sleep(2 * time.Second)
	//	controller.NextSong()
	//	time.Sleep(2 * time.Second)
	//
	//}()

	//go func() {
	//	song1 := playlist.Song{Name: "makSim", Duration: 11 * time.Second, IsPlaying: false}
	//	song2 := playlist.Song{Name: "ranetki", Duration: 11 * time.Second, IsPlaying: false}
	//	controller.AddSong(song2)
	//	time.Sleep(2 * time.Second)
	//	controller.AddSong(song1)
	//
	//}()

	time.Sleep(18 * time.Second)

}

//obj := Obj{}
//server := grpc.NewServer()
//api.RegisterPlaylistServer(server, obj)
//fmt.Println("starting server at :8081")
//
//lis, err := net.Listen("tcp", "0.0.0.0:8081")
//if err != nil {
//	log.Fatalln("cant listen port", err)
//}
//go server.Serve(lis)
//fmt.Scanln()
