package controllers

import (
	"errors"
	"github.com/corona10/goimagehash"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gitlab.ru/new-swapix/api/v4/auth"
	"gitlab.ru/new-swapix/api/v4/components/image"
	"gitlab.ru/new-swapix/api/v4/dto"
	"gitlab.ru/new-swapix/api/v4/models"
	"gitlab.ru/new-swapix/api/v4/response"
	"gitlab.ru/new-swapix/api/v4/services"
	"image/jpeg"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"time"
)

func (s *Server) AdImageCreate(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		response.ERROR(w, 400, errors.New("Field 'file' not found"))
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)

	if err != nil {
		response.ERROR(w, 400, errors.New("Could not read file"))
		return
	}
	contentType := http.DetectContentType(fileBytes)
	fileExt, err := mime.ExtensionsByType(contentType)

	if err != nil {
		response.ERROR(w, 400, errors.New("Could not get extention"))
		return
	}
	if contentType != "image/jpeg" {
		response.ERROR(w, 400, errors.New("Incorrect format. Allow only jpeg"))
		return
	}

	fileName := uuid.New().String() + fileExt[0]
	savePath := image.PreparePathByDirAndFilename("ads", fileName)
	newFile, err := os.Create(savePath)

	if err != nil {
		response.ERROR(w, 400, err)
		return
	}
	defer newFile.Close()
	if _, err := newFile.Write(fileBytes); err != nil {
		response.ERROR(w, 400, err)
		return
	}

	openSavedFile, _ := os.Open(savePath)
	decodedImg, err := jpeg.Decode(openSavedFile)
	if err != nil {
		response.ERROR(w, 400, err)
		return
	}
	defer openSavedFile.Close()
	imageHash, err := goimagehash.PerceptionHash(decodedImg)

	if err != nil {
		response.ERROR(w, 400, err)
		return
	}

	adImage := new(models.AdImage)
	//adImage.AdId временно удалил fk на ad_id ибо не смог записать NULL в поле
	adImage.Original = fileName
	adImage.Default = false
	adImage.Published = true            //TODO возможно поле подлежит удалению
	adImage.Hash = imageHash.ToString() //убедиться в правильности выбора хэша, ибо отличается от старого
	adImage.CreatedAt = uint64(time.Now().Unix())
	adImage.UpdatedAt = uint64(time.Now().Unix())

	serviceAdImage, err := services.CreateAdImageFromStruct(s.DB, adImage)
	if err != nil {
		response.ERROR(w, 400, err)
		return
	}

	response.JSON(w, 200, serviceAdImage)
}

func (s *Server) AdImageDelete(w http.ResponseWriter, r *http.Request) {
	var err error
	imageId := mux.Vars(r)["id"]
	userId, err := auth.GetUserIdByRequest(r)
	if err != nil {
		response.ERROR(w, 400, err)
		return
	}

	adImage := models.AdImage{}
	result := s.DB.Where("ad_images.id = ? AND (ads.user_id = ? OR ad_id = 0)", imageId, userId).
		Joins("LEFT JOIN ads ON ads.id = ad_id").
		First(&adImage)

	if result.RecordNotFound() == true {
		response.ERROR(w, 400, errors.New("Ad image not found"))
		return
	}

	resultDelete := s.DB.Where("id = ?", adImage.ID).Delete(&adImage).Error
	if resultDelete != nil {
		response.ERROR(w, 400, errors.New("Error deleting db record id"+string(adImage.ID)))
		return
	}

	filePath := image.PreparePathByDirAndFilename("ads", adImage.Original)

	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		return
	}

	err = os.Remove(filePath)
	if err != nil {
		response.ERROR(w, 400, err)
		return
	}
}

func (s *Server) AdImageSetDefault(w http.ResponseWriter, r *http.Request) {
	var err error
	imageId := mux.Vars(r)["id"]
	userId, err := auth.GetUserIdByRequest(r)
	if err != nil {
		response.ERROR(w, 400, err)
		return
	}

	adImage := models.AdImage{}
	result := s.DB.Debug().Where("ad_images.id = ? AND ads.user_id = ? AND "+
		"(ad_id != 0 AND ad_id IS NOT NULL)", imageId, userId).
		Joins("INNER JOIN ads ON ads.id = ad_id").First(&adImage)

	if result.RecordNotFound() == true {
		response.ERROR(w, 400, errors.New("Ad image not found"))
		return
	}

	s.DB.Debug().Table("ad_images").Where("ad_id = ?", adImage.AdId).Update("default", false)
	s.DB.Debug().Table("ad_images").Where("id = ?", adImage.ID).Update("default", true)

	response.JSON(w, 200, adImage)
}

//метод получения картинок по id объявления
func (s *Server) AdImageGetForAd(w http.ResponseWriter, r *http.Request) {
	adId := mux.Vars(r)["id"]
	adImages := []dto.AdImagePresenter{}
	result := s.DB.Debug().Table("ad_images").Where("ad_id = ?", adId).Find(&adImages)

	if result.RecordNotFound() == true {
		response.JSON(w, 200, []string{})
		return
	}

	response.JSON(w, 200, adImages)
}
