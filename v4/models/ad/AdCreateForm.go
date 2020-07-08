package ad

import (
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"gitlab.ru/new-swapix/api/v4/components/validation"
)

type AdCreateForm struct {
	Name        string   `validate:"required,lte=64" json:"name"`
	Description string   `validate:"required,lte=3000" json:"description"`
	CategoryId  uint64   `validate:"required,numeric,max=2147000000,exists-category-id" json:"category_id"`
	LocationId  uint64   `validate:"required,numeric,max=2147000000,exists-location-id" json:"location_id"`
	Cost        uint64   `validate:"required,min=0,max=2147000000" json:"cost"`
	ImageIds    []uint64 `validate:"required" json:"image_ids"` //провалидировать в цикле весь масив на инт
	FilterIds   []uint64 `validate:"required" json:"filter_ids"`
}

var validate *validator.Validate
var db *gorm.DB

func (adCreateForm AdCreateForm) Validate(database *gorm.DB) (map[string]string, error) {
	db = database
	validate := validator.New()
	validate.RegisterValidation("exists-category-id", existsCategoryId)
	validate.RegisterValidation("exists-location-id", existsLocationId)

	return validation.CheckStruct(validate, adCreateForm)
}

func existsCategoryId(fl validator.FieldLevel) bool {
	cnt := 0
	db.Debug().Table("categories").Where("id = ?", fl.Field().Uint()).Count(&cnt)
	return cnt > 0
}

func existsLocationId(fl validator.FieldLevel) bool {
	cnt := 0
	db.Debug().Table("locations").Where("id = ?", fl.Field().Uint()).Count(&cnt)
	return cnt > 0
}

/*func existsImageId(fl validator.FieldLevel) bool {
	cnt := 0
	db.Debug().Table("ad_images").Where("id = ?", fl.Field().Uint()).Count(&cnt)
	return cnt > 0
}

func existsFilterId(fl validator.FieldLevel) bool {
	cnt := 0
	db.Debug().Table("filters").Where("id = ?", fl.Field().Uint()).Count(&cnt)
	return cnt > 0
}
*/
