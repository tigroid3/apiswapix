package models

type AdView struct {
	ID              uint64 `json:"id"`
	UserId          uint64 `json:"user_id"`
	CategoryId      uint64 `json:"category_id"`
	LocationId      uint64 `json:"location_id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	Status          uint64 `json:"status"`
	StatusUser      uint64 `json:"status_user"`
	CreatedAt       uint64 `json:"created_at"`
	UpdatedAt       uint64 `json:"updated_at"`
	Alias           string `json:"alias"`
	Cost            uint64 `json:"cost"`
	ExpiresAt       uint64 `json:"expires_at"`
	CountViews      uint64 `json:"count_views"`
	CountViewsToday uint64 `json:"count_views_today"`
}
