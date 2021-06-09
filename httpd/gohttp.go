package main

import (
	"log"

	"example.com/gin-http/httpd/handler"
	"example.com/gin-http/platform/newsfeed"
	"example.com/gin-http/platform/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func init(){
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found. Using default values.")
	}
}

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

	r.GET("/database/user", handler.GetUserFromDB("test"))

	r.Run(":8080")
	

}