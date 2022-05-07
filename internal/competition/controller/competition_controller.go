package controller

import (
	"fmt"
	"strconv"

	"github.com/alimikegami/compnouron/internal/competition/dto"
	"github.com/alimikegami/compnouron/internal/competition/usecase"
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
		r.DELETE("/:id", cc.DeleteCompetition, middleware.JWTWithConfig(config))
		r.PUT("/:id", cc.UpdateCompetition, middleware.JWTWithConfig(config))
	}
}

func (cc *CompetitionController) CreateCompetition(c echo.Context) error {
	userID, _ := utils.GetUserDetails(c)
	competition := new(dto.CompetitionRequest)
	if err := c.Bind(competition); err != nil {
		fmt.Println(err)
	}
	err := cc.CompetitionUC.CreateCompetition(*competition, userID)
	if err != nil {
		fmt.Println("error creating competition")
	}
	return nil
}

func (cc *CompetitionController) DeleteCompetition(c echo.Context) error {
	competitionID := c.Param("id")
	userID, _ := utils.GetUserDetails(c)
	competitionIDUint, err := strconv.ParseUint(competitionID, 10, 32)
	if err != nil {
		return err
	}
	cc.CompetitionUC.DeleteCompetition(uint(competitionIDUint), userID)

	return nil
}

func (cc *CompetitionController) UpdateCompetition(c echo.Context) error {
	competitionID := c.Param("id")
	competition := new(dto.CompetitionRequest)
	if err := c.Bind(competition); err != nil {
		fmt.Println(err)
	}
	competitionIDUint, err := strconv.ParseUint(competitionID, 10, 32)
	if err != nil {
		return err
	}
	cc.CompetitionUC.UpdateCompetition(*competition, uint(competitionIDUint))
	return nil
}
