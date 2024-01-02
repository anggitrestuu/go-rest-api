package middlewares

import (
	"net/http"

	casbin "github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func Authorize(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user role from the request, e.g., JWT token
		// userRole := getUserRole(c)

		userRole := "admin"

		// Check permission
		allowed, err := e.Enforce(userRole, c.Request.URL.Path, c.Request.Method)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			c.Abort()
			return
		}

		// If not allowed, return error
		if !allowed {
			c.JSON(http.StatusForbidden, "Forbidden")
			c.Abort()
			return
		}

		c.Next()
	}
}
