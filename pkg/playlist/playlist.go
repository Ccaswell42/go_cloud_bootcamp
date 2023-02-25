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

	//songa := Song{Name: "za dengi da", Duration: time.Second * 60}

	Playlist = Player{FirstTrack: list.New()}
	go Playlist.Controller()

}
