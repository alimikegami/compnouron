package controller

import (
	"fmt"
	"net/http"

	"github.com/alimikegami/compnouron/internal/team/dto"
	"github.com/alimikegami/compnouron/internal/team/usecase"
	"github.com/alimikegami/compnouron/pkg/response"
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
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}

	err := tc.teamUC.CreateTeam(userID, *team)
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

func CreateNewTeamController(e *echo.Echo, teamUC *usecase.TeamUseCase) *TeamController {
	return &TeamController{router: e, teamUC: teamUC}
}
