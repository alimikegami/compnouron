package controller

import (
	"fmt"
	"net/http"

	"github.com/alimikegami/compnouron/internal/user/dto"
	"github.com/alimikegami/compnouron/internal/user/entity"
	"github.com/alimikegami/compnouron/internal/user/usecase"
	"github.com/alimikegami/compnouron/pkg/response"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type UserController struct {
	router *echo.Echo
	userUC usecase.UserUseCase
}

func (uc *UserController) InitializeUserRoute(config middleware.JWTConfig) {
	uc.router.POST("/users", uc.CreateUser)
	uc.router.POST("/login", uc.Login)
}

func (uc *UserController) CreateUser(c echo.Context) error {
	u := new(entity.User)
	if err := c.Bind(u); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}
	err := uc.userUC.CreateUser(u)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.JSON(http.StatusCreated, response.Response{
		Status:  "success",
		Message: nil,
		Data:    nil,
	})
}

func (uc *UserController) Login(c echo.Context) error {
	credential := new(dto.Credential)
	if err := c.Bind(credential); err != nil {
		fmt.Println("error: error on binding request body")
	}
	token, err := uc.userUC.Login(credential)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, response.Response{
		Status:  "success",
		Message: nil,
		Data: dto.TokenResponse{
			Token:     token,
			TokenType: "JWT",
		},
	})
}

func CreateNewUserController(e *echo.Echo, userUC usecase.UserUseCase) *UserController {
	return &UserController{router: e, userUC: userUC}
}
