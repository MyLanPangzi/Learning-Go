package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

type Album struct {
	Id     int64
	Title  string
	Artist string
	Price  float64
}

func main() {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "!QAZ2wsx",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "recordings",
	}
	dsn := cfg.FormatDSN()
	fmt.Println(dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected")
	defer db.Close()

	artist, err := albumByArtist(db, "John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range artist {
		fmt.Printf("%+v\n", v)
	}

	album, err := albumById(db, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(album)
	id, err := addAlbum(db, Album{
		Title:  "hello",
		Artist: "world",
		Price:  1,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("add album id", id)
}

func albumByArtist(db *sql.DB, name string) (albums []Album, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("albumByArtist %q: %v", name, err)
			albums = nil
		}
	}()
	rows, err := db.Query("select * from album where artist = ?", name)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var e Album
		if err = rows.Scan(&e.Id, &e.Title, &e.Artist, &e.Price); err != nil {
			return
		}
		albums = append(albums, e)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return albums, nil
}

// albumByID queries for the album with the specified ID.
func albumById(db *sql.DB, id int64) (Album, error) {
	var r Album
	row := db.QueryRow("select * from album where id = ?", id)
	if err := row.Scan(&r.Id, &r.Title, &r.Artist, &r.Price); err != nil {
		if err == sql.ErrNoRows {
			return r, fmt.Errorf("albumById %d: no such album", id)
		}
		return r, fmt.Errorf("albumById %d: %v", id, err)
	}
	return r, nil
}

// // addAlbum adds the specified album to the database,
// // returning the album ID of the new entry
func addAlbum(db *sql.DB, a Album) (int64, error) {
	r, err := db.Exec("insert into album (title,artist,price) values (?,?,?)", a.Title, a.Artist, a.Price)
	if err != nil {
		return 0, fmt.Errorf("add album: %v", err)
	}
	id, err := r.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("add album: %v", err)
	}
	return id, nil
}
