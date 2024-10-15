package main

import (
	"context"
	"encoding/json"
	"hash/fnv"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// ENUM for link types
const (
	Redirect = iota
	Paste
)

// Struct for storing in the DB
type link struct {
	Link_type int    `json:"type"`
	Data      string `json:"data"`
	Views     int    `json:"views"`
}

// Struct for binding with paste POST request
type json_data struct {
	Data string `json:"data"`
}

// connection with redis db
var rdb = redis.NewClient(&redis.Options{
	Addr:     "pspbalsaas-db-1:6379",
	Password: "",
	DB:       0,
})

var ctx = context.Background()

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("./src/templates/*")

	router.GET("/:id", getLink)
	router.GET("/", home)

	router.POST("/links", createLink)
	router.POST("/paste", createPaste)

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

	var new_link link

	new_link.Link_type = Redirect
	new_link.Data = url
	new_link.Views = 0

	value, err := json.Marshal(new_link)
	if err != nil {
		log.Fatal(err)
	}

	err = rdb.Set(ctx, hashed_url, string(value), 0).Err()
	if err != nil {
		panic(err)
	}

	returnURL := "http://" + c.Request.Host + "/" + hashed_url

	c.IndentedJSON(http.StatusCreated, returnURL)
}

func createPaste(c *gin.Context) {

	var new_data json_data

	err := c.BindJSON(&new_data)
	if err != nil {
		log.Fatal(err)
	}

	hashed_url := strconv.Itoa(int(url_hash(new_data.Data)))

	var new_link link

	new_link.Link_type = Paste
	new_link.Data = new_data.Data
	new_link.Views = 0

	value, err := json.Marshal(new_link)
	if err != nil {
		log.Fatal(err)
	}

	err = rdb.Set(ctx, hashed_url, string(value), 0).Err()
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
		log.Fatal(err)
	}
	var loaded_link link

	err = json.Unmarshal([]byte(val), &loaded_link)
	if err != nil {
		log.Fatal(err)
	}

	switch loaded_link.Link_type {
	case Redirect:
		c.Redirect(http.StatusMovedPermanently, loaded_link.Data)
		return
	case Paste:
		c.String(http.StatusOK, loaded_link.Data)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "link not found"})
}

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Main website"})
}
