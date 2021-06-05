package handler

import (
	"net/http"

	"example.com/gohttp/platform/newsfeed"
	"github.com/gin-gonic/gin"
)

func NewsFeedGet(feed newsfeed.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := feed.GetAll()
		c.JSON(http.StatusOK,  results)
	}
}