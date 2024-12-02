package middlewares

import (
	"basicAuth/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PermissionMiddleware(requiredPermission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("currentUser")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		currentUser := user.(models.User)
		if currentUser.IsAdmin == false {
			hasPermission := false
			for _, perm := range currentUser.Permission {
				if perm.Name == requiredPermission {
					hasPermission = true
					break
				}
			}

			if !hasPermission {
				c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
				c.Abort()
				return
			}

		}

		c.Next()
	}
}
