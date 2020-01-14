package controllers

import (
	"github.com/tigroid3/apiswapix/v1/models"
	"github.com/tigroid3/apiswapix/v1/response"
	"go/types"
	"net/http"
)

func (s *Server) AdIndex(w http.ResponseWriter, r *http.Request) {
	ad := models.Ad{}
	ads, err := ad.Search(s.DB)
	if err != nil {
		response.JSON(w, http.StatusOK, new(types.Map))
	}

	response.JSON(w, http.StatusOK, ads)
}
