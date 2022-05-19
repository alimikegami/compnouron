package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	mocks "github.com/alimikegami/compnouron/internal/mocks/recruitment/usecase"
	"github.com/alimikegami/compnouron/internal/recruitment/dto"
	"github.com/alimikegami/compnouron/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateRecruitmentApplication(t *testing.T) {
	mockUseCase := mocks.NewRecruitmentUseCase(t)
	t.Run("success", func(t *testing.T) {
		mockUseCase.On("CreateRecruitmentApplication", dto.RecruitmentApplicationRequest{
			RecruitmentID: 1,
		}, uint(1)).Return(nil).Once()

		// construct request body
		reqBody := dto.RecruitmentApplicationRequest{
			RecruitmentID: 1,
		}

		jsonReqBody, err := json.Marshal(&reqBody)
		assert.NoError(t, err, "No marshaling error")

		// setup the endpoint
		req, err := http.NewRequest(http.MethodPost, "/recruitments/applications", bytes.NewBuffer(jsonReqBody))
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")

		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		token := utils.CreateJWTToken(1, "gmail@gmail.com")
		c.Set("user", token)
		// setup controller/handler
		testRecruitmentController := RecruitmentController{
			router:        e,
			recruitmentUC: mockUseCase,
		}

		// get the response
		testRecruitmentController.CreateRecruitmentApplication(c)
		assert.Equal(t, http.StatusCreated, rec.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("foreign-key-error", func(t *testing.T) {
		mockUseCase.On("CreateRecruitmentApplication", dto.RecruitmentApplicationRequest{
			RecruitmentID: 999,
		}, uint(1)).Return(errors.New("Error 1452: Cannot add or update a child row: a foreign key constraint fails (`compnouron`.`recruitment_applications`, CONSTRAINT `fk_recruitment_applications_recruitment` FOREIGN KEY (`recruitment_id`) REFERENCES `recruitments` (`id`))")).Once()

		// construct request body
		reqBody := dto.RecruitmentApplicationRequest{
			RecruitmentID: 999,
		}

		jsonReqBody, err := json.Marshal(&reqBody)
		assert.NoError(t, err, "No marshaling error")

		// setup the endpoint
		req, err := http.NewRequest(http.MethodPost, "/recruitments/applications", bytes.NewBuffer(jsonReqBody))
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")

		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		token := utils.CreateJWTToken(1, "gmail@gmail.com")
		c.Set("user", token)
		// setup controller/handler
		testRecruitmentController := RecruitmentController{
			router:        e,
			recruitmentUC: mockUseCase,
		}

		// get the response
		testRecruitmentController.CreateRecruitmentApplication(c)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUseCase.AssertExpectations(t)
	})
}
