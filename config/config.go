package config

import "fmt"

var (
	logger *Logger
)

func Init() error {
	//Initialize Database
	//Initialize Environment variables
	err := InitializeEnvironment()
	if err != nil {
		return fmt.Errorf("error initializing environment variables: %v", err)
	}

	return nil
}

func GetLogger(p string) *Logger {
	//Initialize logger
	logger = NewLogger(p)

	return logger
}
