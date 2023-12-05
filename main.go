package main

import (
	"log"
	"os"
	"todo/health"
	"todo/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)

	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middleware.AccessLogger)
	r.Use(middleware.Marius)

	health.HealthSetup(r)

	r.Run()
}