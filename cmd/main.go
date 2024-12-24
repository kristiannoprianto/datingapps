package main

import (
	"apps/datingapps/pkg/database"
	"apps/datingapps/routes/auth"
	"apps/datingapps/routes/premium"
	"apps/datingapps/routes/swipes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	if err := database.Connect(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	r := gin.Default()

	// Set up routes
	r.POST("/signup", auth.SignUp)
	r.POST("/login", auth.Login)
	r.POST("/swipe", swipes.Swipe)
	r.POST("/unlock-premium", premium.UnlockPremium)

	// Start the server
	r.Run(":8080")
}
