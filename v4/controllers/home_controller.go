package controllers

import (
	"fmt"
	"gitlab.ru/new-swapix/api/v4/auth/identity"
	"gitlab.ru/new-swapix/api/v4/response"
	"net/http"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, "Welcome To This Awesome API")
}

func (server *Server) Test(w http.ResponseWriter, r *http.Request) {
	fmt.Println(identity.GetUserByRequest(r, server.DB))
}
