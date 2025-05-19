package processor

import (
	"context"

	"github.com/gin-gonic/gin"
)

func (p *Processor) GetWarningUsers(c *gin.Context) {
	userUsage, err := p.Consumer().GetUserUsage(context.Background())
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to get user usage:" + err.Error(),
		})
		return
	}

	// TODO: Implement logic to determine warning users based on userUsage
	// Use userUsage to get warning users
	// For now, just return the userUsage as a placeholder
	c.JSON(200, gin.H{
		"warning_users": userUsage,
	})
}
