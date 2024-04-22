package middleware

import (
	"context"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"server.go/libs"
)

type AuthTokenKey string

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		email, err := libs.ValidateJwtToken(token)
		if err != nil || email == "" {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), AuthTokenKey("token"), token)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func CtxValue(ctx context.Context) string {
	raw, _ := ctx.Value(AuthTokenKey("token")).(string)
	return raw
}

func Auth(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	token := CtxValue(ctx)
	if token == "" {
		return nil, &gqlerror.Error{
			Message: "Unauthorized",
		}
	}

	return next(ctx)
}
