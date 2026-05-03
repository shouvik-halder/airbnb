package helper

import (
	"AuthenticationService/config/logger"
	"AuthenticationService/constants"
	"context"

	"github.com/rs/zerolog"
)

func LoggerFromContext(ctx context.Context) *zerolog.Logger {
	if log, ok := ctx.Value(constants.LoggerKey).(*zerolog.Logger); ok {
		return log
	}
	return logger.Log
}