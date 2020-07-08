package services

import (
	"github.com/jinzhu/gorm"
	"gitlab.ru/new-swapix/api/v4/models/ad"
)

func createAdFromAdCreateForm(db *gorm.DB, form ad.AdCreateForm) (*ad.Ad, error) {
	var err error
	ad := new(ad.Ad)

	ad.Name = form.Name
	ad.Description = form.Description
	ad.CategoryId = form.CategoryId
	ad.LocationId = form.LocationId
	ad.Cost = form.Cost

	err = db.Debug().Create(&ad).Error
	if err != nil {
		return ad, err
	}
	return ad, nil
}
