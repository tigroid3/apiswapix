package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"gitlab.ru/new-swapix/api/v4/auth"
	"gitlab.ru/new-swapix/api/v4/models"
	"gitlab.ru/new-swapix/api/v4/models/forms"
	"gitlab.ru/new-swapix/api/v4/response"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	userLoginForm := forms.UserLoginForm{}
	err = json.Unmarshal(body, &userLoginForm)
	fmt.Println(userLoginForm)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	validate = validator.New()
	err = validate.Struct(userLoginForm)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := server.signIn(userLoginForm.Email, userLoginForm.Password)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	response.JSON(w, http.StatusOK, token)
}

func (server *Server) RefreshToken(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	rtForm := forms.UserRefreshTokenForm{}
	err = json.Unmarshal(body, &rtForm)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	validate = validator.New()
	err = validate.Struct(rtForm)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	claims, err := auth.GetPayoutsFromToken(rtForm.RefreshToken)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	user := models.User{}
	err = server.DB.Debug().Model(models.User{}).Where("id = ?", claims["sub"]).Take(&user).Error
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, errors.New("User not found"))
	}

	if claims["type"].(string) != auth.TypeTokenRefresh {
		response.ERROR(w, http.StatusUnprocessableEntity, errors.New("Invalid token type"))
	}

	tokenPairs, err := auth.GenerateTokenPair(user.ID)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	response.JSON(w, http.StatusOK, tokenPairs)
}

func (server *Server) signIn(email, password string) (map[string]string, error) {
	var err error
	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return nil, err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, err
	}
	return auth.GenerateTokenPair(user.ID)
}
