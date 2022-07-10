package middleware

import (
	"dobledcloud.com/consumers/repository"
	"dobledcloud.com/consumers/server"
	"net/http"
	"strings"
)

func CheckAuthMiddleware(s server.Server) func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			//Search a key into emission's
			apiKey := strings.TrimSpace(r.Header.Get("x-api-key"))
			emission, err := repository.GetEmissionByKey(r.Context(), apiKey)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			//Search a secret key into emission
			secretKey := strings.TrimSpace(r.Header.Get("x-api-secret"))
			exists := repository.GetSecretForEmission(r.Context(), emission.Id, secretKey)

			if !exists {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})

	}
}
