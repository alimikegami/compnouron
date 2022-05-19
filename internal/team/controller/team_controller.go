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
	teamUC usecase.TeamUseCase
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

// UpdateTeam godoc
// @Summary      Update team's data
// @Description  Given the request body and the ID path parameters, this endpoint will update the existing team's data
// @Tags         Teams
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Param id path int true "Team ID"
// @Param data body dto.TeamRequest true "Request Body"
// @Success      200  {object}   response.Response{data=string,status=string,message=string}
// @Failure      400  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /teams/{id} [put]
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
		if err.Error() == "action unauthorized" {
			return c.JSON(http.StatusUnauthorized, response.Response{
				Status:  "error",
				Message: err.Error(),
				Data:    nil,
			})
		}
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

// CreateTeam godoc
// @Summary      Create new team
// @Description  Given the request body, the API will create a new team data
// @Tags         Teams
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param data body dto.TeamRequest true "Request Body"
// @Param Authorization header string true "Bearer"
// @Success      201  {object}   response.Response{data=string,status=string,message=string}
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /teams [post]
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

// DeleteTeam godoc
// @Summary      Delete team's data
// @Description  Given the ID path parameters, this endpoint will delete the existing team's data
// @Tags         Teams
// @Produce      json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Param id path int true "Team ID"
// @Success      200  {object}   response.Response{data=string,status=string,message=string}
// @Failure      400  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /teams/{id} [delete]
func (tc *TeamController) DeleteTeam(c echo.Context) error {
	userID, _ := utils.GetUserDetails(c)
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

	err = tc.teamUC.DeleteTeam(uint(teamIDUint), userID)

	if err != nil {
		fmt.Println(err)
		if err.Error() == "action unauthorized" {
			return c.JSON(http.StatusUnauthorized, response.Response{
				Status:  "error",
				Message: err.Error(),
				Data:    nil,
			})
		}
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

// GetTeamsByUserID godoc
// @Summary      Get team's data by user ID
// @Description  Given the user ID as the path parameter, retrieve the team's data that are associated with that particular user
// @Tags         Teams
// @Produce      json
// @Param id path int true "Team ID"
// @Success      200  {object}   response.Response{data=[]dto.BriefTeamResponse,status=string,message=string}
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /teams/users/{id} [get]
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

// GetTeamDetailsByID godoc
// @Summary      Get detailed team's data by team ID
// @Description  Given the team ID, retrieve the detailed team's data that are associated with that particular ID
// @Tags         Teams
// @Produce      json
// @Param id path int true "Team ID"
// @Success      200  {object}   response.Response{data=[]dto.TeamDetailsResponse,status=string,message=string}
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /teams/{id} [get]
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

func CreateNewTeamController(e *echo.Echo, teamUC usecase.TeamUseCase) *TeamController {
	return &TeamController{router: e, teamUC: teamUC}
}
