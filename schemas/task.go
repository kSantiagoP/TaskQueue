package schemas

import (
	"time"

	"gorm.io/gorm"
)

//task schema configuration

type Task struct {
	gorm.Model
	Type        string
	Status      string
	Error       string
	ElapsedTime int64
}

type TaskResponse struct {
	ID          uint      `json:"id"`
	Type        string    `json:"type"`
	Status      string    `json:"status"`
	Error       string    `json:"error"`
	ElapsedTime int64     `json:"elapsedTime"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt"`
}
