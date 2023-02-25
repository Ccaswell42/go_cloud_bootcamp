package playlist

import (
	"container/list"
	"errors"
	"github.com/ivahaev/timer"
)

func (p *Player) Play() error {
	p.mu.Lock()
	if p.FirstTrack.Len() == 0 {
		p.mu.Unlock()
		return errors.New("playlist is empty")
	}
	if p.NowPlaying == nil {
		p.NowPlaying = p.FirstTrack.Front()
	}
	/// if p.FirstTrack.Len() == 1 {
	// 		go p.PlaySong()
	//   }
	//
	//
	//
	//
	p.PlayChan <- struct{}{}
	p.mu.Unlock()

	return nil
}

func (p *Player) PrevSong() {

}

func (p *Player) PauseSong() {

}

func (p *Player) AddSong(song Song) {
	p.mu.Lock()
	p.FirstTrack.PushBack(song)
	p.mu.Unlock()
}
func (p *Player) DeleteSong(songTarget Song) error {
	var res *list.Element

	for e := p.FirstTrack.Front(); e != nil; e = e.Next() {
		p.mu.Lock()
		song := e.Value.(Song)
		p.mu.Unlock()
		if song.Name == songTarget.Name {
			res = e
			break
		}
	}
	p.mu.Lock()
	cast, ok := res.Value.(Song)
	p.mu.Unlock()
	if !ok || cast.Name != songTarget.Name {
		return errors.New("no such song")
	}
	p.mu.Lock()
	if res.Value.(Song).IsPlaying {
		p.mu.Unlock()
		return errors.New("can't delete, the song is playing now")
	}
	p.FirstTrack.Remove(res)
	p.mu.Unlock()

	/// проверить мьютексы
	return nil
}

func (p *Player) PlaySong(chPlay, chPause, chNext, chPrev chan struct{}) {
	p.mu.Lock()
	song := p.NowPlaying.Value.(Song)
	p.mu.Unlock()
	timerDuration := timer.NewTimer(song.Duration)
	timerDuration.Start()
	for {
		select {
		case <-chPause:
			timerDuration.Pause()
			song.IsPlaying = false
			p.mu.Lock()
			p.NowPlaying.Value = song
			p.mu.Unlock()
		case <-chPlay:
			timerDuration.Start()
			song.IsPlaying = true
			p.mu.Lock()
			p.NowPlaying.Value = song
			p.mu.Unlock()
		case <-timerDuration.C:
			return
		case <-chNext:
			return
		case <-chPrev:
			return
		}
	}
}

func (p *Player) Controller() {

	chPlay := make(chan struct{})
	chPause := make(chan struct{})
	chNext := make(chan struct{})
	chPrev := make(chan struct{})
	for {
		select {
		case <-p.PlayChan:
			go p.PlaySong(chPlay, chPause, chNext, chPrev)
			chPlay <- struct{}{}
		}

	}

}

func (p *Player) CurrentSong() Song {
	p.mu.Lock()
	song := p.NowPlaying.Value.(Song)
	p.mu.Unlock()
	return song
}

func (p *Player) UpdateNextSong(songTarget Song) {
	p.mu.Lock()
	p.FirstTrack.InsertAfter(songTarget, p.NowPlaying)
	p.mu.Unlock()
}
