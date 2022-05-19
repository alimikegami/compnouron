package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alimikegami/compnouron/internal/recruitment/dto"
	"github.com/alimikegami/compnouron/internal/recruitment/usecase"
	"github.com/alimikegami/compnouron/pkg/response"
	"github.com/alimikegami/compnouron/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RecruitmentController struct {
	router        *echo.Echo
	recruitmentUC usecase.RecruitmentUseCase
}

func (rc *RecruitmentController) InitializeRecruitmentRoute(config middleware.JWTConfig) {
	r := rc.router.Group("/recruitments")
	{
		r.POST("", rc.CreateRecruitment, middleware.JWTWithConfig(config))
		r.GET("", rc.GetRecruitments)
		r.PUT("/:id", rc.UpdateRecruitment, middleware.JWTWithConfig(config))
		r.POST("/applications", rc.CreateRecruitmentApplication, middleware.JWTWithConfig(config))
		r.GET("/:id/details", rc.GetRecruitmentByID)
		r.GET("/:id/details", rc.GetRecruitmentDetailsByID, middleware.JWTWithConfig(config))
		r.GET("/teams/:id", rc.GetRecruitmentByTeamID, middleware.JWTWithConfig(config))
		r.PUT("/applications/:id/accept", rc.AcceptRecruitmentApplication, middleware.JWTWithConfig(config))
		r.PUT("/applications/:id/reject", rc.RejectRecruitmentApplication, middleware.JWTWithConfig(config))
		r.DELETE("/:id", rc.DeleteRecruitmentByID, middleware.JWTWithConfig(config))
		r.PUT("/:id/open", rc.OpenRecruitmentApplicationPeriod, middleware.JWTWithConfig(config))
		r.PUT("/:id/close", rc.CloseRecruitmentApplicationPeriod, middleware.JWTWithConfig(config))
	}
}

// CreateRecruitment godoc
// @Summary      Create new recruitment
// @Description  Given the request body, create a recruiment record in the database
// @Tags         Recruitments
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Param data body dto.RecruitmentRequest true "Request Body"
// @Success      200  {object}  response.Response{data=string,status=string,message=string}
// @Failure      400  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /recruitments [post]
func (rc *RecruitmentController) CreateRecruitment(c echo.Context) error {
	userID, _ := utils.GetUserDetails(c)
	recruitment := new(dto.RecruitmentRequest)
	if err := c.Bind(recruitment); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}
	err := rc.recruitmentUC.CreateRecruitment(*recruitment, userID)
	if err != nil {
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

	return c.JSON(http.StatusCreated, response.Response{
		Status:  "success",
		Message: nil,
		Data:    nil,
	})
}

// UpdateRecruitment godoc
// @Summary      Update recruitment's data
// @Description  Given the request body and the ID path parameters, this endpoint will update the existing recruitment's data
// @Tags         Recruitments
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Param id path int true "Recruitment ID"
// @Param data body dto.RecruitmentRequest true "Request Body"
// @Success      200  {object}   response.Response{data=string,status=string,message=string}
// @Failure      400  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /recruitments/{id} [put]
func (rc *RecruitmentController) UpdateRecruitment(c echo.Context) error {
	userID, _ := utils.GetUserDetails(c)
	recruitmentID := c.Param("id")
	recruitmentIDUint, err := strconv.ParseUint(recruitmentID, 10, 32)
	if err != nil {
		if err.Error() == "action unauthorized" {
			return c.JSON(http.StatusUnauthorized, response.Response{
				Status:  "error",
				Message: err.Error(),
				Data:    nil,
			})
		}
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}
	recruitment := new(dto.RecruitmentRequest)
	if err := c.Bind(recruitment); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = rc.recruitmentUC.UpdateRecruitment(*recruitment, uint(recruitmentIDUint), userID)
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

// CreateRecruitmentApplication godoc
// @Summary      Create new recruitment appliation
// @Description  Given the request body, create a recruiment application record in the database
// @Tags         Recruitments
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Param data body dto.RecruitmentApplicationRequest true "Request Body"
// @Success      200  {object}   response.Response{data=string,status=string,message=string}
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /recruitments/applications [post]
func (rc *RecruitmentController) CreateRecruitmentApplication(c echo.Context) error {
	userID, _ := utils.GetUserDetails(c)
	recruitmentApplication := new(dto.RecruitmentApplicationRequest)
	if err := c.Bind(recruitmentApplication); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}

	err := rc.recruitmentUC.CreateRecruitmentApplication(*recruitmentApplication, userID)
	if err != nil {
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

// GetRecruitments godoc
// @Summary      Get recruitments's data
// @Description  This endpoint will return the recruitments data with pagination implemented and also with keyword searching capability
// @Tags         Recruitments
// @Produce      json
// @Param        limit     query      int     true  "rows retrieved limit"
// @Param        offset    query      int     true  "skipped rows"
// @Param        keyword   query      string  false  "recruitment role keyword"
// @Success      200  {object}   response.Response{data=dto.RecruitmentsResponse,status=string,message=string}
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /recruitments [get]
func (rc *RecruitmentController) GetRecruitments(c echo.Context) error {
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

	var result []dto.BriefRecruitmentResponse

	if keyword != "" {
		result, err = rc.recruitmentUC.SearchRecruitment(limitInt, offsetInt, keyword)
	} else {
		result, err = rc.recruitmentUC.GetRecruitments(limitInt, offsetInt)
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
		Data:    result,
	})
}

// GetRecruitmentByID godoc
// @Summary      Get recruitment data
// @Description  Given the recruitment ID on path parameter, this endpoint will return the data associated with that particular recruitment (briefly)
// @Tags         Recruitments
// @Produce      json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Param id path int true "Recruitment ID"
// @Success      200  {object}  response.Response{data=dto.RecruitmentDetailsResponse,status=string,message=string}
// @Failure      400  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /recruitments/{id} [get]
func (rc *RecruitmentController) GetRecruitmentByID(c echo.Context) error {
	recruitmentID := c.Param("id")
	recruitmentIDUint, err := strconv.ParseUint(recruitmentID, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}

	result, err := rc.recruitmentUC.GetRecruitmentByID(uint(recruitmentIDUint))
	if err != nil {
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
		Data:    result,
	})
}

// GetRecruitmentDetailsByID godoc
// @Summary      Get detailed recruitment's data
// @Description  Given the recruitment ID on path parameter, this endpoint will return the detailed data associated with that particular recruitment
// @Tags         Recruitments
// @Produce      json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Param id path int true "Recruitment ID"
// @Success      200  {object}   response.Response{data=dto.RecruitmentDetailsResponse,status=string,message=string}
// @Failure      400  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /recruitments/{id}/details [get]
func (rc *RecruitmentController) GetRecruitmentDetailsByID(c echo.Context) error {
	userID, _ := utils.GetUserDetails(c)

	recruitmentID := c.Param("id")
	recruitmentIDUint, err := strconv.ParseUint(recruitmentID, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}

	result, err := rc.recruitmentUC.GetRecruitmentDetailsByID(uint(recruitmentIDUint), userID)
	if err != nil {
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
		Data:    result,
	})
}

// GetRecruitmentByTeamID godoc
// @Summary      Get recruitment's data of a particular team
// @Description  Given the team ID on path parameter, this endpoint will return the recruitment data associated with that team
// @Tags         Recruitments
// @Produce      json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Param id path int true "Team ID"
// @Success      200  {object}   response.Response{data=dto.RecruitmentsResponse,status=string,message=string}
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /recruitments/teams/{id} [get]
func (rc *RecruitmentController) GetRecruitmentByTeamID(c echo.Context) error {
	teamID := c.Param("id")
	teamIDUint, err := strconv.ParseUint(teamID, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}
	result, err := rc.recruitmentUC.GetRecruitmentByTeamID(uint(teamIDUint))
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

// RejectRecruitmentApplication godoc
// @Summary      Reject recruitment application
// @Description  Given the recruitment application ID path parameters, this endpoint will reject the recruitment application
// @Tags         Recruitments
// @Produce      json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Param id path int true "Recruitment ID"
// @Success      200  {object}   response.Response{data=string,status=string,message=string}
// @Failure      400  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /recruitments/applications/{id}/reject [put]
func (rc *RecruitmentController) RejectRecruitmentApplication(c echo.Context) error {
	userID, _ := utils.GetUserDetails(c)
	recruitmentApplicationID := c.Param("id")
	recruitmentApplicationIDUint, err := strconv.ParseUint(recruitmentApplicationID, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = rc.recruitmentUC.RejectRecruitmentApplication(uint(recruitmentApplicationIDUint), userID)
	if err != nil {
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

// AcceptRecruitmentApplication godoc
// @Summary      Accept recruitment application
// @Description  Given the recruitment application ID path parameters, this endpoint will accept the recruitment application
// @Tags         Recruitments
// @Produce      json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Param id path int true "Recruitment ID"
// @Success      200  {object}   response.Response{data=string,status=string,message=string}
// @Failure      400  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /recruitments/applications/{id}/accept [put]
func (rc *RecruitmentController) AcceptRecruitmentApplication(c echo.Context) error {
	userID, _ := utils.GetUserDetails(c)
	recruitmentApplicationID := c.Param("id")
	recruitmentApplicationIDUint, err := strconv.ParseUint(recruitmentApplicationID, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = rc.recruitmentUC.AcceptRecruitmentApplication(uint(recruitmentApplicationIDUint), userID)
	if err != nil {
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

// DeleteRecruitmentByID godoc
// @Summary      Delete recruitment's data
// @Description  Given the ID path parameters, this endpoint will delete the existing recruitment's data
// @Tags         Recruitments
// @Produce      json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Param id path int true "Recruitment ID"
// @Success      200  {object}   response.Response{data=string,status=string,message=string}
// @Failure      400  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /recruitments/{id} [delete]
func (rc *RecruitmentController) DeleteRecruitmentByID(c echo.Context) error {
	userID, _ := utils.GetUserDetails(c)
	recruitmentID := c.Param("id")
	recruitmentIDUint, err := strconv.ParseUint(recruitmentID, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = rc.recruitmentUC.DeleteRecruitmentByID(uint(recruitmentIDUint), userID)
	if err != nil {
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

// OpenRecruitmentApplicationPeriod godoc
// @Summary      Open recruitment application period
// @Description  Given the recruitment ID path parameters, this endpoint will open the recruitment application period
// @Tags         Recruitments
// @Produce      json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Param id path int true "Recruitment ID"
// @Success      200  {object}   response.Response{data=string,status=string,message=string}
// @Failure      400  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /recruitments/{id}/open [put]
func (rc *RecruitmentController) OpenRecruitmentApplicationPeriod(c echo.Context) error {
	userID, _ := utils.GetUserDetails(c)
	competition := c.Param("id")
	competitionUint, err := strconv.ParseUint(competition, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = rc.recruitmentUC.OpenRecruitmentApplicationPeriod(uint(competitionUint), userID)
	if err != nil {
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

// CloseRecruitmentApplicationPeriod godoc
// @Summary      Close recruitment application period
// @Description  Given the recruitment ID path parameters, this endpoint will close the recruitment application period
// @Tags         Recruitments
// @Produce      json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Param id path int true "Recruitment ID"
// @Success      200  {object}   response.Response{data=string,status=string,message=string}
// @Failure      400  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /recruitments/{id}/close [put]
func (rc *RecruitmentController) CloseRecruitmentApplicationPeriod(c echo.Context) error {
	userID, _ := utils.GetUserDetails(c)
	competition := c.Param("id")
	competitionUint, err := strconv.ParseUint(competition, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = rc.recruitmentUC.CloseRecruitmentApplicationPeriod(uint(competitionUint), userID)
	if err != nil {
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

func CreateNewRecruitmentController(e *echo.Echo, recruitmentUC usecase.RecruitmentUseCase) *RecruitmentController {
	return &RecruitmentController{router: e, recruitmentUC: recruitmentUC}
}
