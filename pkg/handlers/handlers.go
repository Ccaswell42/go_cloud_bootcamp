package handlers

import (
	"context"
	"errors"
	"gocloud/pkg/api"
	"gocloud/pkg/playlist"
	"strconv"
	"time"
)

type ServerPlaylist struct {
	api.UnimplementedPlaylistServer
	Controller playlist.Player
}

func (s ServerPlaylist) Play(ctx context.Context, empty *api.Empty) (*api.Response, error) {
	err := s.Controller.Play()
	if err != nil {
		return &api.Response{Message: err.Error(), StatusCode: 400}, err
	}
	song, err := s.Controller.CurrentSong()
	if err != nil {
		return &api.Response{Message: err.Error(), StatusCode: 500}, err
	}
	return &api.Response{Message: "PLay: " + song.Name, StatusCode: 200}, nil
}

func (s ServerPlaylist) Pause(ctx context.Context, empty *api.Empty) (*api.Response, error) {

	err := s.Controller.PauseSong()
	if err != nil {
		return &api.Response{Message: err.Error(), StatusCode: 400}, err
	}
	song, err := s.Controller.CurrentSong()
	if err != nil {
		return &api.Response{Message: err.Error(), StatusCode: 500}, err
	}
	return &api.Response{Message: "Pause: " + song.Name, StatusCode: 200}, nil
}

func (s ServerPlaylist) Next(ctx context.Context, empty *api.Empty) (*api.Response, error) {
	err := s.Controller.NextSong()
	if err != nil {
		return &api.Response{Message: err.Error(), StatusCode: 400}, err
	}
	song, err := s.Controller.CurrentSong()
	if err != nil {
		return &api.Response{Message: err.Error(), StatusCode: 500}, err
	}
	return &api.Response{Message: "turned on the next song, now playing is: " + song.Name, StatusCode: 200}, nil
}

func (s ServerPlaylist) Prev(ctx context.Context, empty *api.Empty) (*api.Response, error) {
	err := s.Controller.PrevSong()
	if err != nil {
		return &api.Response{Message: err.Error(), StatusCode: 400}, err
	}
	song, err := s.Controller.CurrentSong()
	if err != nil {
		return &api.Response{Message: err.Error(), StatusCode: 500}, err
	}
	return &api.Response{Message: "turned on the previous song, now playing is: " + song.Name, StatusCode: 200}, nil
}

func (s ServerPlaylist) AddSong(ctx context.Context, song *api.Song) (*api.Response, error) {
	newSong := playlist.Song{Name: song.Name}
	durationStr := strconv.Itoa(int(song.Duration)) + "s"
	durationTime, err := time.ParseDuration(durationStr)
	if err != nil {
		return &api.Response{Message: "can't add song", StatusCode: 500}, errors.New("can't add song")
	}
	newSong.Duration = durationTime
	s.Controller.AddSong(newSong)
	return &api.Response{Message: "add new song to the playlist: " + song.Name, StatusCode: 200}, nil
}
