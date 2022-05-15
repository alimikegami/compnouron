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
		r.GET("", cc.GetCompetitions)
		r.POST("/register", cc.Register, middleware.JWTWithConfig(config))
		r.DELETE("/:id", cc.DeleteCompetition, middleware.JWTWithConfig(config))
		r.PUT("/:id", cc.UpdateCompetition, middleware.JWTWithConfig(config))
		r.PUT("/registrations/:id/accept", cc.AcceptCompetitionRegistration, middleware.JWTWithConfig(config))
		r.PUT("/registrations/:id/reject", cc.RejectCompetitionRegistration, middleware.JWTWithConfig(config))
		r.PUT("/:id/open", cc.OpenCompetitionRegistrationPeriod, middleware.JWTWithConfig(config))
		r.PUT("/:id/close", cc.CloseCompetitionRegistrationPeriod, middleware.JWTWithConfig(config))
		r.GET("/:id/registrations", cc.GetCompetitionRegistration, middleware.JWTWithConfig(config))
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

func (cc *CompetitionController) GetCompetitions(c echo.Context) error {
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
	competitionsResponse, err := cc.CompetitionUC.GetCompetitions(limitInt, offsetInt)
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
