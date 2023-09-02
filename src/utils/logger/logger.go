package logger

import (
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

func init() {
	logger = zap.Must(zap.NewDevelopment())
}

func Logger() *zap.Logger {
	return logger
}
