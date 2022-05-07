package controller

import (
	"fmt"

	"github.com/alimikegami/compnouron/internal/team/dto"
	"github.com/alimikegami/compnouron/internal/team/usecase"
	"github.com/alimikegami/compnouron/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TeamController struct {
	router *echo.Echo
	teamUC *usecase.TeamUseCase
}

func (tc *TeamController) InitializeTeamRoute(config middleware.JWTConfig) {
	r := tc.router.Group("/teams")
	{
		r.POST("", tc.CreateTeam, middleware.JWTWithConfig(config))
	}
}

func (tc *TeamController) CreateTeam(c echo.Context) error {
	userID, _ := utils.GetUserDetails(c)
	team := new(dto.TeamRequest)
	if err := c.Bind(team); err != nil {
		fmt.Println(err)
	}

	tc.teamUC.CreateTeam(userID, *team)

	return nil
}

func CreateNewTeamController(e *echo.Echo, teamUC *usecase.TeamUseCase) *TeamController {
	return &TeamController{router: e, teamUC: teamUC}
}
