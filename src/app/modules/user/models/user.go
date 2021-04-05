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
	Password  string     `gorm:"column:password;type:varchar(255);not null" json:"password"`
	Age       uint8      `json:"age"`
	Birthday  *time.Time `json:"birthday"`
	Avatar    *string    `json:"avatar" gorm:"column:image"`
	IsDeleted bool       `json:"-" gorm:"type:boolean;default:false"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type PublicUser struct {
	Id       string `gorm:"primary_key;default:uuid_generate_v4()" json:"id"`
	Name     string `gorm:"column:name;type:varchar(100)" json:"name"`
	Username string `gorm:"column:username;type:varchar(50)" json:"username"`
}

type Users []User

func (users Users) PublicUsers() []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.PublicUser()
	}
	return result
}

func (u *User) PublicUser() interface{} {
	return &PublicUser{
		Id:       u.Id,
		Name:     u.Name,
		Username: u.Username,
	}
}

// hook -> BeforeSave
func (u *User) BeforeSave(*gorm.DB) (err error) {
	hashPassword, err := utils.Hash(u.Password)
	if err != nil {
		return err
	}

	u.Password = string(hashPassword)

	return nil
}
