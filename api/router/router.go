package router

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Initialize() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	port := ":"
	port += os.Getenv("SERVER_PORT")

	if port == "" {
		log.Fatal("Error retrieving .env")
	}

	router := gin.Default()

	//Definir rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(port) // listen and serve on localhost:4000
}
