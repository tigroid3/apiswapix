package identity

import (
	"github.com/jinzhu/gorm"
	"gitlab.ru/new-swapix/api/v4/auth"
	"gitlab.ru/new-swapix/api/v4/models"
	"net/http"
)

type User struct {
	isGuest bool
	models.User
}

func GetUserByRequest(r *http.Request, db *gorm.DB) (*User, error) {
	var err error
	user := User{}
	userId, err := auth.GetUserIdByRequest(r)
	if err != nil {
		return &user, err
	}

	err = db.Table("user").Where("id = ?", userId).Find(&user).Error
	if err != nil {
		return &user, err
	}

	return &user, nil
}
