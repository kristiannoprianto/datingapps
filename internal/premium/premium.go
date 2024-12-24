package premium

import (
	"apps/datingapps/pkg/database"
	"apps/datingapps/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UnlockPremium unlocks premium features for a user
func UnlockPremium(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	// Update the user's premium status
	if err := database.DB.Model(&models.Profile{}).Where("id = ?", userID).Update("premium", true).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unlock premium"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Premium features unlocked"})
}
