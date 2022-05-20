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

func TestAcceptRecruitmentApplication(t *testing.T) {
	mockUseCase := mocks.NewRecruitmentUseCase(t)
	// setup the endpoint
	t.Run("success", func(t *testing.T) {
		mockUseCase.On("AcceptRecruitmentApplication", uint(1), uint(1)).Return(nil).Once()
		req, err := http.NewRequest(http.MethodPut, "/recruitments/applications", nil)
		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		token := utils.CreateJWTToken(uint(1), "gmail@gmail.com")
		c.Set("user", token)
		c.SetPath("/:id/accept")
		c.SetParamNames("id")
		c.SetParamValues("1")
		// setup controller/handler
		compController := RecruitmentController{
			router:        e,
			recruitmentUC: mockUseCase,
		}

		// get the response
		compController.AcceptRecruitmentApplication(c)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUseCase.On("AcceptRecruitmentApplication", uint(1), uint(1)).Return(errors.New("no affected rows")).Once()
		req, err := http.NewRequest(http.MethodPut, "/recruitments/applications", nil)
		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		token := utils.CreateJWTToken(uint(1), "gmail@gmail.com")
		c.Set("user", token)
		c.SetPath("/:id/accept")
		c.SetParamNames("id")
		c.SetParamValues("1")
		// setup controller/handler
		compController := RecruitmentController{
			router:        e,
			recruitmentUC: mockUseCase,
		}

		// get the response
		compController.AcceptRecruitmentApplication(c)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUseCase.AssertExpectations(t)
	})
}

func TestRejectRecruitmentApplication(t *testing.T) {
	mockUseCase := mocks.NewRecruitmentUseCase(t)
	// setup the endpoint
	t.Run("success", func(t *testing.T) {
		mockUseCase.On("RejectRecruitmentApplication", uint(1), uint(1)).Return(nil).Once()
		req, err := http.NewRequest(http.MethodPut, "/recruitments/applications", nil)
		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		token := utils.CreateJWTToken(uint(1), "gmail@gmail.com")
		c.Set("user", token)
		c.SetPath("/:id/reject")
		c.SetParamNames("id")
		c.SetParamValues("1")
		// setup controller/handler
		compController := RecruitmentController{
			router:        e,
			recruitmentUC: mockUseCase,
		}

		// get the response
		compController.RejectRecruitmentApplication(c)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUseCase.On("RejectRecruitmentApplication", uint(1), uint(1)).Return(errors.New("no affected rows")).Once()
		req, err := http.NewRequest(http.MethodPut, "/recruitments/applications", nil)
		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		token := utils.CreateJWTToken(uint(1), "gmail@gmail.com")
		c.Set("user", token)
		c.SetPath("/:id/reject")
		c.SetParamNames("id")
		c.SetParamValues("1")
		// setup controller/handler
		compController := RecruitmentController{
			router:        e,
			recruitmentUC: mockUseCase,
		}

		// get the response
		compController.RejectRecruitmentApplication(c)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUseCase.AssertExpectations(t)
	})
}

func TestOpenRecruitmentApplicationPeriod(t *testing.T) {
	mockUseCase := mocks.NewRecruitmentUseCase(t)
	// setup the endpoint
	t.Run("success", func(t *testing.T) {
		mockUseCase.On("OpenRecruitmentApplicationPeriod", uint(1), uint(1)).Return(nil).Once()
		req, err := http.NewRequest(http.MethodPut, "/recruitments", nil)
		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		token := utils.CreateJWTToken(uint(1), "gmail@gmail.com")
		c.Set("user", token)
		c.SetPath("/:id/open")
		c.SetParamNames("id")
		c.SetParamValues("1")
		// setup controller/handler
		compController := RecruitmentController{
			router:        e,
			recruitmentUC: mockUseCase,
		}

		// get the response
		compController.OpenRecruitmentApplicationPeriod(c)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUseCase.On("OpenRecruitmentApplicationPeriod", uint(1), uint(1)).Return(errors.New("no affected rows")).Once()
		req, err := http.NewRequest(http.MethodPut, "/recruitments", nil)
		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		token := utils.CreateJWTToken(uint(1), "gmail@gmail.com")
		c.Set("user", token)
		c.SetPath("/:id/open")
		c.SetParamNames("id")
		c.SetParamValues("1")
		// setup controller/handler
		compController := RecruitmentController{
			router:        e,
			recruitmentUC: mockUseCase,
		}

		// get the response
		compController.OpenRecruitmentApplicationPeriod(c)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUseCase.AssertExpectations(t)
	})
}

func TestCloseRecruitmentApplicationPeriod(t *testing.T) {
	mockUseCase := mocks.NewRecruitmentUseCase(t)
	// setup the endpoint
	t.Run("success", func(t *testing.T) {
		mockUseCase.On("CloseRecruitmentApplicationPeriod", uint(1), uint(1)).Return(nil).Once()
		req, err := http.NewRequest(http.MethodPut, "/recruitments", nil)
		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		token := utils.CreateJWTToken(uint(1), "gmail@gmail.com")
		c.Set("user", token)
		c.SetPath("/:id/close")
		c.SetParamNames("id")
		c.SetParamValues("1")
		// setup controller/handler
		compController := RecruitmentController{
			router:        e,
			recruitmentUC: mockUseCase,
		}

		// get the response
		compController.CloseRecruitmentApplicationPeriod(c)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUseCase.On("CloseRecruitmentApplicationPeriod", uint(1), uint(1)).Return(errors.New("no affected rows")).Once()
		req, err := http.NewRequest(http.MethodPut, "/recruitments", nil)
		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		token := utils.CreateJWTToken(uint(1), "gmail@gmail.com")
		c.Set("user", token)
		c.SetPath("/:id/close")
		c.SetParamNames("id")
		c.SetParamValues("1")
		// setup controller/handler
		compController := RecruitmentController{
			router:        e,
			recruitmentUC: mockUseCase,
		}

		// get the response
		compController.CloseRecruitmentApplicationPeriod(c)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUseCase.AssertExpectations(t)
	})
}

func TestDeleteRecruitmentByID(t *testing.T) {
	mockUseCase := mocks.NewRecruitmentUseCase(t)
	// setup the endpoint
	t.Run("success", func(t *testing.T) {
		mockUseCase.On("DeleteRecruitmentByID", uint(1), uint(1)).Return(nil).Once()
		req, err := http.NewRequest(http.MethodDelete, "/recruitments", nil)
		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		token := utils.CreateJWTToken(uint(1), "gmail@gmail.com")
		c.Set("user", token)
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		// setup controller/handler
		compController := RecruitmentController{
			router:        e,
			recruitmentUC: mockUseCase,
		}

		// get the response
		compController.DeleteRecruitmentByID(c)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUseCase.On("DeleteRecruitmentByID", uint(1), uint(1)).Return(errors.New("no affected rows")).Once()
		req, err := http.NewRequest(http.MethodDelete, "/recruitments", nil)
		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		token := utils.CreateJWTToken(uint(1), "gmail@gmail.com")
		c.Set("user", token)
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		// setup controller/handler
		compController := RecruitmentController{
			router:        e,
			recruitmentUC: mockUseCase,
		}

		// get the response
		compController.DeleteRecruitmentByID(c)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUseCase.AssertExpectations(t)
	})
}

func TestUpdateRecruitment(t *testing.T) {
	mockUseCase := mocks.NewRecruitmentUseCase(t)
	reqBody := dto.RecruitmentRequest{
		Role:        "Backend Engineer",
		Description: "Bisa cuci piring",
		TeamID:      1,
	}

	jsonReqBody, err := json.Marshal(&reqBody)
	assert.NoError(t, err, "No marshaling error")
	// setup the endpoint
	t.Run("success", func(t *testing.T) {
		mockUseCase.On("UpdateRecruitment", reqBody, uint(1), uint(1)).Return(nil).Once()
		req, err := http.NewRequest(http.MethodPut, "/recruitments", bytes.NewBuffer(jsonReqBody))
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		token := utils.CreateJWTToken(uint(1), "gmail@gmail.com")
		c.Set("user", token)
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		// setup controller/handler
		compController := RecruitmentController{
			router:        e,
			recruitmentUC: mockUseCase,
		}

		// get the response
		compController.UpdateRecruitment(c)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUseCase.On("UpdateRecruitment", reqBody, uint(1), uint(1)).Return(errors.New("no affected rows")).Once()
		req, err := http.NewRequest(http.MethodPut, "/recruitments", bytes.NewBuffer(jsonReqBody))
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		token := utils.CreateJWTToken(uint(1), "gmail@gmail.com")
		c.Set("user", token)
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		// setup controller/handler
		compController := RecruitmentController{
			router:        e,
			recruitmentUC: mockUseCase,
		}

		// get the response
		compController.UpdateRecruitment(c)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUseCase.AssertExpectations(t)
	})
}

func TestGetRecruitmentByID(t *testing.T) {
	mockUseCase := mocks.NewRecruitmentUseCase(t)
	// setup the endpoint
	t.Run("success", func(t *testing.T) {
		mockUseCase.On("GetRecruitmentByID", uint(1)).Return(dto.RecruitmentResponse{
			ID:          1,
			Role:        "Backend Engineer",
			Description: "Bisa cuci piring",
			TeamID:      1,
			TeamName:    "Perceptron",
		}, nil).Once()
		req, err := http.NewRequest(http.MethodGet, "/competitions", nil)
		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		// setup controller/handler
		compController := RecruitmentController{
			router:        e,
			recruitmentUC: mockUseCase,
		}

		// get the response
		compController.GetRecruitmentByID(c)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUseCase.On("GetRecruitmentByID", uint(1)).Return(dto.RecruitmentResponse{}, errors.New("unexpected error occured")).Once()
		req, err := http.NewRequest(http.MethodGet, "/competitions", nil)
		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		// setup controller/handler
		compController := RecruitmentController{
			router:        e,
			recruitmentUC: mockUseCase,
		}

		// get the response
		compController.GetRecruitmentByID(c)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUseCase.AssertExpectations(t)
	})
}

func TestGetRecruitmentByTeamID(t *testing.T) {
	mockUseCase := mocks.NewRecruitmentUseCase(t)
	// setup the endpoint
	t.Run("success", func(t *testing.T) {
		mockUseCase.On("GetRecruitmentByTeamID", uint(1)).Return(dto.RecruitmentsResponse{
			{
				ID:          1,
				Role:        "Backend Engineer",
				Description: "Bisa cuci piring",
				TeamID:      1,
				TeamName:    "Perceptron",
			},
			{
				ID:          1,
				Role:        "Backend Engineer",
				Description: "Bisa cuci piring",
				TeamID:      1,
				TeamName:    "Perceptron",
			},
		}, nil).Once()
		req, err := http.NewRequest(http.MethodGet, "/competitions/teams", nil)
		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		// setup controller/handler
		compController := RecruitmentController{
			router:        e,
			recruitmentUC: mockUseCase,
		}

		// get the response
		compController.GetRecruitmentByTeamID(c)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUseCase.On("GetRecruitmentByTeamID", uint(1)).Return(dto.RecruitmentsResponse{}, errors.New("unexpected error occured")).Once()
		req, err := http.NewRequest(http.MethodGet, "/competitions/teams", nil)
		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		// setup controller/handler
		compController := RecruitmentController{
			router:        e,
			recruitmentUC: mockUseCase,
		}

		// get the response
		compController.GetRecruitmentByTeamID(c)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUseCase.AssertExpectations(t)
	})
}
