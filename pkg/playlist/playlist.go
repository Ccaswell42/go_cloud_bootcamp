package playlist

import (
	"container/list"
	"sync"
	"time"
)

var Playlist Player

type Player struct {
	FirstTrack *list.List
	NowPlaying *list.Element
	mu         *sync.Mutex
	PlayChan   chan struct{}
	PauseChan  chan struct{}
	NextChan   chan struct{}
	PrevChan   chan struct{}
}

type Song struct {
	Name      string
	Duration  time.Duration
	IsPlaying bool
}

func init() {

	Playlist = Player{FirstTrack: list.New(),
		NowPlaying: nil,
		mu:         &sync.Mutex{},
		PlayChan:   make(chan struct{}),
		PauseChan:  make(chan struct{}),
		NextChan:   make(chan struct{}),
		PrevChan:   make(chan struct{}),
	}
	//go Playlist.Controller()

}
