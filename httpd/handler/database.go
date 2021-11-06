package handler

import (
	"database/sql"
	"log"
	"os"

	"manga-api/httpd/manga"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
var err error

func init() {
	Db, err = sql.Open("mysql", "tabe:"+os.Getenv("PASS")+"@/manga")
	if err != nil {
		log.Fatal("DB connection error")
	}
}

func GetAllMangas() []manga.Manga {
	rows, err := Db.Query("SELECT * FROM release_date")
	if err != nil {
		log.Fatal("GetAllMangas error")
	}

	var mangas []manga.Manga
	manga := manga.Manga{}
	for rows.Next() {
		err := rows.Scan(&manga.ID, &manga.Title, &manga.ReleaseDate)
		if err != nil {
			log.Fatal("Scan error")
		} else {
			mangas = append(mangas, manga)
		}
	}

	return mangas
}
