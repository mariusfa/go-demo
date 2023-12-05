package middleware

import "github.com/gin-gonic/gin"

func Marius(c *gin.Context) {
	user := "Marius"
	c.Set("user", user)

	// Process request
	c.Next()
}