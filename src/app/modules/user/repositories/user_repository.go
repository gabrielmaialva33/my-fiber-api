package repositories

import (
	"go-api/src/app/modules/user/interfaces"
	"go-api/src/app/modules/user/models"
)

// - structs
type userRepository struct {
	ur interfaces.UserRepository
}

var _ UserRepositoryInterface = &userRepository{}

// - interfaces
type UserRepositoryInterface interface {
	Index() ([]models.User, error)
	Show(id string) (*models.User, error)
	Create(*models.User) (*models.User, map[string]string)
	Update(*models.User) (*models.User, map[string]string)
	Delete(id string) error
	FindByEmail(email string) (*models.User, error)
	GetUserByEmailAndPassword(*models.User) (*models.User, map[string]string)
}

// - functions
func (u *userRepository) Index() ([]models.User, error) {
	return u.ur.Index()
}

func (u *userRepository) Show(id string) (*models.User, error) {
	return u.ur.Show(id)
}

func (u *userRepository) Create(user *models.User) (*models.User, map[string]string) {
	return u.ur.Create(user)
}

func (u *userRepository) Update(user *models.User) (*models.User, map[string]string) {
	return u.ur.Update(user)
}

func (u *userRepository) Delete(id string) error {
	return u.ur.Delete(id)
}

func (u *userRepository) FindByEmail(email string) (*models.User, error) {
	return u.ur.FindByEmail(email)
}

func (u *userRepository) GetUserByEmailAndPassword(user *models.User) (*models.User, map[string]string) {
	return u.ur.GetUserByEmailAndPassword(user)
}
