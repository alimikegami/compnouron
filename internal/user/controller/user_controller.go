package controller

import (
	"fmt"
	"net/http"

	"github.com/alimikegami/compnouron/internal/user/dto"
	"github.com/alimikegami/compnouron/internal/user/usecase"
	"github.com/alimikegami/compnouron/pkg/response"
	"github.com/alimikegami/compnouron/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type UserController struct {
	router *echo.Echo
	userUC usecase.UserUseCase
}

func (uc *UserController) InitializeUserRoute(config middleware.JWTConfig) {
	uc.router.POST("/users", uc.CreateUser)
	uc.router.POST("/users/login", uc.Login)
	uc.router.GET("/users/competitions/registrations", uc.GetCompetitionRegistrationHistory, middleware.JWTWithConfig(config))
	uc.router.GET("/users/recruitments/applications", uc.GetRecruitmentApplicationHistory, middleware.JWTWithConfig(config))
}

// CreateUser godoc
// @Summary      Create new user account
// @Description  Given the request body, create a new user record in the database
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param data body dto.UserRegistrationRequest true "Request Body"
// @Success      200  {object}   response.Response{data=string,status=string,message=string}
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /users [post]
func (uc *UserController) CreateUser(c echo.Context) error {
	u := new(dto.UserRegistrationRequest)
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

// Login godoc
// @Summary      Login
// @Description  Given the credentials, authenticate the credentials and returns the JWT token if the credentials matched the record in the database
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param data body dto.Credential true "Request Body"
// @Success      200  {object}   response.Response{data=dto.TokenResponse,status=string,message=string}
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /users/login [post]
func (uc *UserController) Login(c echo.Context) error {
	credential := new(dto.Credential)
	if err := c.Bind(credential); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}
	token, err := uc.userUC.Login(credential)
	if err != nil {
		var statusCode int
		fmt.Println(err)
		if err.Error() == "credentials dont match" {
			statusCode = 403
		} else {
			statusCode = 500
		}
		return c.JSON(statusCode, response.Response{
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

func (uc *UserController) GetCompetitionRegistrationHistory(c echo.Context) error {
	userID, _ := utils.GetUserDetails(c)
	result, err := uc.userUC.GetCompetitionRegistrationHistory(userID)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, response.Response{
		Status:  "success",
		Message: nil,
		Data:    result,
	})
}

func (uc *UserController) GetRecruitmentApplicationHistory(c echo.Context) error {
	userID, _ := utils.GetUserDetails(c)
	result, err := uc.userUC.GetRecruitmentApplicationHistory(userID)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, response.Response{
		Status:  "success",
		Message: nil,
		Data:    result,
	})
}

func CreateNewUserController(e *echo.Echo, userUC usecase.UserUseCase) *UserController {
	return &UserController{router: e, userUC: userUC}
}
