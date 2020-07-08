package forms

import (
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"gitlab.ru/new-swapix/api/v4/components/pagination"
	"gitlab.ru/new-swapix/api/v4/components/validation"
	"gitlab.ru/new-swapix/api/v4/models/ad"
	"net/http"
)

type AdSearchForm struct {
	CategoryId    uint64   `validate:"omitempty,gte=1,lte=2148000000" form:"category_id"`
	LocationId    []uint64 `validate:"omitempty,dive,gte=1,lte=2148000000" form:"location_id[]"`
	CreatedAtFrom uint64   `validate:"omitempty,gte=0,lte=2148000000" form:"created_at_from"`
	CostFrom      uint64   `validate:"omitempty,gte=0,lte=2148000000" form:"cost_from"`
	CostTo        uint64   `validate:"omitempty,gte=1,lte=2148000000" form:"cost_to"`
	FilterId      []uint64 `validate:"omitempty,dive,gte=1,lte=2148000000" form:"filter_id[]"`
	Sort          string   `validate:"omitempty" form:"sort"`
}

var validate *validator.Validate

func (adSearchForm AdSearchForm) Search(db *gorm.DB, r *http.Request) *gorm.DB {
	q := db.Select("ads.*, (SELECT id FROM favorites_ads WHERE favorites_ads.user_id = ads.user_id AND ads.id = favorites_ads.ad_id LIMIT 1)").
		Scopes(ad.GetPublishedAds, pagination.AddPagination(r), adSearchForm.AddSort).
		Joins("LEFT JOIN billing_service_hop_up_ad_order ON ads.id = billing_service_hop_up_ad_order.ad_id")

	if adSearchForm.CategoryId != 0 {
		q = q.Where("ads.category_id = ?", adSearchForm.CategoryId)
	}

	if adSearchForm.LocationId != nil {
		q = q.Where("ads.location_id IN (?)", adSearchForm.LocationId)
	}

	if adSearchForm.CostFrom != 0 {
		q = q.Where("ads.cost > ?", adSearchForm.CostFrom)
	}

	if adSearchForm.CostTo != 0 {
		q = q.Where("ads.cost < ?", adSearchForm.CostTo)
	}

	if adSearchForm.CreatedAtFrom != 0 {
		q = q.Where("ads.created_at > ?", adSearchForm.CreatedAtFrom)
	}

	if adSearchForm.FilterId != nil {
		q = q.
			Joins("INNER JOIN filter_ad ON filter_ad.ad_id = ads.id").
			Where("filter_ad.ad_id IN (?)", adSearchForm.FilterId)
	}

	return q
}

func (adSearchForm AdSearchForm) AddSort(db *gorm.DB) *gorm.DB {
	switch adSearchForm.Sort {
	case "cost":
		return db.Order("cost ASC")
	case "-cost":
		return db.Order("cost DESC")
	case "created_at":
		return db.Order("CASE WHEN hop_up_at IS NULL THEN ads.created_at ELSE hop_up_at END DESC")
	case "-created_at":
		return db.Order("CASE WHEN hop_up_at IS NULL THEN ads.created_at ELSE hop_up_at END ASC")
	case "views":
		return db.Order("count_views DESC")
	default:
		return db.Order("CASE WHEN hop_up_at IS NULL THEN ads.created_at ELSE hop_up_at END DESC")
	}
}

func (adSearchForm AdSearchForm) Validate() (map[string]string, error) {
	validate := validator.New()

	return validation.CheckStruct(validate, adSearchForm)
}
