package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alimikegami/compnouron/internal/competition/dto"
	mocks "github.com/alimikegami/compnouron/internal/mocks/competition/usecase"
	"github.com/alimikegami/compnouron/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateCompetition(t *testing.T) {
	mockUseCase := mocks.NewCompetitionUseCase(t)

	// construct request body
	reqBody := dto.CompetitionRequest{
		Name:                 "Technoscape Hackathon 2022",
		Description:          "Hackathon dengan peserta sebanyak 4 orang per tim",
		ContactPerson:        "081239990128",
		IsTeam:               1,
		TeamCapacity:         4,
		Level:                "University Student",
		IsTheSameInstitution: 1,
	}

	jsonReqBody, err := json.Marshal(&reqBody)
	assert.NoError(t, err, "No marshaling error")

	// setup the endpoint
	t.Run("success", func(t *testing.T) {
		mockUseCase.On("CreateCompetition", reqBody, uint(1)).Return(nil).Once()
		req, err := http.NewRequest(http.MethodPost, "/competitions", bytes.NewBuffer(jsonReqBody))
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")

		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		token := utils.CreateJWTToken(uint(1), "gmail@gmail.com")
		c.Set("user", token)
		// setup controller/handler
		compController := CompetitionController{
			router:        e,
			CompetitionUC: mockUseCase,
		}

		// get the response
		compController.CreateCompetition(c)
		assert.Equal(t, http.StatusCreated, rec.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUseCase.On("CreateCompetition", reqBody, uint(1)).Return(errors.New("unexpected DB error")).Once()
		req, err := http.NewRequest(http.MethodPost, "/competitions", bytes.NewBuffer(jsonReqBody))
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")

		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		token := utils.CreateJWTToken(uint(1), "gmail@gmail.com")
		c.Set("user", token)
		// setup controller/handler
		compController := CompetitionController{
			router:        e,
			CompetitionUC: mockUseCase,
		}

		// get the response
		compController.CreateCompetition(c)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUseCase.AssertExpectations(t)
	})
}

func TestDeleteCompetition(t *testing.T) {
	mockUseCase := mocks.NewCompetitionUseCase(t)

	// setup the endpoint
	t.Run("success", func(t *testing.T) {
		mockUseCase.On("DeleteCompetition", uint(1), uint(1)).Return(nil).Once()
		req, err := http.NewRequest(http.MethodDelete, "/competitions", nil)

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
		compController := CompetitionController{
			router:        e,
			CompetitionUC: mockUseCase,
		}

		// get the response
		compController.DeleteCompetition(c)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUseCase.On("DeleteCompetition", uint(1), uint(1)).Return(errors.New("no record found")).Once()
		req, err := http.NewRequest(http.MethodDelete, "/competitions", nil)

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
		compController := CompetitionController{
			router:        e,
			CompetitionUC: mockUseCase,
		}

		// get the response
		compController.DeleteCompetition(c)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUseCase.AssertExpectations(t)
	})
}

func TestUpdateCompetition(t *testing.T) {
	mockUseCase := mocks.NewCompetitionUseCase(t)
	reqBody := dto.CompetitionRequest{
		Name:                 "Technoscape Hackathon 2022",
		Description:          "Hackathon dengan peserta sebanyak 4 orang per tim",
		ContactPerson:        "081239990128",
		IsTeam:               1,
		TeamCapacity:         4,
		Level:                "University Student",
		IsTheSameInstitution: 1,
	}

	jsonReqBody, err := json.Marshal(&reqBody)
	assert.NoError(t, err, "No marshaling error")
	// setup the endpoint
	t.Run("success", func(t *testing.T) {
		mockUseCase.On("UpdateCompetition", reqBody, uint(1), uint(1)).Return(nil).Once()
		req, err := http.NewRequest(http.MethodPut, "/competitions", bytes.NewBuffer(jsonReqBody))
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
		compController := CompetitionController{
			router:        e,
			CompetitionUC: mockUseCase,
		}

		// get the response
		compController.UpdateCompetition(c)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUseCase.On("UpdateCompetition", reqBody, uint(1), uint(1)).Return(errors.New("no affected rows")).Once()
		req, err := http.NewRequest(http.MethodPut, "/competitions", bytes.NewBuffer(jsonReqBody))
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
		compController := CompetitionController{
			router:        e,
			CompetitionUC: mockUseCase,
		}

		// get the response
		compController.UpdateCompetition(c)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUseCase.AssertExpectations(t)
	})
}

func TestRegister(t *testing.T) {
	mockUseCase := mocks.NewCompetitionUseCase(t)

	// construct request body
	reqBodyIndividual := dto.CompetitionRegistrationRequest{
		UserID:        1,
		CompetitionID: 1,
	}

	reqBodyTeam := dto.CompetitionRegistrationRequest{
		TeamID:        1,
		CompetitionID: 1,
	}

	jsonReqBodyIndividual, err := json.Marshal(&reqBodyIndividual)
	jsonReqBodyTeam, err := json.Marshal(&reqBodyTeam)
	assert.NoError(t, err, "No marshaling error")

	// setup the endpoint
	t.Run("success-individual-competition", func(t *testing.T) {
		mockUseCase.On("Register", reqBodyIndividual, uint(1)).Return(nil).Once()
		req, err := http.NewRequest(http.MethodPost, "/competitions/registrations", bytes.NewBuffer(jsonReqBodyIndividual))
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")

		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		token := utils.CreateJWTToken(uint(1), "gmail@gmail.com")
		c.Set("user", token)
		// setup controller/handler
		compController := CompetitionController{
			router:        e,
			CompetitionUC: mockUseCase,
		}

		// get the response
		compController.Register(c)
		assert.Equal(t, http.StatusCreated, rec.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("internal-server-error-individual", func(t *testing.T) {
		mockUseCase.On("Register", reqBodyIndividual, uint(1)).Return(errors.New("unexpected DB error")).Once()
		req, err := http.NewRequest(http.MethodPost, "/competitions/registrations", bytes.NewBuffer(jsonReqBodyIndividual))
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")

		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		token := utils.CreateJWTToken(uint(1), "gmail@gmail.com")
		c.Set("user", token)
		// setup controller/handler
		compController := CompetitionController{
			router:        e,
			CompetitionUC: mockUseCase,
		}

		// get the response
		compController.Register(c)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("success-team-competition", func(t *testing.T) {
		mockUseCase.On("Register", reqBodyTeam, uint(1)).Return(nil).Once()
		req, err := http.NewRequest(http.MethodPost, "/competitions/registrations", bytes.NewBuffer(jsonReqBodyTeam))
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")

		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		token := utils.CreateJWTToken(uint(1), "gmail@gmail.com")
		c.Set("user", token)
		// setup controller/handler
		compController := CompetitionController{
			router:        e,
			CompetitionUC: mockUseCase,
		}

		// get the response
		compController.Register(c)
		assert.Equal(t, http.StatusCreated, rec.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("internal-server-error-team", func(t *testing.T) {
		mockUseCase.On("Register", reqBodyTeam, uint(1)).Return(errors.New("unexpected DB error")).Once()
		req, err := http.NewRequest(http.MethodPost, "/competitions/registrations", bytes.NewBuffer(jsonReqBodyTeam))
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")

		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		token := utils.CreateJWTToken(uint(1), "gmail@gmail.com")
		c.Set("user", token)
		// setup controller/handler
		compController := CompetitionController{
			router:        e,
			CompetitionUC: mockUseCase,
		}

		// get the response
		compController.Register(c)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUseCase.AssertExpectations(t)
	})
}

func TestAcceptCompetitionRegistration(t *testing.T) {
	mockUseCase := mocks.NewCompetitionUseCase(t)

	// setup the endpoint
	t.Run("success", func(t *testing.T) {
		mockUseCase.On("AcceptCompetitionRegistration", uint(1), uint(1)).Return(nil).Once()
		req, err := http.NewRequest(http.MethodPut, "/competitions/registrations", nil)
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
		compController := CompetitionController{
			router:        e,
			CompetitionUC: mockUseCase,
		}

		// get the response
		compController.AcceptCompetitionRegistration(c)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUseCase.On("AcceptCompetitionRegistration", uint(1), uint(1)).Return(errors.New("no affected rows")).Once()
		req, err := http.NewRequest(http.MethodPut, "/competitions/registrations", nil)
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
		compController := CompetitionController{
			router:        e,
			CompetitionUC: mockUseCase,
		}

		// get the response
		compController.AcceptCompetitionRegistration(c)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUseCase.AssertExpectations(t)
	})
}

func TestRejectCompetitionRegistration(t *testing.T) {
	mockUseCase := mocks.NewCompetitionUseCase(t)

	// setup the endpoint
	t.Run("success", func(t *testing.T) {
		mockUseCase.On("RejectCompetitionRegistration", uint(1), uint(1)).Return(nil).Once()
		req, err := http.NewRequest(http.MethodPut, "/competitions/registrations", nil)
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
		compController := CompetitionController{
			router:        e,
			CompetitionUC: mockUseCase,
		}

		// get the response
		compController.RejectCompetitionRegistration(c)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUseCase.On("RejectCompetitionRegistration", uint(1), uint(1)).Return(errors.New("no affected rows")).Once()
		req, err := http.NewRequest(http.MethodPut, "/competitions/registrations", nil)
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
		compController := CompetitionController{
			router:        e,
			CompetitionUC: mockUseCase,
		}

		// get the response
		compController.RejectCompetitionRegistration(c)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUseCase.AssertExpectations(t)
	})
}

func TestOpenCompetitionRegistrationPeriod(t *testing.T) {
	mockUseCase := mocks.NewCompetitionUseCase(t)
	// setup the endpoint
	t.Run("success", func(t *testing.T) {
		mockUseCase.On("OpenCompetitionRegistrationPeriod", uint(1), uint(1)).Return(nil).Once()
		req, err := http.NewRequest(http.MethodPut, "/competitions", nil)
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
		compController := CompetitionController{
			router:        e,
			CompetitionUC: mockUseCase,
		}

		// get the response
		compController.OpenCompetitionRegistrationPeriod(c)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUseCase.On("OpenCompetitionRegistrationPeriod", uint(1), uint(1)).Return(errors.New("no affected rows")).Once()
		req, err := http.NewRequest(http.MethodPut, "/competitions", nil)
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
		compController := CompetitionController{
			router:        e,
			CompetitionUC: mockUseCase,
		}

		// get the response
		compController.OpenCompetitionRegistrationPeriod(c)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUseCase.AssertExpectations(t)
	})
}

func TestCloseCompetitionRegistrationPeriod(t *testing.T) {
	mockUseCase := mocks.NewCompetitionUseCase(t)
	// setup the endpoint
	t.Run("success", func(t *testing.T) {
		mockUseCase.On("CloseCompetitionRegistrationPeriod", uint(1), uint(1)).Return(nil).Once()
		req, err := http.NewRequest(http.MethodPut, "/competitions", nil)
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
		compController := CompetitionController{
			router:        e,
			CompetitionUC: mockUseCase,
		}

		// get the response
		compController.CloseCompetitionRegistrationPeriod(c)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUseCase.On("CloseCompetitionRegistrationPeriod", uint(1), uint(1)).Return(errors.New("no affected rows")).Once()
		req, err := http.NewRequest(http.MethodPut, "/competitions", nil)
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
		compController := CompetitionController{
			router:        e,
			CompetitionUC: mockUseCase,
		}

		// get the response
		compController.CloseCompetitionRegistrationPeriod(c)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUseCase.AssertExpectations(t)
	})
}

func TestGetCompetitionByID(t *testing.T) {
	mockUseCase := mocks.NewCompetitionUseCase(t)
	// setup the endpoint
	t.Run("success", func(t *testing.T) {
		mockUseCase.On("GetCompetitionByID", uint(1)).Return(dto.DetailedCompetitionResponse{
			ID:                       1,
			Name:                     "Technoscape",
			Description:              "Testing",
			ContactPerson:            "081313131",
			IsTheSameInstitution:     1,
			IsTeam:                   1,
			RegistrationPeriodStatus: 1,
			TeamCapacity:             4,
			Level:                    "University Student",
			UserID:                   1,
			UserName:                 "BNCC",
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
		compController := CompetitionController{
			router:        e,
			CompetitionUC: mockUseCase,
		}

		// get the response
		compController.GetCompetitionByID(c)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUseCase.On("GetCompetitionByID", uint(1)).Return(dto.DetailedCompetitionResponse{}, errors.New("unexpected error occured")).Once()
		req, err := http.NewRequest(http.MethodGet, "/competitions", nil)
		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		// setup controller/handler
		compController := CompetitionController{
			router:        e,
			CompetitionUC: mockUseCase,
		}

		// get the response
		compController.GetCompetitionByID(c)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUseCase.AssertExpectations(t)
	})
}

// func TestGetCompetitions(t *testing.T) {
// 	mockUseCase := mocks.NewCompetitionUseCase(t)
// 	// setup the endpoint
// 	t.Run("success", func(t *testing.T) {
// 		mockUseCase.On("GetCompetitions", uint(1)).Return(dto.DetailedCompetitionResponse{
// 			ID:                       1,
// 			Name:                     "Technoscape",
// 			Description:              "Testing",
// 			ContactPerson:            "081313131",
// 			IsTheSameInstitution:     1,
// 			IsTeam:                   1,
// 			RegistrationPeriodStatus: 1,
// 			TeamCapacity:             4,
// 			Level:                    "University Student",
// 			UserID:                   1,
// 			UserName:                 "BNCC",
// 		}, nil).Once()
// 		req, err := http.NewRequest(http.MethodGet, "/competitions", nil)
// 		assert.NoError(t, err, "No request error")
// 		e := echo.New()
// 		rec := httptest.NewRecorder()
// 		c := e.NewContext(req, rec)
// 		c.SetPath("/:id")
// 		c.SetParamNames("id")
// 		c.SetParamValues("1")
// 		// setup controller/handler
// 		compController := CompetitionController{
// 			router:        e,
// 			CompetitionUC: mockUseCase,
// 		}

// 		// get the response
// 		compController.GetCompetitions(c)
// 		assert.Equal(t, http.StatusOK, rec.Code)
// 		mockUseCase.AssertExpectations(t)
// 	})

// 	t.Run("internal-server-error", func(t *testing.T) {
// 		mockUseCase.On("GetCompetitions", uint(1)).Return(dto.DetailedCompetitionResponse{}, errors.New("unexpected error occured")).Once()
// 		req, err := http.NewRequest(http.MethodGet, "/competitions", nil)
// 		assert.NoError(t, err, "No request error")
// 		e := echo.New()
// 		rec := httptest.NewRecorder()
// 		c := e.NewContext(req, rec)
// 		c.SetPath("/:id")
// 		c.SetParamNames("id")
// 		c.SetParamValues("1")
// 		// setup controller/handler
// 		compController := CompetitionController{
// 			router:        e,
// 			CompetitionUC: mockUseCase,
// 		}

// 		// get the response
// 		compController.GetCompetitions(c)
// 		assert.Equal(t, http.StatusInternalServerError, rec.Code)
// 		mockUseCase.AssertExpectations(t)
// 	})
// }
