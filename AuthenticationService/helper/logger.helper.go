package helper

import (
	"AuthenticationService/config/logger"
	"AuthenticationService/middlewares"
	"context"

	"github.com/rs/zerolog"
)

func LoggerFromContext(ctx context.Context) *zerolog.Logger {
	if log, ok := ctx.Value(middlewares.LoggerKey).(*zerolog.Logger); ok {
		return log
	}
	return logger.Log
}