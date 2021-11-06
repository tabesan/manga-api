package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RefAllMangas() gin.HandlerFunc {
	return func(c *gin.Context) {
		result := GetAllMangas()
		c.JSON(http.StatusOK, result)
	}
}
