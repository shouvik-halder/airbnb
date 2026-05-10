package validators

import (
	"AuthenticationService/utils"
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/go-chi/chi/v5"
)

func Validate[T any]() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var payload T

			if err := utils.ReadJSONBody(r, &payload); err != nil {
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

func ValidateParams[T any]() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var payload T

			v := reflect.ValueOf(&payload).Elem()
			t := v.Type()
			for i := range t.NumField() {
				field := t.Field(i)
				paramName := field.Tag.Get("json")
				if paramName == "" {
					continue
				}
				paramValue := chi.URLParam(r, paramName)
				if paramValue == "" {
					continue
				}
				ptr := reflect.New(field.Type)
				if _, err := fmt.Sscan(paramValue, ptr.Interface()); err != nil {
					utils.WriteError(w, http.StatusBadRequest, fmt.Sprintf("Invalid param '%s'", paramName))
					return
				}
				v.Field(i).Set(ptr.Elem())
			}

			if err := utils.Validator.Struct(payload); err != nil {
				utils.WriteError(w, http.StatusUnprocessableEntity, fmt.Sprintf("Validation error: %s", err.Error()))
				return
			}
			ctx := context.WithValue(r.Context(), utils.ParamsContextKey, &payload)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
