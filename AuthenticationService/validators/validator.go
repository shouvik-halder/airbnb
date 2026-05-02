package validators

import (
	"AuthenticationService/utils"
	"context"
	"fmt"
	"net/http"
)

func Validate[T any]() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var payload T

			if err := utils.ReadJSON(r, &payload); err != nil {
				utils.WriteError(w, http.StatusBadRequest, fmt.Sprintf("Validation error %s", err.Error()))
				return
			}

			if err := utils.Validator.Struct(payload); err != nil {
				utils.WriteError(w, http.StatusUnprocessableEntity, fmt.Sprintf("Validation error %s", err.Error()))
				return
			}

			ctx := context.WithValue(r.Context(), utils.ValidatorContextKey, &payload)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
