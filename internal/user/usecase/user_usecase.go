package usecase

import (
	"fmt"

	"github.com/alimikegami/compnouron/internal/user/dto"
	"github.com/alimikegami/compnouron/internal/user/entity"
	"github.com/alimikegami/compnouron/internal/user/repository"
	"github.com/alimikegami/compnouron/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	CreateUser(user *entity.User) error
	Login(credential *dto.Credential) (string, error)
}

type UserUseCaseImpl struct {
	ur repository.UserRepository
}

func CreateNewUserUseCase(ur repository.UserRepository) UserUseCase {
	return &UserUseCaseImpl{ur: ur}
}

func (us *UserUseCaseImpl) CreateUser(user *entity.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	err = us.ur.CreateUser(user)

	return err
}

func (us *UserUseCaseImpl) Login(credential *dto.Credential) (string, error) {
	user := us.ur.GetUserByEmail(credential.Email)
	if user == nil {
		return "", fmt.Errorf("user not found")
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credential.Password))
	if err != nil {
		return "", fmt.Errorf("credentials dont match")
	}
	token, err := utils.CreateJWTToken(user.ID, user.Email)
	if err != nil {
		fmt.Println(err)
	}

	return token, nil
}
