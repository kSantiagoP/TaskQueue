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

	//Definir rotas
	initializeRoutes(router)

	router.Run(port) // listen and serve on localhost:4000
}
