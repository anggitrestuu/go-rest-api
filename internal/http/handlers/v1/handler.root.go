package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Root(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"staus":   true,
		"message": "v1 online...",
	})
}
