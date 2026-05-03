package middlewares

import (
	"AuthenticationService/config"
	"AuthenticationService/helper"
	"AuthenticationService/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := helper.LoggerFromContext(r.Context())
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			log.Error().Msg("no authorization token available ")
			utils.WriteError(w, http.StatusUnauthorized, "no authorization token available ")
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			log.Error().Msg("no bearer authorization token available ")
			utils.WriteError(w, http.StatusUnauthorized, "no bearer authorization token available ")
			return
		}

		authToken := strings.TrimPrefix(authHeader, "Bearer ")
		if authToken == "" {
			log.Error().Msg("no bearer authorization token available ")
			utils.WriteError(w, http.StatusUnauthorized, "no bearer authorization token available ")
			return
		}
		tokenSecret := config.GetConfig().Auth.TokenSecret
		_, err := jwt.Parse(authToken, func(t *jwt.Token) (any, error) {
			return []byte(tokenSecret), nil
		},
			jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

		if err != nil {
			log.Error().Msg(fmt.Sprintf("error parsing token %s", err.Error()))
			utils.WriteError(w, http.StatusUnauthorized, fmt.Sprintf("error parsing token %s", err.Error()))
			return 
		}

		next.ServeHTTP(w,r)

	})
}
