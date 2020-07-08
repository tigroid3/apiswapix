package ad

import (
	"github.com/jinzhu/gorm"
	"gitlab.ru/new-swapix/api/v4/models"
)

//* Так уж вышло, что страница Index и Search используют один и тот же метод: ad/index
//* Но сортировка на index должна быть по дате создания (created_at),
//* а при поиске по дате поднятия (hop_up_at).
//* Пока не переделают, будет такое "обходное" средство.

func GetPublishedAds(db *gorm.DB) *gorm.DB {
	q := db.Debug().
		Joins("INNER JOIN categories ON categories.id = ads.category_id").
		Joins("INNER JOIN \"user\" ON \"user\".id = ads.user_id").
		Joins("INNER JOIN locations ON locations.id = ads.location_id").
		Joins("INNER JOIN ad_images ON ad_images.ad_id = ads.id").
		Where("ads.status = ? AND ads.status_user = ?", AD_STATUS_CONFIRMED, AD_STATUS_USER_ACTIVE).
		Where("ads.name IS NOT NULL").
		Where("locations.visible = TRUE AND locations.is_deleted = FALSE").
		Where("categories.visible = TRUE AND categories.is_deleted = FALSE").
		Where("\"user\".status != ?", models.STATUS_BLOCKED)

	return q
}
