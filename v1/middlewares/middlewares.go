package middlewares

import (
	"errors"
	"github.com/tigroid3/apiswapix/v1/auth"
	"github.com/tigroid3/apiswapix/v1/response"
	"net/http"
)

func SetMiddlewareJSON(httpHandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		httpHandler(w, r)
	}
}

func SetMiddlewareAuthentication(httpHandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			response.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		httpHandler(w, r)
	}
}
