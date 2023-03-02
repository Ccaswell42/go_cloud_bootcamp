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
	Controller *playlist.Player
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

	return &api.Response{Message: "turned on the next song", StatusCode: 200}, nil
}

func (s ServerPlaylist) Prev(ctx context.Context, empty *api.Empty) (*api.Response, error) {
	err := s.Controller.PrevSong()
	if err != nil {
		return &api.Response{Message: err.Error(), StatusCode: 400}, err
	}
	return &api.Response{Message: "turned on the previous song", StatusCode: 200}, nil
}

func (s ServerPlaylist) AddSong(ctx context.Context, song *api.Song) (*api.Response, error) {
	newSong := playlist.Song{Name: song.Name}
	durationStr := strconv.Itoa(int(song.Duration)) + "s"
	durationTime, err := time.ParseDuration(durationStr)
	if err != nil {
		return &api.Response{Message: "can't add song", StatusCode: 500}, errors.New("can't add song")
	}
	newSong.Duration = durationTime
	err = s.Controller.AddSong(newSong)
	if err != nil {
		return &api.Response{Message: "can't add song", StatusCode: 500}, errors.New("can't add song")
	}
	return &api.Response{Message: "add new song to the playlist: " + song.Name, StatusCode: 200}, nil
}

func (s ServerPlaylist) GetCurrentSong(ctx context.Context, empty *api.Empty) (*api.Response, error) {
	song, err := s.Controller.CurrentSong()
	if err != nil {
		return &api.Response{Message: "no song is playing now", StatusCode: 200}, nil
	}
	if song.IsPlaying == false {
		return &api.Response{Message: "no song is playing now", StatusCode: 200}, nil
	}
	return &api.Response{Message: "now playing is: " + song.Name, StatusCode: 200}, nil

}

func (s ServerPlaylist) DeleteSong(ctx context.Context, song *api.Song) (*api.Response, error) {
	targetSong := playlist.Song{Name: song.Name}
	durationStr := strconv.Itoa(int(song.Duration)) + "s"
	durationTime, err := time.ParseDuration(durationStr)
	if err != nil {
		return &api.Response{Message: "can't delete song", StatusCode: 500}, errors.New("can't delete song")
	}
	targetSong.Duration = durationTime
	err, status := s.Controller.DeleteSong(targetSong)
	if err != nil {
		if status == 400 {
			return &api.Response{Message: err.Error(), StatusCode: 400}, err
		}
		if status == 500 {
			return &api.Response{Message: "can't delete song", StatusCode: 500}, errors.New("can't delete song")
		}
	}
	return &api.Response{Message: "deleted: " + song.Name, StatusCode: 200}, nil
}

func (s ServerPlaylist) UpdateNextSong(ctx context.Context, song *api.Song) (*api.Response, error) {
	newSong := playlist.Song{Name: song.Name}
	durationStr := strconv.Itoa(int(song.Duration)) + "s"
	durationTime, err := time.ParseDuration(durationStr)
	if err != nil {
		return &api.Response{Message: "can't update next song", StatusCode: 500}, errors.New("can't update next song")
	}
	newSong.Duration = durationTime
	err = s.Controller.UpdateNextSong(newSong)
	if err != nil {
		return &api.Response{Message: "can't update next song", StatusCode: 500}, errors.New("can't update next song")
	}
	return &api.Response{Message: "next song updated: " + song.Name, StatusCode: 200}, nil
}
