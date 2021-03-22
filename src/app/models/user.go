package models

import "github.com/gofrs/uuid"

type User struct {
	Id       uuid.Generator
	Name     string
	Nickname string
	Email    string
	Password string
}
