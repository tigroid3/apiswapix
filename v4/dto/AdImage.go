package dto

type AdImagePresenter struct {
	ID       uint64 `json:"id"`
	AdId     uint64 `json:"-"`
	UrlSmall string `json:"size_small"`
	UrlBig   string `json:"size_big"`
	Default  bool   `json:"default"`
}

func (AdImagePresenter) TableName() string {
	return "ad_images"
}
