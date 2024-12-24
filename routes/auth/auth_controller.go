package auth

import (
	"apps/datingapps/internal/auth"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	// Call internal auth logic for sign-up
	auth.SignUp(c)
}

func Login(c *gin.Context) {
	// Call internal auth logic for login
	auth.Login(c)
}
