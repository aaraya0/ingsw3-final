package app

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartRoute() {
	mapUrls()

	port := os.Getenv("PORT")
	if port == "" {
		port = "80" // Usa un puerto predeterminado si no se especifica PORT
	}

	address := "0.0.0.0:80"
	log.Printf("Starting server on %s\n", address)
	router.Run(address)
}
