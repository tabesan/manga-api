package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RefAllMangas() gin.HandlerFunc {
	return func(c *gin.Context) {
		result := GetAllMangas()
		c.JSON(http.StatusOK, result)
	}
}

func RefSpecifiedManga() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Param("title"))
		result := GetSpecifiedManga(c)
		c.JSON(http.StatusOK, result)
	}
}
