package models

import (
	"go-api/src/app/utils"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        string     `gorm:"primary_key;default:uuid_generate_v4()" json:"id"`
	Name      string     `gorm:"column:name;type:varchar(100)" json:"name"`
	Username  string     `gorm:"column:username;type:varchar(50)" json:"username"`
	Email     string     `gorm:"column:email;size:100;not null;unique;unique;index;" json:"email"`
	Password  string     `gorm:"column:password;type:varchar(255);not null" json:"password,omitempty"`
	Age       uint8      `json:"age"`
	Birthday  *time.Time `json:"birthday"`
	Avatar    *string    `json:"avatar" gorm:"column:image"`
	IsDeleted bool       `json:"-" gorm:"type:boolean;default:false"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type PublicUser struct {
	Id        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string `gorm:"size:100;not null;" json:"first_name"`
	LastName  string `gorm:"size:100;not null;" json:"last_name"`
}

// hook -> BeforeCreate
func (u *User) BeforeCreate(*gorm.DB) (err error) {
	hashPassword, err := utils.Hash(u.Password)
	if err != nil {
		return err
	}

	u.Password = string(hashPassword)

	return nil
}
