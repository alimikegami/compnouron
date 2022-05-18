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
	CompetitionUC usecase.CompetitionUseCase
}

func CreateNewCompetitionController(e *echo.Echo, CompetitionUC usecase.CompetitionUseCase) *CompetitionController {
	return &CompetitionController{router: e, CompetitionUC: CompetitionUC}
}

func (cc *CompetitionController) InitializeCompetitionRoute(config middleware.JWTConfig) {
	r := cc.router.Group("/competitions")
	{
		r.POST("", cc.CreateCompetition, middleware.JWTWithConfig(config))
		r.GET("", cc.GetCompetitions)
		r.POST("/registrations", cc.Register, middleware.JWTWithConfig(config))
		r.DELETE("/:id", cc.DeleteCompetition, middleware.JWTWithConfig(config))
		r.PUT("/:id", cc.UpdateCompetition, middleware.JWTWithConfig(config))
		r.PUT("/registrations/:id/accept", cc.AcceptCompetitionRegistration, middleware.JWTWithConfig(config))
		r.PUT("/registrations/:id/reject", cc.RejectCompetitionRegistration, middleware.JWTWithConfig(config))
		r.PUT("/:id/open", cc.OpenCompetitionRegistrationPeriod, middleware.JWTWithConfig(config))
		r.PUT("/:id/close", cc.CloseCompetitionRegistrationPeriod, middleware.JWTWithConfig(config))
		r.GET("/:id", cc.GetCompetitionRegistration, middleware.JWTWithConfig(config))
		r.GET("/:id/registrations", cc.GetCompetitionRegistration, middleware.JWTWithConfig(config))
	}
}

// CreateCompetition godoc
// @Summary      Create new competition
// @Description  Given the request body, create a competition record in the database
// @Tags         Competitions
// @Accept       json
// @Produce      json
// @Param data body dto.CompetitionRequest true "Request Body"
// @Success      200  {object}   response.Response{data=string,status=string,message=string}
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /competitions [post]
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

// DeleteCompetition godoc
// @Summary      Delete competition's data
// @Description  Given the ID path parameters, this endpoint will delete the existing competition's data
// @Tags         Competitions
// @Produce      json
// @Param id path int true "Competition ID"
// @Success      200  {object}   response.Response{data=string,status=string,message=string}
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /competitions/{id} [delete]
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

// UpdateCompetition godoc
// @Summary      Update competition's data
// @Description  Given the request body and the ID path parameters, this endpoint will update the existing competition's data
// @Tags         Competitions
// @Accept       json
// @Produce      json
// @Param id path int true "Competition ID"
// @Success      200  {object}   response.Response{data=string,status=string,message=string}
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /competitions/{id} [put]
func (cc *CompetitionController) UpdateCompetition(c echo.Context) error {
	competitionID := c.Param("id")
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
	competitionIDUint, err := strconv.ParseUint(competitionID, 10, 32)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}
	err = cc.CompetitionUC.UpdateCompetition(*competition, uint(competitionIDUint), userID)
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

// GetCompetitions godoc
// @Summary      Get competitions data
// @Description  This endpoint will return the competitions data with pagination implemented and also with keyword searching capability
// @Tags         Competitions
// @Produce      json
// @Param        limit     query      int     true  "rows retrieved limit"
// @Param        offset    query      int     true  "skipped rows"
// @Param        keyword   query      string  false  "competition name keyword"
// @Success      200  {object}   response.Response{data=[]dto.CompetitionResponse,status=string,message=string}
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /competitions [get]
func (cc *CompetitionController) GetCompetitions(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}

	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}
	var competitionsResponse []dto.CompetitionResponse
	if keyword != "" {
		competitionsResponse, err = cc.CompetitionUC.SearchCompetition(limitInt, offsetInt, keyword)

	} else {
		competitionsResponse, err = cc.CompetitionUC.GetCompetitions(limitInt, offsetInt)

	}
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
		Data:    competitionsResponse,
	})
}

// Register godoc
// @Summary      Create new competition registrations
// @Description  Given the request body, create a competition registration record in the database
// @Tags         Competitions
// @Accept       json
// @Produce      json
// @Success      200  {object}   response.Response{data=string,status=string,message=string}
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /competitions/registrations [post]
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

	return c.JSON(http.StatusCreated, response.Response{
		Status:  "success",
		Message: nil,
		Data:    nil,
	})
}

// RejectCompetitionRegistration godoc
// @Summary      Reject competition application
// @Description  Given the competition registration ID path parameters, this endpoint will reject the competition registration
// @Tags         Competitions
// @Produce      json
// @Param id path int true "Competition ID"
// @Success      200  {object}   response.Response{data=string,status=string,message=string}
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /competitions/registrations/{id}/reject [put]
func (cc *CompetitionController) RejectCompetitionRegistration(c echo.Context) error {
	competitionRegistrationID := c.Param("id")
	competitionRegistrationIDUint, err := strconv.ParseUint(competitionRegistrationID, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = cc.CompetitionUC.RejectCompetitionRegistration(uint(competitionRegistrationIDUint))
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
		Data:    nil,
	})
}

// AcceptCompetitionRegistration godoc
// @Summary      Accept competition application
// @Description  Given the competition registration ID path parameters, this endpoint will accept the competition registration
// @Tags         Competitions
// @Produce      json
// @Param id path int true "Competition ID"
// @Success      200  {object}   response.Response{data=string,status=string,message=string}
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /competitions/registrations/{id}/accept [put]
func (cc *CompetitionController) AcceptCompetitionRegistration(c echo.Context) error {
	competitionRegistrationID := c.Param("id")
	competitionRegistrationIDUint, err := strconv.ParseUint(competitionRegistrationID, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = cc.CompetitionUC.AcceptCompetitionRegistration(uint(competitionRegistrationIDUint))
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
		Data:    nil,
	})
}

// CloseCompetitionRegistrationPeriod godoc
// @Summary      Close competition registration period
// @Description  Given the competition ID path parameters, this endpoint will close the competition registration period
// @Tags         Competitions
// @Produce      json
// @Param id path int true "Competition ID"
// @Success      200  {object}   response.Response{data=string,status=string,message=string}
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /competitions/{id}/close [put]
func (cc *CompetitionController) CloseCompetitionRegistrationPeriod(c echo.Context) error {
	competition := c.Param("id")
	competitionUint, err := strconv.ParseUint(competition, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = cc.CompetitionUC.CloseCompetitionRegistrationPeriod(uint(competitionUint))
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
		Data:    nil,
	})
}

// OpenCompetitionRegistrationPeriod godoc
// @Summary      Open competition registration period
// @Description  Given the competition ID path parameters, this endpoint will open the competition registration period
// @Tags         Competitions
// @Produce      json
// @Param id path int true "Competition ID"
// @Success      200  {object}   response.Response{data=string,status=string,message=string}
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /competitions/{id}/open [put]
func (cc *CompetitionController) OpenCompetitionRegistrationPeriod(c echo.Context) error {
	competition := c.Param("id")
	competitionUint, err := strconv.ParseUint(competition, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = cc.CompetitionUC.OpenCompetitionRegistrationPeriod(uint(competitionUint))
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
		Data:    nil,
	})
}

// GetCompetitionRegistration godoc
// @Summary      Get competition registration data
// @Description  Given the ID path parameters and the status query parameteres, this endpoint will retrieve the competition registration data of a particular ID and accepted status if the query parameters are given
// @Tags         Competitions
// @Produce      json
// @Param        status    query      int     true  "filter to get accepted registrations record"
// @Param id path int true "Competition ID"
// @Success      200  {object}   response.Response{data=interface{},status=string,message=string}
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /competitions/{id}/registrations [get]
func (cc *CompetitionController) GetCompetitionRegistration(c echo.Context) error {
	status := c.QueryParam("status")
	competitionID := c.Param("id")
	competitionIDUint, err := strconv.ParseUint(competitionID, 10, 32)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}
	var res interface{}
	if status == "accepted" {
		res, err = cc.CompetitionUC.GetAcceptedCompetitionParticipants(uint(competitionIDUint))
	} else {
		res, err = cc.CompetitionUC.GetCompetitionRegistration(uint(competitionIDUint))
	}
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
		Data:    res,
	})
}
