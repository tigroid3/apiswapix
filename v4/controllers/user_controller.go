package controllers

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gitlab.ru/new-swapix/api/v4/dto"
	"gitlab.ru/new-swapix/api/v4/models"
	"gitlab.ru/new-swapix/api/v4/models/forms"
	"gitlab.ru/new-swapix/api/v4/response"
	"gitlab.ru/new-swapix/api/v4/services"
	"io/ioutil"
	"net/http"
)

var validate *validator.Validate

func (s *Server) UserCreate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	userRegisterForm := forms.UserRegisterForm{}
	err = json.Unmarshal(body, &userRegisterForm)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	userRegisterForm.Prepare()
	validate = validator.New()
	err = validate.Struct(userRegisterForm)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	user, err := services.CreateFromForm(s.DB, userRegisterForm)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	response.JSON(w, http.StatusOK, user)
}

func (s *Server) UserUpdate(w http.ResponseWriter, r *http.Request) {
	//body, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	response.ERROR(w, http.StatusUnprocessableEntity, err)
	//}
}

func (s *Server) UserView(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	modelUser := models.User{}
	err := s.DB.Table("user").Where("id = ?", vars["id"]).Find(&modelUser).Error

	if err != nil {
		response.JSON(w, http.StatusNotFound, nil)
		return
	}

	userPresenter := dto.UserPresenter{}
	userPresenter.LoadFromModel(modelUser)

	response.JSON(w, http.StatusOK, userPresenter)
}
