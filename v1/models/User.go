package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
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
	CountryId         int64  `gorm:"size:64" json:"country_id"`
	LocationId        int64  `gorm:"default:null" json:"location_id"`
	CompanyId         int64  `gorm:"default:null" json:"company_id"`
	AuthKey           string `gorm:"default:null;size:32" json:"auth_key"`
	Status            int64  `gorm:"not null" json:"status"`
	EmailConfirm      bool   `gorm:"default:false" json:"email_confirm"`
	EmailSubscription bool   `gorm:"default:true" json:"email_subscription"`
	AddFrom           int64  `gorm:"not null" json:"add_from"`
	Photo             string `gorm:"default:null;size:255" json:"photo"`
	CreatedAt         int64  `gorm:"default:0;not null" json:"created_at"`
	UpdatedAt         int64  `gorm:"default:0;not null" json:"updated_at"`
	PhoneConfirm      int16  `gorm:"default:0;" json:"phone_confirm"`
	LastVisitAt       int32  `gorm:"default:null;" json:"last_visit_at"`
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

//
//func (u *User) SaveUser(db *gorm.DB) (*User, error) {
//	var err error
//	err = db.Debug().Create(&u).Error
//	if err != nil {
//		return &User{}, err
//	}
//	return u, nil
//}
//
//func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
//	var err error
//	users := []User{}
//	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
//	if err != nil {
//		return &[]User{}, err
//	}
//	return &users, err
//}
//
//func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
//	var err error
//	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
//	if err != nil {
//		return &User{}, err
//	}
//	if gorm.IsRecordNotFoundError(err) {
//		return &User{}, errors.New("User Not Found")
//	}
//	return u, err
//}
//
//func (u *User) UpdateAUser(db *gorm.DB, uid uint32) (*User, error) {
//
//	// To hash the password
//	err := u.BeforeSave()
//	if err != nil {
//		log.Fatal(err)
//	}
//	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
//		map[string]interface{}{
//			"password":  u.Password,
//			"nickname":  u.Name,
//			"email":     u.Email,
//			"update_at": time.Now(),
//		},
//	)
//	if db.Error != nil {
//		return &User{}, db.Error
//	}
//	// This is the display the updated user
//	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
//	if err != nil {
//		return &User{}, err
//	}
//	return u, nil
//}
