package dto

import (
	"gitlab.ru/new-swapix/api/v4/models"
	"os"
)

type UserPresenter struct {
	ID           uint64 `json:"id"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	LocationId   uint64 `json:"location_id"`
	Status       uint64 `json:"status"`
	EmailConfirm bool   `json:"email_confirm"`
	PhoneConfirm uint16 `json:"phone_confirm"`
	Photo        string `json:"photo"`
	//LanguageId   uint64 `json:"language_id"`
}

func (userPresenter *UserPresenter) LoadFromModel(userModel models.User) *UserPresenter {
	userPresenter.ID = userModel.ID
	userPresenter.Email = userModel.Email
	userPresenter.Name = userModel.Name
	userPresenter.Phone = userModel.Phone
	userPresenter.LocationId = userModel.LocationId
	userPresenter.Status = userModel.Status
	userPresenter.EmailConfirm = userModel.EmailConfirm
	userPresenter.PhoneConfirm = userModel.PhoneConfirm
	userPresenter.Photo = getAvatarUrl(userModel.Photo)

	return userPresenter
}

func getAvatarUrl(avatar string) string {
	return os.Getenv("FRONTEND_URL") + "avatars/" + avatar
}
