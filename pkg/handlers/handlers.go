package handlers

import (
	"context"
	"gocloud/pkg/api"
	"gocloud/pkg/playlist"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"strconv"
	"time"
)

type ServerPlaylist struct {
	api.UnimplementedPlaylistServer
	Controller *playlist.Player
}

func (s ServerPlaylist) Play(ctx context.Context, empty *api.Empty) (*api.Response, error) {
	err := s.Controller.Play()
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}
	song, err := s.Controller.CurrentSong()
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "internal server error")
	}
	resp := api.Response{Message: "PLay: " + song.Name}
	log.Println(resp.Message)
	return &resp, nil
}

func (s ServerPlaylist) Pause(ctx context.Context, empty *api.Empty) (*api.Response, error) {

	err := s.Controller.PauseSong()
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}
	song, err := s.Controller.CurrentSong()
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "internal server error")
	}
	resp := api.Response{Message: "Pause: " + song.Name}
	log.Println(resp.Message)
	return &resp, nil
}

func (s ServerPlaylist) Next(ctx context.Context, empty *api.Empty) (*api.Response, error) {
	err := s.Controller.NextSong()
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.OutOfRange, err.Error())
	}
	resp := api.Response{Message: "turned on the next song"}
	log.Println(resp.Message)
	return &resp, nil
}

func (s ServerPlaylist) Prev(ctx context.Context, empty *api.Empty) (*api.Response, error) {
	err := s.Controller.PrevSong()
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.OutOfRange, err.Error())
	}
	resp := api.Response{Message: "turned on the previous song"}
	log.Println(resp.Message)
	return &resp, nil
}

func (s ServerPlaylist) AddSong(ctx context.Context, song *api.Song) (*api.Response, error) {
	newSong := playlist.Song{Name: song.Name}
	durationStr := strconv.Itoa(int(song.Duration)) + "s"
	durationTime, err := time.ParseDuration(durationStr)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "can't add song")
	}
	newSong.Duration = durationTime
	err = s.Controller.AddSong(newSong)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "can't add song")
	}
	resp := api.Response{Message: "add new song to the playlist: " + song.Name}
	log.Println(resp.Message)
	return &resp, nil
}

func (s ServerPlaylist) GetCurrentSong(ctx context.Context, empty *api.Empty) (*api.Response, error) {
	song, err := s.Controller.CurrentSong()
	if err != nil {
		log.Println(err)
		return &api.Response{Message: err.Error()}, nil
	}
	if song.IsPlaying == false {
		log.Println(err)
		return &api.Response{Message: "no song is playing now"}, nil
	}
	resp := api.Response{Message: "now playing is: " + song.Name}
	log.Println(resp.Message)
	return &resp, nil

}

func (s ServerPlaylist) DeleteSong(ctx context.Context, song *api.Song) (*api.Response, error) {
	targetSong := playlist.Song{Name: song.Name}
	durationStr := strconv.Itoa(int(song.Duration)) + "s"
	durationTime, err := time.ParseDuration(durationStr)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "can't delete song")
	}
	targetSong.Duration = durationTime
	err, stcode := s.Controller.DeleteSong(targetSong)
	if err != nil {
		if stcode == 400 {
			log.Println(err)
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		if stcode == 300 {
			log.Println(err)
			return nil, status.Error(codes.FailedPrecondition, err.Error())
		}
		if stcode == 500 {
			log.Println(err)
			return nil, status.Error(codes.Internal, "can't delete song")
		}
	}
	resp := api.Response{Message: "deleted: " + song.Name}
	log.Println(resp.Message)
	return &resp, nil
}

func (s ServerPlaylist) UpdateNextSong(ctx context.Context, song *api.Song) (*api.Response, error) {
	newSong := playlist.Song{Name: song.Name}
	durationStr := strconv.Itoa(int(song.Duration)) + "s"
	durationTime, err := time.ParseDuration(durationStr)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "can't update next song")
	}
	newSong.Duration = durationTime
	err = s.Controller.UpdateNextSong(newSong)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "can't update next song")
	}
	resp := api.Response{Message: "next song updated: " + song.Name}
	log.Println(resp.Message)
	return &resp, nil
}
