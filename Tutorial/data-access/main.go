package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {
	//Capture connetcion properties
	cfg := mysql.NewConfig()
	cfg.User = "root"
	cfg.Passwd = "MysqlD4d0!"
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "recordings"

	//Get a databse handle
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	//Hard-code ID 2 here to test the query
	alb, err := albumById(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", alb)

	albId, err := addAlbum(Album{
		Title:  "The modern sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albId)
}

// album by artist query for albums that gave the specified artist name.
func albumsByArtist(name string) ([]Album, error) {
	// An album slice to hold the data from returned rows
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign colum data to struct field
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumByArtist %q: %v", name, err)
	}
	return albums, nil
}

// albumById queries for the album with the specified ID
func albumById(id int64) (Album, error) {
	//An album to hold data for the returned row
	var alb Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		return alb, fmt.Errorf("AlbumsById %d:%v", id, err)
	}
	return alb, nil
}

// addAlbum adds the specified album to the databse,
// returning the album ID of the new entry
func addAlbum(alb Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title,artist,price) VALUES (?,?,?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}
