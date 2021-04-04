package database

import (
	"go-api/src/app/modules/user/models"
	"go-api/src/app/modules/user/repositories"
	"go-api/src/app/modules/user/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repositories struct {
	User repositories.UserRepositoryInterface
	db   *gorm.DB
}

func NewRepositories(dsn string) (*Repositories, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	return &Repositories{
		User: services.UserServices(db),
		db:   db,
	}, nil
}

func (s *Repositories) Automigrate() error {
	return s.db.AutoMigrate(&models.User{})
}
