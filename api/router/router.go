package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kSantiagoP/TaskQueue/config"
)

func Initialize() {
	logger := config.GetLogger("router")

	port, err := config.GetPort()
	if err != nil {
		logger.Errorf("error retrieving port %v", err)
		return
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
