package services

import (
	"github.com/jinzhu/gorm"
	"gitlab.ru/new-swapix/api/v4/models"
)

func CreateAdImageFromStruct(db *gorm.DB, adImage *models.AdImage) (*models.AdImage, error) {
	var err error

	err = db.Debug().Create(adImage).Error
	if err != nil {
		return adImage, err
	}
	return adImage, nil
}
