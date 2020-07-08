package dto

type AdPresenter struct {
	ID          uint64 `json:"id"`
	UserId      uint64 `json:"user_id"`
	CategoryId  uint64 `json:"category_id"`
	LocationId  uint64 `json:"location_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   uint64 `json:"created_at"`
	Alias       string `json:"alias"`
	Cost        uint64 `json:"cost"`
	CountViews  uint64 `json:"count_views"`
}
