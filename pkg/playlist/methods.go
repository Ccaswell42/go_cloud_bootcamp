package playlist

import (
	"container/list"
	"errors"
	"github.com/ivahaev/timer"
	"gocloud/pkg/storage"
	"log"
)

func (p *Player) Play() error {
	p.mutex.Lock()
	if p.songList.Len() == 0 {
		p.mutex.Unlock()
		return errors.New("playlist is empty")
	}
	if p.nowPlaying == nil {
		p.nowPlaying = p.songList.Front()
		go p.PlaySong()
	}
	p.mutex.Unlock()
	p.playChan <- struct{}{}
	return nil
}

func (p *Player) PrevSong() error {
	p.mutex.Lock()
	if p.nowPlaying == nil {
		p.mutex.Unlock()
		return errors.New("playlist is empty")
	}
	if p.nowPlaying.Prev() == nil {
		p.mutex.Unlock()
		return errors.New("now playing is first song")
	}
	p.prevChan <- struct{}{}
	p.mutex.Unlock()
	return nil
}

func (p *Player) PauseSong() error {
	p.mutex.Lock()
	if p.nowPlaying == nil {
		p.mutex.Unlock()
		return errors.New("playlist is empty")
	}
	if p.nowPlaying.Value.(Song).IsPlaying == false {
		p.mutex.Unlock()
		return errors.New("no song is playing now")
	}
	song := p.nowPlaying.Value.(Song)
	song.IsPlaying = false
	p.nowPlaying.Value = song
	p.mutex.Unlock()
	p.pauseChan <- struct{}{}

	return nil
	//смотри мьютексы и nowplaying на nil
}
func (p *Player) NextSong() error {

	p.mutex.Lock()
	/////// тут не nowplaying а p.firsttrack.frpnt()
	if p.nowPlaying == nil {
		p.mutex.Unlock()
		return errors.New("playlist is empty")
	}
	if p.nowPlaying.Next() == nil {
		p.mutex.Unlock()
		return errors.New("now playing is last song")
	}
	p.nextChan <- struct{}{}
	p.mutex.Unlock()
	return nil
	//протестировать конкурентность в горутинах и узнать нужен ли мьютекс
	// в сабже написано что эта функция должна быть с учетом конкурентности, значит запись в канал
	//надо облажить мьютексами, а значит все записи в каналы надо обложить мьютексами
	// затесть это!
}

func (p *Player) AddSong(song Song) error {
	songDB := storage.Song{Name: song.Name, Duration: song.Duration}
	err := p.repo.Set(songDB)
	if err != nil {
		return err
	}
	p.mutex.Lock()
	p.songList.PushBack(song)
	p.mutex.Unlock()
	return nil
}
func (p *Player) DeleteSong(songTarget Song) (error, int) {
	var res *list.Element

	p.mutex.Lock()
	for e := p.songList.Front(); e != nil; e = e.Next() {
		song := e.Value.(Song)
		if song.Name == songTarget.Name {
			res = e
			break
		}
	}
	if res == nil {
		p.mutex.Unlock()
		return errors.New("no such song"), 400
	}
	cast, ok := res.Value.(Song)
	if !ok || cast.Name != songTarget.Name {
		p.mutex.Unlock()
		return errors.New("no such song"), 400
	}
	if res.Value.(Song).IsPlaying {
		p.mutex.Unlock()
		return errors.New("can't delete, the song is playing now"), 400
	}
	songDB := storage.Song{Name: songTarget.Name, Duration: songTarget.Duration}
	err := p.repo.DeleteSong(songDB)
	if err != nil {
		return errors.New("can't delete from DB"), 500
	}
	p.songList.Remove(res)
	p.mutex.Unlock()

	/// проверить мьютексы!!!!!! тут с ними очень все плохо. и может вообще
	/// выкинуть мьютексы и оставить их в хендлерах
	return nil, 200
}

func (p *Player) PlaySong() {
	//затестить мьютексы
	for e := p.songList.Front(); e != nil; e = e.Next() {
		p.mutex.Lock()
		if p.flagPrev {
			p.flagPrev = false
			e = e.Prev().Prev()
		}
		p.nowPlaying = e
		log.Println(p.nowPlaying)
		song := p.nowPlaying.Value.(Song)
		song.IsPlaying = true
		p.nowPlaying.Value = song
		timerDuration := timer.NewTimer(song.Duration)
		timerDuration.Start()
		p.mutex.Unlock()
	LOOP:
		for {
			select {
			case <-p.playChan:
				timerDuration.Start()
				song.IsPlaying = true
				p.mutex.Lock()
				p.nowPlaying.Value = song
				p.mutex.Unlock()
			case <-p.pauseChan:
				timerDuration.Pause()
				song.IsPlaying = false
				p.mutex.Lock()
				p.nowPlaying.Value = song
				p.mutex.Unlock()
			case <-timerDuration.C:
				song.IsPlaying = false
				p.mutex.Lock()
				p.nowPlaying.Value = song
				p.mutex.Unlock()
				break LOOP
			case <-p.nextChan:
				song.IsPlaying = false
				p.mutex.Lock()
				p.nowPlaying.Value = song
				p.mutex.Unlock()
				break LOOP
			case <-p.prevChan:
				song.IsPlaying = false
				p.mutex.Lock()
				p.flagPrev = true
				p.nowPlaying.Value = song
				p.mutex.Unlock()
				break LOOP
			}
		}
	}

	p.mutex.Lock()
	p.nowPlaying = nil
	p.mutex.Unlock()
}

func (p *Player) CurrentSong() (Song, error) {
	song := Song{}
	p.mutex.Lock()
	if p.nowPlaying == nil {
		p.mutex.Unlock()
		return song, errors.New("no song is playing now")
	}
	song = p.nowPlaying.Value.(Song)
	p.mutex.Unlock()
	return song, nil
}

func (p *Player) UpdateNextSong(song Song) error {
	songDB := storage.Song{Name: song.Name, Duration: song.Duration}
	err := p.repo.Set(songDB)
	if err != nil {
		return err
	}
	p.mutex.Lock()
	if p.nowPlaying == nil {
		p.songList.PushBack(song)
		p.mutex.Unlock()
		return nil
	}
	p.songList.InsertAfter(song, p.nowPlaying)
	p.mutex.Unlock()
	//проверить на ошибки в случае если nowPlaying == nil
	// можно это проверить на более верхнем уровне
	return nil
}
