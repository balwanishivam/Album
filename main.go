package main

import (
	"handler/albumHandler"
	"github.com/gin-gonic/gin"
)

// type Album struct {
// 	ID     string  `json:"id"`
// 	Title  string  `json:"title"`
// 	Artist string  `json:"artist"`
// 	Price  float64 `json:"price"`
// }

// // albums slice to seed record album data.
// var albums = []Album{
// 	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

// // getAlbums responds with the list of all albums as JSON.
// func getAlbums(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, gin.H{"success": "true", "data": albums})
// }

// // postAlbums adds an album from JSON received in the request body.
// func postAlbums(c *gin.Context) {
// 	var newAlbum Album

// 	// Call BindJSON to bind the received JSON to
// 	// newAlbum.
// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		return
// 	}

// 	// Add the new album to the slice.
// 	albums = append(albums, newAlbum)
// 	c.IndentedJSON(http.StatusOK, gin.H{"success": "true", "data": newAlbum})
// }

// func deleteAlbum(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	for i, album := range albums {
// 		if album.ID == id {
// 			albums = append(albums[:i], albums[i+1:]...)
// 			c.JSON(http.StatusOK, gin.H{"success": "true", "error": "album deleted"})
// 			return
// 		}
// 	}
// 	c.JSON(http.StatusBadRequest, gin.H{"id #" + id: "not found"})
// }
// func getAlbumByID(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	for _, album := range albums {
// 		if album.ID == id {
// 			c.IndentedJSON(http.StatusOK, gin.H{"success": "true", "data": album})
// 			return
// 		}
// 	}
// 	c.JSON(http.StatusBadRequest, gin.H{"success": "false", "error": "album not found"})
// }

// func updateAlbum(c *gin.Context) {
// 	var newAlbum Album

// 	// Call BindJSON to bind the received JSON to
// 	// newAlbum.
// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		return
// 	}
// 	for i, album := range albums {
// 		if album.ID == newAlbum.ID {
// 			albums[i].ID = newAlbum.ID
// 			albums[i].Title = newAlbum.Title
// 			albums[i].Artist = newAlbum.Artist
// 			albums[i].Price = newAlbum.Price
// 			c.IndentedJSON(http.StatusOK, gin.H{"success": "true", "data": albums[i]})
// 			return
// 		}
// 	}
// 	c.JSON(http.StatusBadRequest, gin.H{"success": "false", "error": "album to be updated not found"})
// }
func main() {
	router := gin.Default()
	router.GET("/albums/", getAlbums)
	router.POST("/albums/", postAlbums)
	router.DELETE("/albums/:id/", deleteAlbum)
	router.GET("/albums/:id/", getAlbumByID)
	router.PUT("/albums/", updateAlbum)
	router.Run("localhost:8080")
}
