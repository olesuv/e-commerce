package middleware

import (
	"context"
	"net/http"

	"server.go/libs"
)

type AuthTokenKey string

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("auth")
		if err != nil || c == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		userEmail, err := libs.ValidateJwtToken(context.Background(), c.Value)
		if err != nil || userEmail == "" {
			http.Error(w, "Invalid cookie", http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), AuthTokenKey("auth"), userEmail)

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func CtxValue(ctx context.Context) string {
	raw, _ := ctx.Value(AuthTokenKey("auth")).(string)
	return raw
}
