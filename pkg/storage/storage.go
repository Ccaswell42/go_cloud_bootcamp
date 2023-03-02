package storage

import (
	"container/list"
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"gocloud/config"
	"log"
	"strconv"
	"time"
)

type Repository struct {
	Db *sql.DB
}

type Song struct {
	Name     string
	Duration time.Duration
}

func ConnectToDb(conf *config.Config) (*Repository, error) {

	db, err := sql.Open(conf.DriverName, conf.DSN)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)

	err = db.Ping()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &Repository{db}, nil
}

func (r *Repository) GetAll() (*list.List, error) {
	rows, err := r.Db.Query("SELECT name, duration from playlist")
	if err != nil {
		log.Println("query err:", err)
		return nil, err
	}
	defer rows.Close()

	songList := list.New()
	for rows.Next() {
		song := Song{}
		var durationNum int
		err = rows.Scan(&song.Name, &durationNum)
		str := strconv.Itoa(durationNum) + "s"
		durationTime, err := time.ParseDuration(str)
		if err != nil {
			log.Println("parse duration err:", err)
			err = errors.New("problem downloading song from database")
			continue
		}
		song.Duration = durationTime
		songList.PushBack(song)
	}
	return songList, err
}

func (r *Repository) DeleteSong(song Song) error {

	_, err := r.Db.Exec("DELETE from playlist where name = $1", song.Name)
	if err != nil {
		log.Println("delete error:", err)
		return err
	}
	return nil
}

func (r *Repository) Set(song Song) error {
	durationNum := int(song.Duration.Seconds())
	_, err := r.Db.Exec("insert into playlist values (default, $1, $2)",
		song.Name, durationNum)
	if err != nil {
		log.Println("insert error: ", err)
		return err
	}
	return nil
}
