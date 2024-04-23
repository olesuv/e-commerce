package middleware

import (
	"context"
	"net/http"
	"time"

	"server.go/libs"
)

type ContextKey string

type authResponseWriter struct {
	http.ResponseWriter
	userEmailToResolver string
	userEmailFromCookie string
}

func (w *authResponseWriter) Write(b []byte) (int, error) {
	if w.userEmailFromCookie != w.userEmailToResolver {
		token, err := libs.GenearteJwtToken(context.Background(), w.userEmailToResolver)
		if err != nil || token == "" {
			return 0, err
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "auth",
			Value:    token,
			Expires:  time.Now().Add(time.Hour * 24 * 40),
			HttpOnly: true,
			Path:     "/",
		})
	}
	return w.ResponseWriter.Write(b)
}

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		arw := authResponseWriter{w, "", ""}
		userEmailContextKey := ContextKey("userEmail")

		c, err := r.Cookie("auth")
		if err != nil || c.Value == "" {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		arw.userEmailFromCookie = c.Value
		arw.userEmailToResolver = c.Value

		ctx := context.WithValue(r.Context(), userEmailContextKey, &arw.userEmailToResolver)

		r = r.WithContext(ctx)
		next.ServeHTTP(&arw, r)
	})
}
