package controller

import (
	"fmt"
	"net/http"
	"strconv"

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
		r.PUT("/:id", tc.UpdateTeam, middleware.JWTWithConfig(config))
		r.DELETE("/:id", tc.DeleteTeam, middleware.JWTWithConfig(config))
		r.GET("/users/:id", tc.GetTeamsByUserID)
		r.GET("/:id", tc.GetTeamDetailsByID)
	}
}

func (tc *TeamController) UpdateTeam(c echo.Context) error {
	teamID := c.Param("id")
	teamIDUint, err := strconv.ParseUint(teamID, 10, 32)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}
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

	err = tc.teamUC.UpdateTeam(userID, *team, uint(teamIDUint))
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

func (tc *TeamController) DeleteTeam(c echo.Context) error {
	teamID := c.Param("id")
	// userID, _ := utils.GetUserDetails(c)
	teamIDUint, err := strconv.ParseUint(teamID, 10, 32)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = tc.teamUC.DeleteTeam(uint(teamIDUint))

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

func (tc *TeamController) GetTeamsByUserID(c echo.Context) error {
	userID := c.Param("id")
	userIDUint, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}

	result, err := tc.teamUC.GetTeamsByUserID(uint(userIDUint))
	if err != nil {
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

func (tc *TeamController) GetTeamDetailsByID(c echo.Context) error {
	teamID := c.Param("id")
	teamIDUint, err := strconv.ParseUint(teamID, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}

	result, err := tc.teamUC.GetTeamDetailsByID(uint(teamIDUint))
	if err != nil {
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

func CreateNewTeamController(e *echo.Echo, teamUC *usecase.TeamUseCase) *TeamController {
	return &TeamController{router: e, teamUC: teamUC}
}
