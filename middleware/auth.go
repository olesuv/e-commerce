package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"server.go/libs"
)

type AuthTokenKey string

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			next.ServeHTTP(w, r)
			return
		}

		bearer := "Bearer "
		auth = auth[len(bearer):]

		email, err := libs.ValidateJwtToken(context.Background(), auth)
		if err != nil || email == "" {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), AuthTokenKey("auth"), fmt.Sprintf("Bearer %s", auth))
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func CtxValue(ctx context.Context) string {
	raw, _ := ctx.Value(AuthTokenKey("auth")).(string)
	return raw
}

func Auth(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	auth := CtxValue(ctx)
	if auth == "" {
		return nil, &gqlerror.Error{
			Message: "Unauthorized",
		}
	}

	return next(ctx)
}
