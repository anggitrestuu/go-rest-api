package middlewares

import (
	"net/http"

	"github.com/anggitrestuu/go-rest-api/internal/constants"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", constants.AllowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", constants.AllowCredential)
		c.Writer.Header().Set("Access-Control-Allow-Headers", constants.AllowHeader)
		c.Writer.Header().Set("Access-Control-Allow-Methods", constants.AllowMethods)
		c.Writer.Header().Set("Access-Control-Max-Age", constants.MaxAge)

		println("CORS Middleware")

		// if !helpers.IsArrayContains(strings.Split(constants.AllowMethods, ", "), c.Request.Method) {
		// 	logger.InfoF("method %s is not allowed\n", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryCORS}, c.Request.Method)
		// 	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden with CORS policy 1"})
		// 	return
		// }

		// for key, value := range c.Request.Header {
		// 	fmt.Println(key, value)
		// 	if !helpers.IsArrayContains(strings.Split(constants.AllowHeader, ", "), key) {
		// 		logger.InfoF("in header %s: %s\n", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryCORS}, key, value)
		// 		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden with CORS policy 2"})
		// 		return
		// 	}
		// }

		// if constants.AllowOrigin != "*" {
		// 	if !helpers.IsArrayContains(strings.Split(constants.AllowOrigin, ", "), c.Request.Host) {
		// 		logger.InfoF("host '%s' is not part of '%v'\n", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryCORS}, c.Request.Host, constants.AllowOrigin)
		// 		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden with CORS policy 3"})
		// 		return
		// 	}
		// }

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}
