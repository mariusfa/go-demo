package health

import (
	"github.com/gin-gonic/gin"
)

func getHealth(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.String(500, "user not found")
		return
	}
	c.String(200, "ok %s", user)
}

func HealthSetup(r *gin.Engine) {
	r.GET("/health", getHealth)
}