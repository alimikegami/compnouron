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
		r.GET("/:id/applications", rc.GetRecruitmentDetailsByID, middleware.JWTWithConfig(config))
		r.GET("/user", rc.GetRecruitmentByUserID, middleware.JWTWithConfig(config))
		r.PUT("/applications/:id/accept", rc.AcceptRecruitmentApplication, middleware.JWTWithConfig(config))
		r.PUT("/applications/:id/reject", rc.RejectRecruitmentApplication, middleware.JWTWithConfig(config))
		r.DELETE("/:id", rc.DeleteRecruitmentByID, middleware.JWTWithConfig(config))
		r.PUT("/:id/open", rc.OpenRecruitmentApplicationPeriod, middleware.JWTWithConfig(config))
		r.PUT("/:id/close", rc.CloseRecruitmentApplicationPeriod, middleware.JWTWithConfig(config))

	}
}

func (rc *RecruitmentController) CreateRecruitment(c echo.Context) error {
	recruitment := new(dto.RecruitmentRequest)
	if err := c.Bind(recruitment); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}
	err := rc.recruitmentUC.CreateRecruitment(*recruitment)
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

func (rc *RecruitmentController) UpdateRecruitment(c echo.Context) error {
	recruitmentID := c.Param("id")
	recruitmentIDUint, err := strconv.ParseUint(recruitmentID, 10, 32)
	if err != nil {
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

	err = rc.recruitmentUC.UpdateRecruitment(*recruitment, uint(recruitmentIDUint))
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

func (rc *RecruitmentController) GetRecruitments(c echo.Context) error {
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

	result, err := rc.recruitmentUC.GetRecruitments(limitInt, offsetInt)
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

func (rc *RecruitmentController) GetRecruitmentDetailsByID(c echo.Context) error {
	recruitmentID := c.Param("id")
	recruitmentIDUint, err := strconv.ParseUint(recruitmentID, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}

	result, err := rc.recruitmentUC.GetRecruitmentDetailsByID(uint(recruitmentIDUint))
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

func (rc *RecruitmentController) GetRecruitmentByUserID(c echo.Context) error {
	userID, _ := utils.GetUserDetails(c)
	result, err := rc.recruitmentUC.GetRecruitmentByUserID(userID)
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

func (rc *RecruitmentController) RejectRecruitmentApplication(c echo.Context) error {
	recruitmentApplicationID := c.Param("id")
	recruitmentApplicationIDUint, err := strconv.ParseUint(recruitmentApplicationID, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = rc.recruitmentUC.RejectRecruitmentApplication(uint(recruitmentApplicationIDUint))
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

func (rc *RecruitmentController) AcceptRecruitmentApplication(c echo.Context) error {
	recruitmentApplicationID := c.Param("id")
	recruitmentApplicationIDUint, err := strconv.ParseUint(recruitmentApplicationID, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = rc.recruitmentUC.AcceptRecruitmentApplication(uint(recruitmentApplicationIDUint))
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
	recruitment := new(dto.RecruitmentRequest)
	if err := c.Bind(recruitment); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}

	res, err := rc.recruitmentUC.GetRecruitmentByID(uint(recruitmentIDUint))
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

func (rc *RecruitmentController) DeleteRecruitmentByID(c echo.Context) error {
	recruitmentID := c.Param("id")
	recruitmentIDUint, err := strconv.ParseUint(recruitmentID, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = rc.recruitmentUC.DeleteRecruitmentByID(uint(recruitmentIDUint))
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

func (rc *RecruitmentController) OpenRecruitmentApplicationPeriod(c echo.Context) error {
	competition := c.Param("id")
	competitionUint, err := strconv.ParseUint(competition, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = rc.recruitmentUC.OpenRecruitmentApplicationPeriod(uint(competitionUint))
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

func (rc *RecruitmentController) CloseRecruitmentApplicationPeriod(c echo.Context) error {
	competition := c.Param("id")
	competitionUint, err := strconv.ParseUint(competition, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = rc.recruitmentUC.CloseRecruitmentApplicationPeriod(uint(competitionUint))
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

func CreateNewRecruitmentController(e *echo.Echo, recruitmentUC usecase.RecruitmentUseCase) *RecruitmentController {
	return &RecruitmentController{router: e, recruitmentUC: recruitmentUC}
}
