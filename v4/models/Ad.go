package models

import (
	"github.com/jinzhu/gorm"
)

type Ad struct {
	ID                      uint64 `gorm:"primary_key;auto_increment" json:"id"`
	UserId                  uint64 `gorm:"size:64;not null;unique" json:"user_id"`
	CategoryId              uint64 `gorm:"size:64;" json:"category_id"`
	CountryId               uint64 `gorm:"size:64;" json:"country_id"`
	LocationId              uint64 `gorm:"size:64;" json:"location_id"`
	Name                    string `gorm:"size:255;not null;" json:"name"`
	Description             string `gorm:"size:64;" json:"description"`
	Status                  uint64 `gorm:"size:64;not null;" json:"status"`
	StatusUser              uint64 `gorm:"size:64;not null;default:1;" json:"status_user"`
	AddFrom                 uint64 `gorm:"size:64;not null;" json:"add_from"`
	CheckUnique             uint64 `gorm:"size:64;not null;" json:"check_unique"`
	CreatedAt               uint64 `gorm:"size:64;not null;" json:"created_at"`
	UpdatedAt               uint64 `gorm:"size:64;not null;" json:"updated_at"`
	Alias                   string `gorm:"size:255;" json:"alias"`
	ParserId                uint64 `gorm:"size:64;" json:"parser_id"`
	Cost                    uint64 `gorm:"size:64;" json:"cost"`
	ExpiresAt               uint64 `gorm:"size:64;not null;" json:"expires_at"`
	CountViews              uint64 `gorm:"size:64;not null;" json:"count_views"`
	CountViewsToday         uint64 `gorm:"size:64;" json:"count_views_today"`
	RejectReason            uint64 `gorm:"size:64;" json:"reject_reason"`
	SetStatusOnModerationAt uint64 `gorm:"size:64;not null;" json:"set_on_moderation_at"`
	CountDeclined           uint32 `gorm:"size:64;not null;" json:"count_declined"`
	FtsVector               string `gorm:"size:64;" json:"fts_vector"`
	User                    User   `gorm:"user"`
}

func (ad *Ad) Search(db *gorm.DB) ([]Ad, error) {
	ads := []Ad{}
	err := db.Where("status = ?", 2).Limit(20).Find(&ads).Error
	if err != nil {
		return nil, err
	}

	return ads, nil
}
