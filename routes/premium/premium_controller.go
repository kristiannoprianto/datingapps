package premium

import (
	"apps/datingapps/internal/premium"

	"github.com/gin-gonic/gin"
)

// UnlockPremium handles the unlock premium route
func UnlockPremium(c *gin.Context) {
	// Call the internal logic to unlock premium features for the user
	premium.UnlockPremium(c)
}
