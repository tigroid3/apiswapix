package controllers

import (
	"github.com/tigroid3/apiswapix/v1/response"
	"net/http"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {

	response.JSON(w, http.StatusOK, "Welcome To This Awesome API")
}

func (server *Server) Test(w http.ResponseWriter, r *http.Request) {

}
