package middlewares

import (
	"errors"
	"gitlab.ru/new-swapix/api/v4/auth"
	"gitlab.ru/new-swapix/api/v4/response"
	"net/http"
)

func SetMiddlewareJSON(httpHandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		//w.Header().Set("Access-Control-Allow-Origin", "*")
		//w.Header().Set("Vary", "Origin")
		//w.Header().Set("Access-Control-Allow-Credentials", "true")
		//w.Header().Set("Access-Control-Expose-Headers", "Content-Length, ETag, Link, X-RateLimit-Limit, X-RateLimit-Remaining, X-Requested-With")
		//w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, HEAD, OPTIONS, PATCH")
		//w.Header().Set("Access-Control-Allow-Headers", "*")

		httpHandler(w, r)
	}
}

func SetMiddlewareAuthentication(httpHandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValidFromRequest(r)
		if err != nil {
			response.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		httpHandler(w, r)
	}
}
