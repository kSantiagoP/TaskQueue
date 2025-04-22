package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func InitializeEnvironment() error {
	logger := GetLogger("envs")
	err := godotenv.Load(".env")
	if err != nil {
		logger.Errorf("Error loading .env file: %v", err)
		return err
	}

	return nil
}

func GetPort() (string, error) {
	//port := ":"
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		return "", fmt.Errorf("Error retrieving .env")
	}

	return ":" + port, nil
}
