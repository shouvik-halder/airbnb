package middlewares

import (
	"AuthenticationService/config/logger"
	"AuthenticationService/constants"
	"context"
	"net/http"

	"github.com/google/uuid"
)



func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cid := r.Header.Get("X-Correlation-ID")
		if cid == "" {
			cid = uuid.New().String()
		}

		reqLogger := logger.Log.
			With().
			Str("correlation-id", cid).
			Logger()

		ctx := context.WithValue(r.Context(), constants.LoggerKey, &reqLogger)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
