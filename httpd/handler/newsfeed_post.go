package handler

import (
	"net/http"

	"example.com/gohttp/platform/newsfeed"
	"github.com/gin-gonic/gin"
)

type newsFeedPostRequest struct  {
	Title string  `json:"title"`
	Post string `json:"post"`
}

func NewsFeedPost(feed newsfeed.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := newsFeedPostRequest{}
		c.Bind(&requestBody)

		item := newsfeed.Item{
			Title: requestBody.Title,
			Post: requestBody.Post,
		}

		feed.Add(item)
		c.Status(http.StatusNoContent)
	}
}