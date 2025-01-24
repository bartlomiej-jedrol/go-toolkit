package aws

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	iLog "github.com/bartlomiej-jedrol/go-toolkit/log"
)

var Service = "go-toolkit"

// LoadDefaultConfig loads default AWS config.
func LoadDefaultConfig() (*aws.Config, error) {
	function := "LoadDefaultConfig"

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		iLog.Error("failed to load default AWS config", "", err, Service, function)
		return nil, err
	}
	return &cfg, nil
}

// GetEnvironmentVariable returns environment variable's value.
func GetEnvironmentVariable(envVarName string) (string, error) {
	function := "GetEnvironmentVariable"

	ev := os.Getenv(envVarName)
	if ev == "" {
		iLog.Error("failed to get environment variable", envVarName, nil, Service, function)
		return "", nil
	}
	return ev, nil
}
