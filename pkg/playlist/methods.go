package playlist

import (
	"container/list"
	"errors"
	"github.com/ivahaev/timer"
)

func (p *Player) NextSong() {

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
	cast, ok := res.Value.(Song)
	if !ok || cast.Name != songTarget.Name {
		return errors.New("no such song")
	}
	if res.Value.(Song).IsPlaying {
		return errors.New("can't delete, the song is playing now")
	}
	p.mu.Lock()
	p.FirstTrack.Remove(res)
	p.mu.Unlock()

	/// проверить мьютексы
	return nil
}

func (p *Player) PlaySong() {
	p.mu.Lock()
	song := p.NowPlaying.Value.(Song)
	p.mu.Unlock()
	timerDuration := timer.NewTimer(song.Duration)
	timerDuration.Start()
	for {
		select {
		case <-p.Pause:
			timerDuration.Pause()
			song.IsPlaying = false
			p.mu.Lock()
			p.NowPlaying.Value = song
			p.mu.Unlock()
		case <-p.Play:
			timerDuration.Start()
			song.IsPlaying = true
			p.mu.Lock()
			p.NowPlaying.Value = song
			p.mu.Unlock()
		case <-timerDuration.C:
			return
		}
	}
}

func (p *Player) Playing() {
	e := p.FirstTrack.Front()
	for e := p.FirstTrack.Front(); e != nil; e = e.Next() {
		// do something with e.Value
	}

	for {
		if e == nil {

		}

	}

}
