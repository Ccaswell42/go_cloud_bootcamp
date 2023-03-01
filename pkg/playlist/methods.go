package playlist

import (
	"container/list"
	"errors"
	"fmt"
	"github.com/ivahaev/timer"
	"log"
)

func (p *Player) Play() error {
	p.Mu.Lock()
	if p.FirstTrack.Len() == 0 {
		p.Mu.Unlock()
		return errors.New("playlist is empty")
	}
	if p.NowPlaying == nil {
		p.NowPlaying = p.FirstTrack.Front()
		go p.PlaySong()
	}
	p.Mu.Unlock()
	p.PlayChan <- struct{}{}
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
	log.Println("PAUSECONGFUNC::::", p.NowPlaying.Value.(Song).IsPlaying)
	if p.NowPlaying.Value.(Song).IsPlaying == false {
		p.Mu.Unlock()
		return errors.New("no song is playing now")
	}
	song := p.NowPlaying.Value.(Song)
	song.IsPlaying = false
	p.NowPlaying.Value = song
	p.Mu.Unlock()
	p.PauseChan <- struct{}{}

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

	p.Mu.Lock()
	for e := p.FirstTrack.Front(); e != nil; e = e.Next() {
		song := e.Value.(Song)
		if song.Name == songTarget.Name {
			res = e
			break
		}
	}
	if res == nil {
		p.Mu.Unlock()
		return errors.New("no such song")
	}
	cast, ok := res.Value.(Song)
	if !ok || cast.Name != songTarget.Name {
		p.Mu.Unlock()
		return errors.New("no such song")
	}
	if res.Value.(Song).IsPlaying {
		p.Mu.Unlock()
		return errors.New("can't delete, the song is playing now")
	}
	fmt.Println("DEELLLETEE:", res)
	p.FirstTrack.Remove(res)
	p.Mu.Unlock()

	/// проверить мьютексы!!!!!! тут с ними очень все плохо. и может вообще
	/// выкинуть мьютексы и оставить их в хендлерах
	return nil
}

func (p *Player) PlaySong() {
	//затестить мьютексы
	for e := p.FirstTrack.Front(); e != nil; e = e.Next() {
		fmt.Println("new iteration")
		p.Mu.Lock()
		if p.FlagPrev {
			p.FlagPrev = false
			e = e.Prev().Prev()
		}
		p.NowPlaying = e
		song := p.NowPlaying.Value.(Song)
		p.Mu.Unlock()
		timerDuration := timer.NewTimer(song.Duration)
		timerDuration.Start()
	LOOP:
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
				fmt.Println("playyer", p.NowPlaying.Value)
				song.IsPlaying = false
				p.Mu.Lock()
				p.NowPlaying.Value = song
				fmt.Println("playyer2222222222", p.NowPlaying.Value)
				p.Mu.Unlock()
			case <-timerDuration.C:
				fmt.Println("дошли до конца таймера")
				break LOOP
			case <-p.NextChan:
				break LOOP
			case <-p.PrevChan:
				p.Mu.Lock()
				p.FlagPrev = true
				p.Mu.Unlock()
				break LOOP
			}
			fmt.Println("вышли из селекта, сейчас в бесконечности")
		}
		fmt.Println("вышли из бесконечности, сейчас в основном цикле")
	}
	fmt.Println("выходим из горутины почему-то")
	p.NowPlaying = nil
}

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
