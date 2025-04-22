package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kSantiagoP/TaskQueue/api/handler"
)

func initializeRoutes(router *gin.Engine) {
	//initialize handler
	handler.InitializeHandler()

	basePath := "/api/v1"

	v1 := router.Group(basePath)
	{
		v1.GET("/ping", handler.PlaceHolderHandler)
		v1.GET("/task", handler.PlaceHolderHandler)
		v1.POST("/task", handler.PlaceHolderHandler)
		v1.GET("/tasks", handler.PlaceHolderHandler)
		v1.PUT("/task/cancel", handler.PlaceHolderHandler)
	}
}
