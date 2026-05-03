package middlewares

import (
	"AuthenticationService/constants"
	"context"
	"net/http"

	"github.com/google/uuid"
)

func CorrelationId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		correlationID := r.Header.Get("X-Correlation-ID")

		if correlationID == "" {
			correlationID = uuid.New().String()
		}

		ctx := context.WithValue(r.Context(), constants.CorrelationIDKey, correlationID)

		w.Header().Set("X-Correlation-Id", correlationID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
