package main

import (
	"example.com/gin-http/httpd/handler"
	"example.com/gin-http/platform/newsfeed"
	"example.com/gin-http/platform/user"
	"github.com/gin-gonic/gin"
)


func main() {
	feed := newsfeed.New()
	user := user.New()
	r := gin.Default()

	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsFeedGet(feed))
	r.POST("/newsfeed", handler.NewsFeedPost(feed))
	r.GET("/user", handler.GetAllUsers(user))
	r.GET("/user/:username", handler.GetUser(user))
	r.POST("/user/new", handler.CreateAccount(user))

	r.Run(":8080")
	

}