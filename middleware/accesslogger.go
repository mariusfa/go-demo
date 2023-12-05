package middleware

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func AccessLogger(c *gin.Context) {
	log.SetPrefix("[AccessLogger] ")
	path := c.Request.URL.Path
	logStart := fmt.Sprintf("Start of request %s", path)
	log.Println(logStart)

	// Process request
	c.Next()

	logEnd := fmt.Sprintf("End of request %s", path)
	log.Println(logEnd)
}