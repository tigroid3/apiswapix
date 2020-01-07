package userService

import (
	"github.com/jinzhu/gorm"
	"github.com/tigroid3/apiswapix/v1/models"
	"github.com/tigroid3/apiswapix/v1/models/forms"
)

func CreateFromForm(db *gorm.DB, form forms.UserRegisterForm) (*models.User, error) {
	var err error
	u := new(models.User)
	u.Name = form.Name
	u.Password = form.Password
	u.Phone = form.Phone
	u.Email = form.Email

	err = db.Debug().Create(&u).Error
	if err != nil {
		return u, err
	}
	return u, nil
}
