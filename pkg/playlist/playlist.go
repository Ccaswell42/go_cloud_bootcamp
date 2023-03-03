package playlist

import (
	"container/list"
	"gocloud/pkg/storage"
	"sync"
	"time"
)

type Player struct {
	songList   *list.List
	nowPlaying *list.Element
	mutex      *sync.Mutex
	playChan   chan struct{}
	pauseChan  chan struct{}
	nextChan   chan struct{}
	prevChan   chan struct{}
	flagPrev   bool
	repo       *storage.Repository
}

type Song struct {
	Name      string
	Duration  time.Duration
	IsPlaying bool
}

func New(repo *storage.Repository) (*Player, error) {
	Playlist := Player{
		nowPlaying: nil,
		mutex:      &sync.Mutex{},
		playChan:   make(chan struct{}),
		pauseChan:  make(chan struct{}),
		nextChan:   make(chan struct{}),
		prevChan:   make(chan struct{}),
		repo:       repo,
	}
	songList, err := Playlist.repo.GetAll()
	if err != nil {
		Playlist.songList = list.New()
		return &Playlist, err
	}
	for e := songList.Front(); e != nil; e = e.Next() {
		songDB := e.Value.(storage.Song)
		song := Song{Name: songDB.Name, Duration: songDB.Duration, IsPlaying: false}
		e.Value = song
	}

	Playlist.songList = songList
	return &Playlist, nil
}
