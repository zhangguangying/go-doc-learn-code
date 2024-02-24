package db

import "database/sql"
import "github.com/go-sql-driver/mysql"
import "log"
import "fmt"

var db *sql.DB

type Album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

func init() {
	cfg := mysql.Config{
		User: "homestead",
		Passwd: "secret",
		Net: "tcp",
		Addr: "192.168.10.10:3306",
		DBName: "recordings",
		AllowNativePasswords: true,
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
}

func GetAlbums() ([]Album, error) {
	var albums []Album

	rows, err := db.Query("SELECT * FROM album")
	if err != nil {
		return nil, fmt.Errorf("getAlbums: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("getAlbums: %v", err)
		}
		albums = append(albums, alb)
	}

	return albums, nil
}