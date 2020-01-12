package controllers

import (
	"fmt"
	"github.com/tigroid3/apiswapix/v1/auth"
	"github.com/tigroid3/apiswapix/v1/response"
	"net/http"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {

	response.JSON(w, http.StatusOK, "Welcome To This Awesome API")
}

func (server *Server) Test(w http.ResponseWriter, r *http.Request) {
	passHash := "$2y$13$6Qm7Bf2UAHMCGLVBZMMTS.F5AwC79uAe2PBODH8H61iFC.17OaKri"
	claims, _ := auth.GetPayoutsFromToken(auth.ExtractToken(r))

	fmt.Println(auth.EqualsHmacForHashPassword(claims["secret"].(string), passHash))
}
