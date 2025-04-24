package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kSantiagoP/TaskQueue/schemas"
)

func CreateTaskHandler(c *gin.Context) {
	request := CreateTaskRequest{}

	c.BindJSON(&request)

	task := schemas.Task{
		Type:        request.Type,
		Status:      "pending",
		Error:       "",
		ElapsedTime: request.ElapsedTime,
	}

	if err := db.Create(&task).Error; err != nil {
		logger.Errorf("error creating task: %v", err)
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(c, "createTask", task)
}
