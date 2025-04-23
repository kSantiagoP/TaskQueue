package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	//Initialize Database
	var err error

	db, err = InitializeSQlite()
	if err != nil {
		return fmt.Errorf("error initializing databases: %v", err)
	}

	//Initialize Environment variables
	err = InitializeEnvironment()
	if err != nil {
		return fmt.Errorf("error initializing environment variables: %v", err)
	}

	return nil
}

func GetSQlite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	//Initialize logger
	logger = NewLogger(p)

	return logger
}
