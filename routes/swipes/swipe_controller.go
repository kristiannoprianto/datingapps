package swipes

import (
	"apps/datingapps/internal/swipes"

	"github.com/gin-gonic/gin"
)

// Swipe handles the swipe route (like or pass)
func Swipe(c *gin.Context) {
	// Call the internal logic for handling swipes
	swipes.Swipe(c)
}
