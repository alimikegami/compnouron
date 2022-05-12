package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alimikegami/compnouron/internal/competition/dto"
	"github.com/alimikegami/compnouron/internal/competition/usecase"
	"github.com/alimikegami/compnouron/pkg/response"
	"github.com/alimikegami/compnouron/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CompetitionController struct {
	router        *echo.Echo
	CompetitionUC *usecase.CompetitionUseCase
}

func CreateNewCompetitionController(e *echo.Echo, CompetitionUC *usecase.CompetitionUseCase) *CompetitionController {
	return &CompetitionController{router: e, CompetitionUC: CompetitionUC}
}

func (cc *CompetitionController) InitializeCompetitionRoute(config middleware.JWTConfig) {
	r := cc.router.Group("/competitions")
	{
		r.POST("", cc.CreateCompetition, middleware.JWTWithConfig(config))
		r.POST("/register", cc.Register, middleware.JWTWithConfig(config))
		r.DELETE("/:id", cc.DeleteCompetition, middleware.JWTWithConfig(config))
		r.PUT("/:id", cc.UpdateCompetition, middleware.JWTWithConfig(config))
	}
}

func (cc *CompetitionController) CreateCompetition(c echo.Context) error {
	userID, _ := utils.GetUserDetails(c)
	competition := new(dto.CompetitionRequest)
	if err := c.Bind(competition); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}
	err := cc.CompetitionUC.CreateCompetition(*competition, userID)
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

func (cc *CompetitionController) DeleteCompetition(c echo.Context) error {
	competitionID := c.Param("id")
	userID, _ := utils.GetUserDetails(c)
	competitionIDUint, err := strconv.ParseUint(competitionID, 10, 32)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}
	err = cc.CompetitionUC.DeleteCompetition(uint(competitionIDUint), userID)

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
		Data:    nil,
	})
}

func (cc *CompetitionController) UpdateCompetition(c echo.Context) error {
	competitionID := c.Param("id")
	competition := new(dto.CompetitionRequest)
	if err := c.Bind(competition); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}
	competitionIDUint, err := strconv.ParseUint(competitionID, 10, 32)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}
	err = cc.CompetitionUC.UpdateCompetition(*competition, uint(competitionIDUint))
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
		Data:    nil,
	})
}

func (cc *CompetitionController) Register(c echo.Context) error {
	competitionRegistration := new(dto.CompetitionRegistrationRequest)
	if err := c.Bind(competitionRegistration); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}
	err := cc.CompetitionUC.Register(*competitionRegistration)

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
		Data:    nil,
	})
}
