package helper

import (
	"AuthenticationService/constants"
	"AuthenticationService/utils"
	"context"
)

func GetCorrelationID(ctx context.Context) string {
	if id, ok := ctx.Value(constants.CorrelationIDKey).(string); ok {
		return id
	}
	return "no-correlation-id"
}

func GetPayLoad[T any](ctx context.Context) (*T, bool) {
	payload, ok := ctx.Value(utils.ValidatorContextKey).(*T)
	return payload, ok
}
