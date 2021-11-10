package handler

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"manga-api/httpd/manga"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
var err error

const (
	getSpecifiedManga = "SELECT * FROM release_date WHERE title = ?;"
)

func init() {
	Db, err = sql.Open("mysql", "tabe:"+os.Getenv("PASS")+"@/manga")
	if err != nil {
		log.Fatal("DB connection error", err)
	}
}

func GetAllMangas() []manga.Manga {
	rows, err := Db.Query("SELECT * FROM release_date")
	if err != nil {
		log.Fatal("GetAllMangas error", err)
	}

	var mangas []manga.Manga
	manga := manga.Manga{}
	for rows.Next() {
		err := rows.Scan(&manga.ID, &manga.Title, &manga.ReleaseDate)
		if err != nil {
			log.Fatal("Scan error", err)
		} else {
			mangas = append(mangas, manga)
		}
	}

	return mangas
}

func GetSpecifiedManga(ctx *gin.Context) manga.Manga {
	title := ctx.Param("title")
	fmt.Println("title", title)
	stmt, err := Db.Prepare(getSpecifiedManga)
	if err != nil {
		log.Fatal("stmt error", err)
	}
	defer stmt.Close()

	row, err := stmt.Query(title)
	if err != nil {
		log.Fatal("query error", err)
	}

	mng := manga.Manga{}
	for row.Next() {
		err = row.Scan(&mng.ID, &mng.Title, &mng.ReleaseDate)
	}
	if err != nil {
		log.Fatal("Row Scan error", err)
	}

	return mng
}
