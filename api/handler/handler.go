package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kSantiagoP/TaskQueue/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQlite()
}

func PlaceHolderHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
