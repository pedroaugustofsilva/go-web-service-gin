package main

import (
	"log"
	"net/http"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var db, dbErr = sql.Open("postgres", "postgres://postgres:postgres@postgres:5432?sslmode=disable")

func getAlbums(c *gin.Context) {
	var albums = []*album{}
	albums_db, err := db.Query("SELECT * FROM albums")

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	for albums_db.Next() {
		a := new(album)
		err := albums_db.Scan(&a.ID, &a.Title, &a.Artist, &a.Price)

		if err != nil {
			log.Fatal(err)
		}

		albums = append(albums, a)
	}

	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	stmt, err := db.Prepare("INSERT INTO albums (title, artist, price) VALUES ($1, $2, $3)")

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(newAlbum.Title, newAlbum.Artist, newAlbum.Price)

	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"error": err})
	}
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	var album = new(album)

	err := db.QueryRow("SELECT * FROM albums where id = $1", id).Scan(&album.ID, &album.Title, &album.Artist, &album.Price)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.IndentedJSON(http.StatusOK, album)
}

func main() {
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("0.0.0.0:3000")
}
