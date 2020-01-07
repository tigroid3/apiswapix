package controllers

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/tigroid3/apiswapix/v1/models/forms"
	"github.com/tigroid3/apiswapix/v1/response"
	"github.com/tigroid3/apiswapix/v1/services/userService"
	"io/ioutil"
	"net/http"
)

var validate *validator.Validate

func (server *Server) Register(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	userRegisterForm := forms.UserRegisterForm{}
	json.Unmarshal(body, &userRegisterForm)

	userRegisterForm.Prepare()
	validate = validator.New()
	err = validate.Struct(userRegisterForm)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	user, err := userService.CreateFromForm(server.DB, userRegisterForm)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	response.JSON(w, http.StatusOK, user)
}
