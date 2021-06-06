package handler

import (
	"net/http"
	"time"

	"example.com/gin-http/platform/user"
	"github.com/gin-gonic/gin"
)

type userPostRequest struct  {
	Username string  `json:"username"`
	Password string `json:"password"`
}

func CreateAccount(u user.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := userPostRequest{}
		c.Bind(&requestBody)

		newUser := user.User{
			Username: requestBody.Username,
			Password: requestBody.Password,
			CreatedOn: time.Now().Format(time.RFC3339),
		}

		u.Add(newUser)
		c.Status(http.StatusNoContent)
	}
}