package controller

import (
	"fmt"
	"net/http"

	"github.com/alimikegami/compnouron/internal/recruitment/dto"
	"github.com/alimikegami/compnouron/internal/recruitment/usecase"
	"github.com/alimikegami/compnouron/pkg/response"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RecruitmentController struct {
	router        *echo.Echo
	recruitmentUC *usecase.RecruitmentUseCase
}

func (rc *RecruitmentController) InitializeRecruitmentRoute(config middleware.JWTConfig) {
	r := rc.router.Group("/recruitments")
	{
		r.POST("", rc.CreateRecruitment, middleware.JWTWithConfig(config))
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
		return c.JSON(http.StatusBadRequest, response.Response{
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

func CreateNewRecruitmentController(e *echo.Echo, recruitmentUC *usecase.RecruitmentUseCase) *RecruitmentController {
	return &RecruitmentController{router: e, recruitmentUC: recruitmentUC}
}
