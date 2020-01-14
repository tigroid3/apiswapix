package controllers

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/tigroid3/apiswapix/v1/auth"
	"github.com/tigroid3/apiswapix/v1/models"
	"github.com/tigroid3/apiswapix/v1/models/forms"
	"github.com/tigroid3/apiswapix/v1/response"
	"github.com/tigroid3/apiswapix/v1/services/userService"
	"golang.org/x/crypto/bcrypt"
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

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	userLoginForm := forms.UserLoginForm{}
	err = json.Unmarshal(body, &userLoginForm)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	validate = validator.New()
	err = validate.Struct(userLoginForm)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	token, err := server.SignIn(userLoginForm.Email, userLoginForm.Password)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	response.JSON(w, http.StatusOK, token)
}

func (server *Server) SignIn(email, password string) (map[string]string, error) {
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
	return auth.GenerateTokenPair(user.ID, user.Password)
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
	err = server.DB.Debug().Model(models.User{}).Where("id = ?", claims["user_id"]).Take(&user).Error
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, errors.New("User not found"))
	}

	err = auth.EqualsHmacForHashPassword(claims["secret"].(string), user.Password)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	tokenPairs, err := auth.GenerateTokenPair(user.ID, user.Password)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	response.JSON(w, http.StatusOK, tokenPairs)
}
