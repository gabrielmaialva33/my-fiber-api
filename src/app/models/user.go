package models

type User struct {
	Id       string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name     string
	Username string
	Email    string
	Password []byte
}
