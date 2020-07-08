package pagination

import (
	"github.com/gorilla/schema"
	"github.com/jinzhu/gorm"
	"github.com/prometheus/common/log"
	"net/http"
)

var defaultLimit uint16 = 20

type Pagination struct {
	Limit  uint16 `schema:"limit"`
	Offset uint16 `schema:"offset"`
}

func AddPagination(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if err := r.ParseForm(); err != nil {
			log.Warn(err)
			return db.Limit(defaultLimit).Offset(0)
		}

		pagination := Pagination{}
		if err := schema.NewDecoder().Decode(&pagination, r.Form); err != nil {
			log.Warn(err)
			return db.Limit(defaultLimit).Offset(0)
		}

		q := db

		if pagination.Limit == 0 {
			q = q.Limit(defaultLimit)
		} else {
			q = q.Limit(pagination.Limit)
		}

		if pagination.Offset != 0 {
			q = q.Offset(pagination.Offset)
		}

		return q
	}
}
