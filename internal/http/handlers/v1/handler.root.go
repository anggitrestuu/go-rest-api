package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Show an account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {string} string	"ok"
// @Router /accounts/{id} [get]
func Root(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"staus":   true,
		"message": "v1 online...",
	})
}
