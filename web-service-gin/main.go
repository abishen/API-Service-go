package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"reflect"
)

func main () {
	router := gin.Default()
	router.GET("/albums",getAlbums)
	//router.GET("/albums/:id",getAlbumByID)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
}

// func getAlbumByID(c * gin.Context) {
// 	id := c.Param("id")
	
// 	foundAlbum := getAlbum(id)
// 	if  foundAlbum != false {
// 		c.IndentedJSON(http.StatusOK,foundAlbum)
// 	}

// 	for _,a := range albums {
// 		if a.ID == id {
// 			return a
// 		}
// 	}
	
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "album not found"})
// }

// func getAlbum(id string) album {

// 	item := album {}
// 	return item
// }


func getAlbums(c * gin.Context) {
	c.IndentedJSON(http.StatusOK,albums)
}

func postAlbums(c * gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}


type album struct {
	ID string 
	Title string
	Artist string
	Price float64
}
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}