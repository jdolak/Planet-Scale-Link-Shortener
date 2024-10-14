package main

import (
	"context"
	"hash/fnv"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

//type link struct {
//	ID    string `json:"id"`
//	url   string `json:"url"`
//	views int    `json:"views"`
//}

type json_data struct {
	data string `json:"views"`
}

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

	url := c.Query("url")
	hashed_url := strconv.Itoa(int(url_hash(url)))

	err := rdb.Set(ctx, hashed_url, url, 0).Err()
	if err != nil {
		panic(err)
	}

	returnURL := "http://" + c.Request.Host + "/" + hashed_url

	c.IndentedJSON(http.StatusCreated, returnURL)
}

func createPaste(c *gin.Context) {

	var new_data json_data

	if err := c.BindJSON(&new_data); err != nil {
		return
	}

	c.IndentedJSON(http.StatusCreated, new_data)

	hashed_url := strconv.Itoa(int(url_hash(new_data.data)))

	err := rdb.Set(ctx, hashed_url, new_data.data, 0).Err()
	if err != nil {
		panic(err)
	}

	returnURL := "http://" + c.Request.Host + "/" + hashed_url

	c.IndentedJSON(http.StatusCreated, returnURL)
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

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "link not found"})
}

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Main website"})
}
