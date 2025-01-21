package aws

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	iLog "github.com/bartlomiej-jedrol/go-toolkit/log"
	"github.com/joho/godotenv"
)

var service = "go-toolkit"

// LoadDefaultConfig loads default AWS config.
func LoadDefaultConfig() (*aws.Config, error) {
	function := "LoadDefaultConfig"
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		iLog.Error("failed to load default AWS config", "", err, service, function)
		return nil, err
	}
	return &cfg, nil
}

// GetEnvironmentVariable returns environment variable's value.
func GetEnvironmentVariable(envVarName string) (string, error) {
	function := "GetEnvironmentVariable"
	// Loading .env file does not return an error because lambda in its runtime relies
	// on environment variables.
	err := godotenv.Load()
	if err != nil {
		iLog.Error("failed to load .env file", nil, err, service, function)
	}

	ev := os.Getenv(envVarName)
	if ev == "" {
		iLog.Error("failed to get environment variable", nil, err, service, function)
		return "", err
	}
	return ev, nil
}
