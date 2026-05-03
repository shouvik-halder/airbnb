package helper

import (
	"AuthenticationService/middlewares"
	"AuthenticationService/utils"
	"context"
)

func GetCorrelationID(ctx context.Context) string {
	if id, ok := ctx.Value(middlewares.CorrelationIDKey).(string); ok {
		return id
	}
	return "no-correlation-id"
}

func GetPayLoad[T any](ctx context.Context) (*T, bool) {
	payload, ok := ctx.Value(utils.ValidatorContextKey).(*T)
	return payload, ok
}
