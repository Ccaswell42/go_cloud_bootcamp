package main

import (
	"fmt"
	"gocloud/config"
	"gocloud/pkg/api"
	"gocloud/pkg/handlers"
	"gocloud/pkg/playlist"
	"gocloud/pkg/storage"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	host    = "0.0.0.0:"
	network = "tcp"
)

func main() {
	conf, err := config.GetConfig()
	if err != nil {
		log.Fatal("config:", err)
	}
	repo, err := storage.ConnectToDb(conf)
	if err != nil {
		log.Fatal("connect to DB:", err)
	}

	controller, err := playlist.New(repo)
	if err != nil {
		log.Println(err)
	}
	srvPlaylist := handlers.ServerPlaylist{Controller: controller}
	lis, err := net.Listen(network, host+conf.Port)
	if err != nil {
		log.Fatal("can't listen port:", err)
	}
	server := grpc.NewServer()
	api.RegisterPlaylistServer(server, srvPlaylist)

	go func() {
		fmt.Println("starting server at ", conf.Port)
		err := server.Serve(lis)
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
		fmt.Println("server completed successfully")
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	server.GracefulStop()
	log.Println("Shutdown completed successfully")
}
