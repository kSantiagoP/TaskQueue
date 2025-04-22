package main

import (
	"github.com/kSantiagoP/TaskQueue/api/router"
	"github.com/kSantiagoP/TaskQueue/config"
)

var (
	logger *config.Logger
)

func main() {

	//initialize configs
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Initialize()
}
