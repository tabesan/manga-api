package main

import (
	"manga-api/httpd/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	defer handler.Db.Close()
	router := gin.Default()
	router.GET("/", handler.RefAllMangas())

	router.Run()
}
