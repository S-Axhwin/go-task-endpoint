package middleware

import (
	"context"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type ctxKey string

const UserIdKey ctxKey = "user_email"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("access_token")
		if err != nil {
			http.Error(w, "unauthorized: no token", http.StatusUnauthorized)
			return
		}

		tokenStr := cookie.Value
		claims := &jwt.RegisteredClaims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "unauthorized: token invalid", http.StatusUnauthorized)
			return
		}
		if claims.Subject == "" {
			http.Error(w, "unauthorized: missing email", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIdKey, claims.Subject)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
