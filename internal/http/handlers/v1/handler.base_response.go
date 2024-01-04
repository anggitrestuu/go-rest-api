package v1

import (
	"net/http"

	"github.com/anggitrestuu/go-rest-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type BaseResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewSuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {

	c.JSON(statusCode, BaseResponse{
		Status:  true,
		Message: message,
		Data:    data,
	})
}

func NewErrorResponse(c *gin.Context, statusCode int, err string) {

	logger.ErrorF("Error Response", logrus.Fields{"status": statusCode, "message": err})

	c.JSON(statusCode, BaseResponse{
		Status:  false,
		Message: err,
	})

}

func NewAbortResponse(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": message})
}
