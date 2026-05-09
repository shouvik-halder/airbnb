package middlewares

import (
	"ReviewService/constants"
	"context"
	"net/http"

	"github.com/google/uuid"
)

func AttachCorrelationId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		correlationId := r.Header.Get("X-Correlation-Id")

		if correlationId == "" {
			correlationId = uuid.New().String()
		}

		ctx := context.WithValue(r.Context(), constants.CorrelationIdKey, correlationId)
		w.Header().Set("X-Correlation-Id", correlationId)
		r.WithContext(ctx)
	})
}
