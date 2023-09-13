package log

import (
	"go.uber.org/zap"
)

func NewLogger() *zap.Logger {
	logger := zap.NewExample()
	defer logger.Sync()
	logger.Info("Log initialized")

	return logger
}
