package dto

//Презентер для страницы главная/поиск
type AdListDto struct {
	ID         uint64             `json:"id"`
	UserId     uint64             `json:"user_id"`
	CategoryId uint64             `json:"category_id"`
	LocationId uint64             `json:"location_id"`
	Status     uint64             `json:"status"`
	Name       string             `json:"name"`
	FavoriteId uint64             `json:"favorite_id"`
	Alias      string             `json:"alias"`
	Cost       uint64             `json:"cost"`
	Images     []AdImagePresenter `gorm:"foreignkey:ad_id" json:"images"`
}

//
func (AdListDto) TableName() string {
	return "ads"
}
