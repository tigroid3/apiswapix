package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/form"
	"github.com/gorilla/mux"
	"gitlab.ru/new-swapix/api/v4/components/pagination"
	"gitlab.ru/new-swapix/api/v4/dto"
	"gitlab.ru/new-swapix/api/v4/models"
	"gitlab.ru/new-swapix/api/v4/models/ad"
	"gitlab.ru/new-swapix/api/v4/models/forms"
	"gitlab.ru/new-swapix/api/v4/response"
	"io/ioutil"
	"net/http"
	"net/url"
)

var decoder *form.Decoder

func (s *Server) AdIndex(w http.ResponseWriter, r *http.Request) {
	var err error

	adListDto := []dto.AdListDto{}
	err = s.DB.Select("ads.*, (SELECT id FROM favorites_ads WHERE favorites_ads.user_id = ads.user_id AND ads.id = favorites_ads.ad_id LIMIT 1)").
		Table("ads").
		Scopes(ad.GetPublishedAds, pagination.AddPagination(r)).
		Preload("Images").
		Order("created_at DESC").
		Group("ads.id").
		Find(&adListDto).Error

	if err != nil {
		fmt.Printf("%s", err)
		response.ERROR(w, http.StatusUnprocessableEntity, errors.New("Error receiving ads"))
		return
	}

	response.JSON(w, http.StatusOK, adListDto)
}

func (s *Server) AdSearch(w http.ResponseWriter, r *http.Request) {
	var err error
	adSearchForm := forms.AdSearchForm{}
	decoder = form.NewDecoder()
	values, err := url.ParseQuery(r.URL.RawQuery)

	if err != nil {
		response.ERROR(w, http.StatusBadRequest, errors.New("Incorrect url params"))
	}

	err = decoder.Decode(&adSearchForm, values)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	formatterErrors, err := adSearchForm.Validate()

	if err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, formatterErrors)
		return
	}

	adListsDto := []dto.AdListDto{}

	query := adSearchForm.Search(s.DB, r)
	err = query.Find(&adListsDto).Error

	if err != nil {
		fmt.Printf("%s", err)
		response.ERROR(w, http.StatusUnprocessableEntity, errors.New("Error receiving ads"))
		return
	}

	response.JSON(w, http.StatusOK, adListsDto)
}

func (s *Server) AdView(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	adView := models.AdView{}
	err := s.DB.Table("ads").Where("id = ?", vars["id"]).Find(&adView).Error

	if err != nil {
		response.JSON(w, http.StatusNotFound, nil)
		return
	}

	response.JSON(w, http.StatusOK, adView)
}

func (s *Server) AdCreate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	adCreateForm := ad.AdCreateForm{}

	if err != nil {
		response.JSON(w, http.StatusBadRequest, nil)
		return
	}

	err = json.Unmarshal(body, &adCreateForm)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	formatterErrors, err := adCreateForm.Validate(s.DB)

	if err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, formatterErrors)
		return
	}

	//tx := s.DB.Begin()
	//
	//
	//
	//tx.Rollback()
	//tx.Commit()

	//транзакция
	//создание объявления
	//линк картинок
	//линк фильтров

	//транзакция завершена

}

func (s *Server) AdDelete(w http.ResponseWriter, r *http.Request) {
	//написать метод получения struct user by auth token
}

func (s *Server) AdUpdate(w http.ResponseWriter, r *http.Request) {

}
