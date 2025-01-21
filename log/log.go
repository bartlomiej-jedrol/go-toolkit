package log

import (
	"fmt"
	"log"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	Enpoint  string = "endpoint"
	Service  string = "service"
	Function string = "function"
	EnvVar   string = "environment_variable"
)

type field interface {
}

func New() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.DisableStacktrace = true
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	config.DisableCaller = false

	logger, err := config.Build()
	if err != nil {
		log.Printf("ERROR: New - failed to initiliaze zap logger")
		return nil, err
	}

	// 0 - log.go as a caller
	// 1 - actual caller
	logger = logger.WithOptions(zap.AddCallerSkip(1))
	return logger, nil
}

func buildMessage(message string, field field) string {
	if field == nil {
		return message
	}
	return fmt.Sprintf(message+": %+v", field)
}

func buildFields(message string, errorMessage error, service, function, endpoint, environmentVariable string) []zapcore.Field {
	if message == "" {
		message = "no message"
	}

	fields := []zapcore.Field{}
	if errorMessage != nil {
		fields = append(fields, zap.Error(errorMessage))
	}
	if service != "" {
		fields = append(fields, zap.String(Service, service))
	}
	if function != "" {
		fields = append(fields, zap.String(Function, function))
	}
	if endpoint != "" {
		fields = append(fields, zap.String(Enpoint, endpoint))
	}
	if environmentVariable != "" {
		fields = append(fields, zap.String(EnvVar, environmentVariable))
	}
	return fields
}

func Error(message string, field field, errorMessage error, service, function, endpoint, environmentVariable string) {
	logger, err := New()
	if err != nil {
		return
	}

	msg := buildMessage(message, field)
	fields := buildFields(msg, errorMessage, service, function, endpoint, environmentVariable)
	logger.Error(msg, fields...)
}

func Info(message string, field field, errorMessage error, service, function, endpoint, environmentVariable string) {
	logger, err := New()
	if err != nil {
		return
	}

	msg := buildMessage(message, field)
	fields := buildFields(msg, nil, service, function, endpoint, environmentVariable)
	logger.Info(msg, fields...)
}
