package log

import (
	"log"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	Enpoint  string = "endpoint"
	Service  string = "serviceName"
	Function string = "function"
	EnvVar   string = "environmentVariable"
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

func Error(msg, service, function string, endpoint, envVar *string) {
	logger, err := New()
	if err != nil {
		return
	}

	logger.Error(msg,
		zap.String(Service, service),
		zap.String(Function, function),
		zap.String("fileId", "id"))
}
