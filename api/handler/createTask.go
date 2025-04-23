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
		Status:      request.Status,
		Error:       request.Error,
		ElapsedTime: request.ElapsedTime,
	}

	if err := db.Create(&task).Error; err != nil {
		logger.Errorf("error creating task: %v", err)
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(c, "createTask", task)
}

type CreateTaskRequest struct {
	Type        string `json:"type"`
	Status      string `json:"status"`
	Error       string `json:"error"`
	ElapsedTime int64  `json:"elapsedTime"`
}
