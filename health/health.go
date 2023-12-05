package health

import (
	"github.com/gin-gonic/gin"
)

func getHealth(c *gin.Context) {
	c.String(200, "ok")
}

func HealthSetup(r *gin.Engine) {
	r.GET("/health", getHealth)
}