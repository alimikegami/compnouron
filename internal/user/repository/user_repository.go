package repository

import (
	"github.com/alimikegami/compnouron/internal/user/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user entity.User) (uint, error)
	GetUserByEmail(email string) *entity.User
	AddUserSkills(skill []entity.Skill) error
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func (ur *userRepositoryImpl) CreateUser(user entity.User) (uint, error) {
	result := ur.db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	return user.ID, nil
}

func (ur *userRepositoryImpl) AddUserSkills(skill []entity.Skill) error {
	result := ur.db.Create(&skill)
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
