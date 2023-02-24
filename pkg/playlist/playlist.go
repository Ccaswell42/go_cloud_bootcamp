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
	Play       chan struct{}
	Pause      chan struct{}
	Next       chan struct{}
	Prev       chan struct{}
}

type Song struct {
	Name      string
	Duration  time.Duration
	Stop      time.Duration
	IsPlaying bool
}

func init() {

	//songa := Song{Name: "za dengi da", Duration: time.Second * 60}

	Playlist = Player{FirstTrack: list.New(), Stop: false}
}
