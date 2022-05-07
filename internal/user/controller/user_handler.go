package controller

import (
	"fmt"
	"net/http"

	"github.com/alimikegami/compnouron/internal/user/dto"
	"github.com/alimikegami/compnouron/internal/user/entity"
	"github.com/alimikegami/compnouron/internal/user/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type UserController struct {
	router *echo.Echo
	userUC *usecase.UserUseCase
}

func (uc *UserController) InitializeUserRoute(config middleware.JWTConfig) {
	uc.router.POST("/users", uc.CreateUser)
	uc.router.POST("/login", uc.Login)
}

func (uc *UserController) CreateUser(c echo.Context) error {
	u := new(entity.User)
	if err := c.Bind(u); err != nil {
		fmt.Println("Error: Error on binding request body")
	}
	err := uc.userUC.CreateUser(u)
	return err
}

func (uc *UserController) Login(c echo.Context) error {
	credential := new(dto.Credential)
	if err := c.Bind(credential); err != nil {
		fmt.Println("error: error on binding request body")
	}
	token, err := uc.userUC.Login(credential)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(token)
	return c.JSON(http.StatusOK, &dto.TokenResponse{
		Token:     token,
		TokenType: "JWT",
	})
}

func CreateNewUserController(e *echo.Echo, userUC *usecase.UserUseCase) *UserController {
	return &UserController{router: e, userUC: userUC}
}
