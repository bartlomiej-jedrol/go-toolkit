package aws

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/joho/godotenv"
)

// LoadDefaultConfig loads default AWS config.
func LoadDefaultConfig() (*aws.Config, error) {
	log.Printf("INFO: LoadDefaultConfig - Entering LoadDefaultConfig")
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("ERROR: LoadDefaultConfig - failed to load default AWS config: %v", err)
		return nil, err
	}
	return &cfg, nil
}

// GetEnvironmentVariable returns environment variable's value.
func GetEnvironmentVariable(envVarName string) (string, error) {
	log.Printf("INFO: GetEnvironmentVariable - Entering GetEnvironmentVariable")

	// Loading .env file does not return an error because lambda in its runtime relies
	// on environment variables.
	err := godotenv.Load()
	if err != nil {
		log.Printf("ERROR: GetEnvironmentVariable - failed to load .env file, %v", err)
	}

	ev := os.Getenv(envVarName)
	if ev == "" {
		log.Printf("ERROR: GetEnvironmentVariable - failed to get environment variable: %s, %v", envVarName, err)
		return "", err
	}
	return ev, nil
}
