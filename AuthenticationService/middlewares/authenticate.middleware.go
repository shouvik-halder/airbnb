package middlewares

import (
	"AuthenticationService/config"
	dbconfig "AuthenticationService/config/db"
	"AuthenticationService/constants"
	dbrepo "AuthenticationService/db/repositories"
	"AuthenticationService/helper"
	"AuthenticationService/utils"
	"context"
	"fmt"
	"net/http"
	"strconv"
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
		claims := &jwt.RegisteredClaims{}
		token, err := jwt.ParseWithClaims(authToken, claims, func(t *jwt.Token) (any, error) {
			return []byte(tokenSecret), nil
		},
			jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

		if err != nil {
			log.Error().Msg(fmt.Sprintf("error parsing token %s", err.Error()))
			utils.WriteError(w, http.StatusUnauthorized, fmt.Sprintf("error parsing token %s", err.Error()))
			return
		}
		if !token.Valid {
			log.Error().Msg("invalid token")
			utils.WriteError(w, http.StatusUnauthorized, "invalid token")
			return
		}

		userID, err := strconv.ParseInt(claims.Subject, 10, 64)
		if err != nil || userID <= 0 {
			log.Error().Msg("invalid token subject")
			utils.WriteError(w, http.StatusUnauthorized, "invalid token subject")
			return
		}

		ctx := context.WithValue(r.Context(), constants.AuthenticatedUserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

func RequirePermission(permissionName string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID, ok := helper.GetAuthenticatedUserID(r.Context())
			if !ok {
				utils.WriteError(w, http.StatusUnauthorized, "authenticated user not found")
				return
			}

			userRolesRepo := dbrepo.NewUserRolesRepository(dbconfig.GetDB())
			hasPermission, err := userRolesRepo.HasPermission(userID, permissionName)
			if err != nil {
				utils.WriteError(w, http.StatusInternalServerError, err.Error())
				return
			}
			if !hasPermission {
				utils.WriteError(w, http.StatusForbidden, "missing required permission")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func RequireRole(roleName string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID, ok := helper.GetAuthenticatedUserID(r.Context())
			if !ok {
				utils.WriteError(w, http.StatusUnauthorized, "authenticated user not found")
				return
			}

			userRolesRepo := dbrepo.NewUserRolesRepository(dbconfig.GetDB())
			hasRole, err := userRolesRepo.HasRole(userID, roleName)
			if err != nil {
				utils.WriteError(w, http.StatusInternalServerError, err.Error())
				return
			}
			if !hasRole {
				utils.WriteError(w, http.StatusForbidden, "missing required role")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
