package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"Artist"`
	Price  float64 `json:"price"`
}

// album slice to seed record album data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Green Train", Artist: "Tom Coltrane", Price: 66.99},
	{ID: "3", Title: "Red Train", Artist: "Harry Coltrane", Price: 76.99},
}

// getAlbums responds with list of all albums as JSON
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)

	router.Run("localhost:8080")
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind received JSON to new album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// locates the album by matching ID
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over list of albums to match
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
