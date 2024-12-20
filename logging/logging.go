package logging

import (
	"log"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	Service  string = "serviceName"
	Enpoint  string = "endpoint"
	EnvVar   string = "environmentVariable"
	Function string = "function"
)

func New() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.DisableStacktrace = true
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	logger, err := config.Build()
	if err != nil {
		log.Printf("ERROR: New - failed to initiliaze zap logger")
		return nil, err
	}
	return logger, nil
}
