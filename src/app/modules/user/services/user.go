package services

import (
	"errors"
	"go-api/src/app/modules/user/interfaces"
	"go-api/src/app/modules/user/models"
	"go-api/src/app/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strings"
)

type UserRepo struct {
	db *gorm.DB
}

func UserServices(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

var _ interfaces.UserRepository = &UserRepo{}

func (r UserRepo) Index() ([]models.User, error) {
	panic("implement me")
}

func (r UserRepo) Show(id string) (*models.User, error) {
	var user models.User
	err := r.db.Debug().Where("id = ?", id).Take(&user).Error
	if err != nil {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (r UserRepo) Create(user *models.User) (*models.User, map[string]string) {
	dbErr := map[string]string{}
	err := r.db.Debug().Create(&user).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["email_taken"] = "email already taken"
			return nil, dbErr
		}
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return user, nil
}

func (r UserRepo) Update(user *models.User) (*models.User, map[string]string) {
	panic("implement me")
}

func (r UserRepo) Delete(id string) error {
	panic("implement me")
}

func (r UserRepo) FindByEmail(email string) (*models.User, error) {
	var user *models.User
	//database.DB.Where("email = ?").First(&user)
	return user, nil
}

func (r *UserRepo) GetUserByEmailAndPassword(u *models.User) (*models.User, map[string]string) {
	var user models.User

	dbErr := map[string]string{}

	err := r.db.Debug().Where("email = ?", u.Email).Take(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		dbErr["no_user"] = "user not found"
		return nil, dbErr
	}
	if err != nil {
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}

	// -> verify the password
	err = utils.VerifyPassword([]byte(user.Password), u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		dbErr["incorrect_password"] = "incorrect password"
		return nil, dbErr
	}
	return &user, nil
}
