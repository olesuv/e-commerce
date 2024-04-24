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

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		arw := authResponseWriter{w, "", ""}
		userEmailContextKey := ContextKey("userEmail")

		c, err := r.Cookie("auth")
		if err != nil || c == nil {
			next(&arw, r)
			return
		}

		userEmail, err := libs.ValidateJwtToken(context.Background(), c.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		arw.userEmailFromCookie = userEmail
		arw.userEmailToResolver = userEmail

		ctx := context.WithValue(r.Context(), userEmailContextKey, &arw.userEmailToResolver)

		r = r.WithContext(ctx)
		next(&arw, r)
	}
}
