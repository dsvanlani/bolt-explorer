package logger

import "go.uber.org/zap"

var logger *zap.Logger

func Logger() *zap.SugaredLogger {
	if logger == nil {
		logger, _ = zap.NewDevelopment()
	}
	return logger.Sugar()
}
