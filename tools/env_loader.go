package tools

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVariable(key string) (string, error) {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
		return "", err
	}

	result, present := os.LookupEnv(key)
	if !present {
		return "", fmt.Errorf(".env variable with key %s doesn't exist", key)
	}

	return result, nil
}
