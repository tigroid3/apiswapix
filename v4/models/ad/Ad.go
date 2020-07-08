package ad

const (
	AD_STATUS_ON_MODEARATION         = 1
	AD_STATUS_CONFIRMED              = 2
	AD_STATUS_DECLINED               = 3
	AD_STATUS_INSTEP                 = 5
	AD_STATUS_DUPLICATE              = 6
	AD_STATUS_BANNED                 = 7
	AD_STATUS_DUPLICATE_VERIFICATION = 8
	AD_STATUS_WAIT_PARSE_PHOTO       = 9
	AD_STATUS_DELETE                 = 10
	AD_STATUS_WAIT_PAYMENT           = 11
)

const (
	AD_STATUS_USER_ACTIVE = iota + 1
	AD_STATUS_USER_INACTIVE
	AD_STATUS_USER_EXPIRED
	AD_STATUS_USER_SOLD
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
}
