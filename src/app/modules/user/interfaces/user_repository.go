package interfaces

import "go-api/src/app/modules/user/models"

type UserRepository interface {
	Index() ([]models.User, error)
	Show(id string) (*models.User, error)
	Create(*models.User) (*models.User, map[string]string)
	Update(*models.User) (*models.User, map[string]string)
	Delete(id string) error
	FindByEmail(email string) (*models.User, error)
	GetUserByEmailAndPassword(*models.User) (*models.User, map[string]string)
}
