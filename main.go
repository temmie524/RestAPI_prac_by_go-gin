package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// レコードのアルバムのデータの素となるスライス`albums`
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbumsはJSON形式の全てのアルバムのリストを返えします。
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.PUT("albums/:id", updateAlbumByID)
	router.DELETE("albums/:id", deleteAlbumByID)
	router.Run("localhost:8080")
}

// postAlbumsはリクエストボディのJSONからアルバムを追加します
func postAlbums(c *gin.Context) {
	var newAlbum album

	// 受け取ったJSONを`newAlbum`にバインドするために`BindJSON`を呼び出す
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// スライスへ新しいアルバムを追加する
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// `getAlbumByID`は`id`にマッチするIDを持つアルバムの場所を取得します。
// クライアントからパラメタが送られたら、レスポンスとしてアルバムを返します。
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// IDの値とマッチするパラメタをもつアルバムを探すために
	// リストのアルバムをループします.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func updateAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var album album

	if err := c.BindJSON(&album); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, a := range albums {
		if a.ID == id {
			albums[i] = album
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")
	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Album has been deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
