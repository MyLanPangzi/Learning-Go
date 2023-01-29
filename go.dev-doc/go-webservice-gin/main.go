package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
	"web-service-gin/entity"
	"web-service-gin/service"
	"web-service-gin/service/impl"
)

func makeGinHandler(fn func(c *gin.Context, s service.AlbumService), s service.AlbumService) gin.HandlerFunc {
	return func(context *gin.Context) {
		fn(context, s)
	}
}
func main() {

	db := openDb()
	defer db.Close()

	var albumService service.AlbumService = impl.NewAlbumService(db)

	router := gin.Default()
	router.GET("/albums", makeGinHandler(getAlbums, albumService))
	router.GET("/albums/:id", makeGinHandler(getAlbumById, albumService))
	router.POST("/albums", makeGinHandler(postAlbums, albumService))

	log.Fatal(router.Run("localhost:8080"))
}

func openDb() *sql.DB {
	var cfg = mysql.Config{
		User:   "root",
		Passwd: "!QAZ2wsx",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "recordings",
	}
	var db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func getAlbums(c *gin.Context, albumService service.AlbumService) {
	c.IndentedJSON(http.StatusOK, albumService.GetAll())
}

func postAlbums(c *gin.Context, albumService service.AlbumService) {
	var newAlbum entity.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	add, err := albumService.Add(newAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "interval server error"})
		return
	}
	c.IndentedJSON(http.StatusCreated, add)
}

func getAlbumById(c *gin.Context, albumService service.AlbumService) {
	id, err := strconv.ParseInt(c.Param("id"), 0, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id format err"})
		return
	}
	album, err := albumService.Get(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}
