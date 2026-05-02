package helper

import (
	"AuthenticationService/middlewares"
	"context"
)

func GetCorrelationID(ctx context.Context) string {
	if id, ok := ctx.Value(middlewares.CorrelationIDKey).(string); ok {
		return id
	}
	return ""
}
