package main

import (
	"example.com/gohttp/httpd/handler"
	"example.com/gohttp/platform/newsfeed"
	"github.com/gin-gonic/gin"
)


func main() {
	feed := newsfeed.New()
	r := gin.Default()

	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsFeedGet(feed))
	r.POST("/newsfeed", handler.NewsFeedPost(feed))

	r.Run(":8080")
	

}