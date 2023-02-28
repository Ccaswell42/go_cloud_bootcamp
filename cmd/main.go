package main

import (
	"fmt"
	"gocloud/pkg/playlist"
	"sync"
)

func main() {
	controller := playlist.Player{NowPlaying: nil, Mu: &sync.Mutex{}}
	anys, err := controller.CurrentSong()
	fmt.Println(anys, err)
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
