package main

import (
	"context"
	"hash/fnv"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type link struct {
	ID    string `json:"id"`
	url   string `json:"url"`
	views int    `json:"views"`
}

// {ID: "a2", url: "https://www3.nd.edu/~pbui/teaching/cse.40842.fa24/reading01.html", views: 0},

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     "pspbalsaas-db-1:6379",
	Password: "",
	DB:       0,
})

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("./src/templates/*")
	router.GET("/:id", getLink)
	router.GET("/", home)
	router.POST("/links", createLink)
	router.POST("/pastebin", home)

	router.Run("0.0.0.0:80")
}

func url_hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func createLink(c *gin.Context) {

	//var new_link link

	//if err := c.BindJSON(&new_link); err != nil {
	//	return
	//}

	// links = append(links, newLink)
	// c.IndentedJSON(http.StatusCreated, newLink)

	url := c.Query("url")
	hashed_url := strconv.Itoa(int(url_hash(url)))

	err := rdb.Set(ctx, hashed_url, url, 0).Err()
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusCreated, "http://s.jdolak.com/"+hashed_url)
}

func getLink(c *gin.Context) {

	id := c.Param("id")

	val, err := rdb.Get(ctx, id).Result()
	if err == redis.Nil {
		c.JSON(http.StatusNotFound, "")
	} else if err != nil {
		panic(err)
	} else {
		c.Redirect(http.StatusMovedPermanently, val)
	}

	//for _, a := range links {
	//    if a.ID == id {
	//		c.Redirect(http.StatusMovedPermanently, a.url)
	//        return
	//    }
	//}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Main website"})
}
