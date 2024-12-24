package swipes

import (
	"apps/datingapps/pkg/database"
	"apps/datingapps/pkg/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Swipe handles the user swipe actions (like or pass)
func Swipe(c *gin.Context) {
	var input struct {
		ProfileID uint   `json:"profile_id"`
		Action    string `json:"action"` // "like" or "pass"
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userID := c.MustGet("user_id").(uint) // Extract user ID from JWT
	if input.Action != "like" && input.Action != "pass" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid action"})
		return
	}

	// Check if the user has already swiped today
	var swipes []models.Swipe
	if err := database.DB.Where("user_id = ? AND created_at > ?", userID, time.Now().Truncate(24*time.Hour)).Find(&swipes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking swipes"})
		return
	}

	if len(swipes) >= 10 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Swipe limit reached for today"})
		return
	}

	// Create a new swipe record
	swipe := models.Swipe{UserID: userID, ProfileID: input.ProfileID, Action: input.Action}
	if err := database.DB.Create(&swipe).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Swipe action failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Swipe recorded"})
}
