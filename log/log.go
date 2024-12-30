package log

import (
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

func Error(msg, service, function string, endpoint, envVar *string) {
	logger, err := New()
	if err != nil {
		return
	}

	if msg == "" {
		msg = "no message"
	}

	fields := []zapcore.Field{}
	if service != "" {
		fields = append(fields, zap.String(Service, service))
	}
	if function != "" {
		fields = append(fields, zap.String(Function, function))
	}
	if endpoint != nil {
		fields = append(fields, zap.String(Enpoint, *endpoint))
	}
	if envVar != nil {
		fields = append(fields, zap.String(EnvVar, *envVar))
	}

	logger.Error(msg, fields...)
}
