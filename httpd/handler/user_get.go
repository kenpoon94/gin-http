package handler

import (
	"net/http"

	"example.com/gin-http/platform/user"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(u user.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := u.GetAll()
		c.JSON(http.StatusOK,  results)
	}
}

func GetUser(u user.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")
		results := u.Get(username)
		if results == nil {
			c.JSON(http.StatusNoContent, gin.H{
				"message":"User not found",
			})
		} else {
			c.JSON(http.StatusOK,  results)
		}
	}
}