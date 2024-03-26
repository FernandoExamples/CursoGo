package main

import (
	"net/http"
	"rest-api-gin/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Albums)
}

func addAlbum(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	models.Albums = append(models.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	idNum, err := strconv.Atoi(id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	for _, album := range models.Albums {
		if album.ID == idNum {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", addAlbum)
	router.GET("/albums/:id", getAlbumByID)

	router.Run("localhost:3000")
}
