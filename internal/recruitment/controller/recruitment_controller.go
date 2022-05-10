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
		r.PUT("/:id", rc.UpdateRecruitment, middleware.JWTWithConfig(config))
		r.POST("/applications", rc.CreateRecruitmentApplication, middleware.JWTWithConfig(config))
		r.GET("/:id/applications", rc.GetRecruitmentDetailsByID, middleware.JWTWithConfig(config))

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

func CreateNewRecruitmentController(e *echo.Echo, recruitmentUC usecase.RecruitmentUseCase) *RecruitmentController {
	return &RecruitmentController{router: e, recruitmentUC: recruitmentUC}
}
