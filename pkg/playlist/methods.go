package playlist

import (
	"container/list"
	"errors"
	"github.com/ivahaev/timer"
)

func (p *Player) Play() error {
	p.Mu.Lock()
	if p.FirstTrack.Len() == 0 {
		p.Mu.Unlock()
		return errors.New("playlist is empty")
	}
	if p.NowPlaying == nil {
		p.NowPlaying = p.FirstTrack.Front()
		//go p.PlaySong()
	}
	p.PlayChan <- struct{}{}
	p.Mu.Unlock()

	return nil
}

func (p *Player) PrevSong() error {
	p.Mu.Lock()
	if p.NowPlaying == nil {
		p.Mu.Unlock()
		return errors.New("playlist is empty")
	}
	if p.NowPlaying.Prev() == nil {
		p.Mu.Unlock()
		return errors.New("now playing is first song")
	}
	p.PrevChan <- struct{}{}
	p.Mu.Unlock()
	return nil
}

func (p *Player) PauseSong() error {
	p.Mu.Lock()
	if p.NowPlaying == nil {
		p.Mu.Unlock()
		return errors.New("playlist is empty")
	}
	if p.NowPlaying.Value.(Song).IsPlaying == false {
		p.Mu.Unlock()
		return errors.New("no song is playing now")
	}
	p.PauseChan <- struct{}{}
	p.Mu.Unlock()
	return nil
	//смотри мьютексы и nowplaying на nil
}
func (p *Player) NextSong() error {

	p.Mu.Lock()
	if p.NowPlaying == nil {
		p.Mu.Unlock()
		return errors.New("playlist is empty")
	}
	if p.NowPlaying.Next() == nil {
		p.Mu.Unlock()
		return errors.New("now playing is last song")
	}
	p.NextChan <- struct{}{}
	p.Mu.Unlock()
	return nil
	//протестировать конкурентность в горутинах и узнать нужен ли мьютекс
	// в сабже написано что эта функция должна быть с учетом конкурентности, значит запись в канал
	//надо облажить мьютексами, а значит все записи в каналы надо обложить мьютексами
	// затесть это!
}

func (p *Player) AddSong(song Song) {
	p.Mu.Lock()
	p.FirstTrack.PushBack(song)
	p.Mu.Unlock()
}
func (p *Player) DeleteSong(songTarget Song) error {
	var res *list.Element

	for e := p.FirstTrack.Front(); e != nil; e = e.Next() {
		p.Mu.Lock()
		song := e.Value.(Song)
		p.Mu.Unlock()
		if song.Name == songTarget.Name {
			res = e
			break
		}
	}
	p.Mu.Lock()
	cast, ok := res.Value.(Song)
	p.Mu.Unlock()
	if !ok || cast.Name != songTarget.Name {
		return errors.New("no such song")
	}
	p.Mu.Lock()
	if res.Value.(Song).IsPlaying {
		p.Mu.Unlock()
		return errors.New("can't delete, the song is playing now")
	}
	p.FirstTrack.Remove(res)
	p.Mu.Unlock()

	/// проверить мьютексы
	return nil
}

//func (p *Player) PlaySong(chPlay, chPause, chNext, chPrev chan struct{}) {
//	p.mu.Lock()
//	song := p.NowPlaying.Value.(Song)
//	p.mu.Unlock()
//	timerDuration := timer.NewTimer(song.Duration)
//	timerDuration.Start()
//	for {
//		select {
//		case <-chPause:
//			timerDuration.Pause()
//			song.IsPlaying = false
//			p.mu.Lock()
//			p.NowPlaying.Value = song
//			p.mu.Unlock()
//		case <-chPlay:
//			timerDuration.Start()
//			song.IsPlaying = true
//			p.mu.Lock()
//			p.NowPlaying.Value = song
//			p.mu.Unlock()
//		case <-timerDuration.C:
//			return
//		case <-chNext:
//			return
//		case <-chPrev:
//			return
//		}
//	}
//}

func (p *Player) PlaySong() {
	//затестить мьютексы
	for e := p.FirstTrack.Front(); e != nil; e = e.Next() {
		p.Mu.Lock()
		p.NowPlaying = e
		song := p.NowPlaying.Value.(Song)
		p.Mu.Unlock()
		timerDuration := timer.NewTimer(song.Duration)
		for {
			select {
			case <-p.PlayChan:
				timerDuration.Start()
				song.IsPlaying = true
				p.Mu.Lock()
				p.NowPlaying.Value = song
				p.Mu.Unlock()
			case <-p.PauseChan:
				timerDuration.Pause()
				song.IsPlaying = false
				p.Mu.Lock()
				p.NowPlaying.Value = song
				p.Mu.Unlock()
			case <-timerDuration.C:
				break
			case <-p.NextChan:
				break
			case <-p.PrevChan:
				e = e.Prev().Prev()
				break
			}

		}
	}
}

//func (p *Player) Controller() {
//
//	chPlay := make(chan struct{})
//	chPause := make(chan struct{})
//	chNext := make(chan struct{})
//	chPrev := make(chan struct{})
//	for {
//		select {
//		case <-p.PlayChan:
//			go p.PlaySong(chPlay, chPause, chNext, chPrev)
//			chPlay <- struct{}{}
//		}
//
//	}
//
//}

func (p *Player) CurrentSong() (Song, error) {
	song := Song{}
	p.Mu.Lock()
	if p.NowPlaying == nil {
		p.Mu.Unlock()
		return song, errors.New("no song is playing now")
	}
	song = p.NowPlaying.Value.(Song)
	p.Mu.Unlock()
	return song, nil
}

func (p *Player) UpdateNextSong(song Song) {
	p.Mu.Lock()
	if p.NowPlaying == nil {
		p.FirstTrack.PushBack(song)
		p.Mu.Unlock()
		return
	}
	p.FirstTrack.InsertAfter(song, p.NowPlaying)
	p.Mu.Unlock()
	//проверить на ошибки в случае если nowPlaying == nil
	// можно это проверить на более верхнем уровне
}
