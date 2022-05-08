package repository

import (
	"github.com/alimikegami/compnouron/internal/user/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *entity.User) error
	GetUserByEmail(email string) *entity.User
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func (ur *userRepositoryImpl) CreateUser(user *entity.User) error {
	result := ur.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ur *userRepositoryImpl) GetUserByEmail(email string) *entity.User {
	var user entity.User
	ur.db.First(&user, "email = ?", email)

	return &user
}

func CreateNewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}
