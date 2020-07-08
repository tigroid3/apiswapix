package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

const (
	STATUS_NEW = iota
	STATUS_CONFIRM
	STATUS_BLOCKED
	STATUS_DELETED
	STATUS_FAKE
)

type User struct {
	ID                uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Email             string `gorm:"size:64;not null;unique" json:"email"`
	Name              string `gorm:"size:255;" json:"name"`
	Social            bool   `gorm:"default:false;not null" json:"social"`
	AutoPassword      bool   `gorm:"default:false;not null" json:"auto_password"`
	Password          string `gorm:"size:64;" json:"password"`
	Phone             string `gorm:"default:null;size:64" json:"phone"`
	FacebookUrl       string `gorm:"default:null;size:255" json:"facebook_url"`
	CountryId         uint64 `gorm:"size:64" json:"country_id"`
	LocationId        uint64 `gorm:"default:null" json:"location_id"`
	CompanyId         uint64 `gorm:"default:null" json:"company_id"`
	AuthKey           string `gorm:"default:null;size:32" json:"auth_key"`
	Status            uint64 `gorm:"not null" json:"status"`
	EmailConfirm      bool   `gorm:"default:false" json:"email_confirm"`
	EmailSubscription bool   `gorm:"default:true" json:"email_subscription"`
	AddFrom           uint64 `gorm:"not null" json:"add_from"`
	Photo             string `gorm:"default:null;size:255" json:"photo"`
	CreatedAt         uint64 `gorm:"default:0;not null" json:"created_at"`
	UpdatedAt         uint64 `gorm:"default:0;not null" json:"updated_at"`
	PhoneConfirm      uint16 `gorm:"default:0;" json:"phone_confirm"`
	LastVisitAt       uint32 `gorm:"default:null;" json:"last_visit_at"`
}

func (User) TableName() string {
	return "user"
}

func (user *User) Create(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
