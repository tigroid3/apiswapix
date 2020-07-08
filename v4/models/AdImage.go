package models

type AdImage struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	AdId      uint64 `gorm:"size:64;null;omitempty" json:"ad_id"`
	Original  string `gorm:"size:255;" json:"original"`
	Default   bool   `gorm:"not null;" json:"default"`
	Published bool   `gorm:"not null;" json:"published"`
	Hash      string `gorm:"size:32;" json:"hash"`
	CreatedAt uint64 `gorm:"size:64;not null" json:"created_at"`
	UpdatedAt uint64 `gorm:"size:64;not null;" json:"updated_at"`
}

func (adImage AdImage) getUrls() {
	//тут будет получение урлов по Original, надо реализовать компонент settings для получения инфы из БД
}
