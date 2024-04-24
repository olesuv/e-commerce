package middleware

import (
	"context"
	"net/http"

	"server.go/libs"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie("auth")
			if err != nil || c.Value == "" {
				next.ServeHTTP(w, r)
				// http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			userEmail, err := libs.ValidateJwtToken(context.Background(), c.Value)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), userCtxKey, userEmail)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func CtxValue(ctx context.Context) string {
	raw, _ := ctx.Value(userCtxKey).(string)
	return raw
}
