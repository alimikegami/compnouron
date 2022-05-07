package repository

import (
	"fmt"

	"github.com/alimikegami/compnouron/internal/user/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func CreateNewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) CreateUser(user *entity.User) error {
	result := ur.db.Create(&user)
	fmt.Println(result)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ur *UserRepository) GetUserByEmail(email string) *entity.User {
	var user entity.User
	ur.db.First(&user, "email = ?", email)

	return &user
}
