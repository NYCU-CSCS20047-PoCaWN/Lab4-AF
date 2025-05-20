package processor

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/NYCU-CSCS20047-PoCaWN/lab4-af/internal/logger"
	"github.com/NYCU-CSCS20047-PoCaWN/lab4-af/internal/models"
)

func (p *Processor) GetWarningUsers(c *gin.Context) {
	userUsage, err := p.Consumer().GetUserUsage(context.Background())
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to get user usage:" + err.Error(),
		})
		return
	}

	var warningUsers []models.WarningUser

	// TODO: Implement logic to determine warning users based on userUsage
	// Use userUsage to get warning users
	// For now, just return the userUsage as a placeholder
	for _, usage := range userUsage {
		// Check if the user is a warning user
		logger.ProcessorLog.Errorf("Debug: %v", usage)
	}

	c.JSON(http.StatusOK, warningUsers)
}
