package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	compDto "github.com/alimikegami/compnouron/internal/competition/dto"
	mocks "github.com/alimikegami/compnouron/internal/mocks/user/usecase"
	"github.com/alimikegami/compnouron/internal/user/dto"
	"github.com/alimikegami/compnouron/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	mockUseCase := mocks.NewUserUseCase(t)
	mockUseCase.On("CreateUser", &dto.UserRegistrationRequest{
		Name:              "Alim Ikegami",
		Email:             "sdafsfa@gmail.com",
		PhoneNumber:       "081111111111",
		Password:          "asdfasfas",
		SchoolInstitution: "Udayana University",
		Skills: []dto.SkillRequest{
			{
				Name: "Node.J",
			},
		},
	}).Return(nil)

	// construct request body
	reqBody := dto.UserRegistrationRequest{
		Name:              "Alim Ikegami",
		Email:             "sdafsfa@gmail.com",
		PhoneNumber:       "081111111111",
		Password:          "asdfasfas",
		SchoolInstitution: "Udayana University",
		Skills: []dto.SkillRequest{
			{
				Name: "Node.J",
			},
		},
	}

	jsonReqBody, err := json.Marshal(&reqBody)
	assert.NoError(t, err, "No marshaling error")

	// setup the endpoint
	req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonReqBody))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	assert.NoError(t, err, "No request error")
	e := echo.New()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// setup controller/handler
	testUserController := UserController{
		router: e,
		userUC: mockUseCase,
	}

	// get the response
	testUserController.CreateUser(c)
	assert.Equal(t, http.StatusCreated, rec.Code)
	mockUseCase.AssertExpectations(t)
}

func TestCreateUserDuplicateEmail(t *testing.T) {
	mockUseCase := mocks.NewUserUseCase(t)
	mockUseCase.On("CreateUser", &dto.UserRegistrationRequest{
		Name:              "Alim Ikegami",
		Email:             "sdafsfa@gmail.com",
		PhoneNumber:       "081111111111",
		Password:          "asdfasfas",
		SchoolInstitution: "Udayana University",
		Skills: []dto.SkillRequest{
			{
				Name: "Node.J",
			},
		},
	}).Return(errors.New("duplicate emails"))

	// construct request body
	reqBody := dto.UserRegistrationRequest{
		Name:              "Alim Ikegami",
		Email:             "sdafsfa@gmail.com",
		PhoneNumber:       "081111111111",
		Password:          "asdfasfas",
		SchoolInstitution: "Udayana University",
		Skills: []dto.SkillRequest{
			{
				Name: "Node.J",
			},
		},
	}

	jsonReqBody, err := json.Marshal(&reqBody)
	assert.NoError(t, err, "No marshaling error")

	// setup the endpoint
	req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonReqBody))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	assert.NoError(t, err, "No request error")
	e := echo.New()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// setup controller/handler
	testUserController := UserController{
		router: e,
		userUC: mockUseCase,
	}

	// get the response
	testUserController.CreateUser(c)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUseCase.AssertExpectations(t)
}

func TestLogin(t *testing.T) {
	mockUseCase := mocks.NewUserUseCase(t)
	mockUseCase.On("Login", &dto.Credential{
		Email:    "sdafsfa@gmail.com",
		Password: "asdfasfas",
	}).Return("sdafasfasfsafasdfasdfasfasfasdf", nil)

	// construct request body
	reqBody := dto.Credential{
		Email:    "sdafsfa@gmail.com",
		Password: "asdfasfas",
	}

	jsonReqBody, err := json.Marshal(&reqBody)
	assert.NoError(t, err, "No marshaling error")

	// setup the endpoint
	req, err := http.NewRequest(http.MethodPost, "/users/login", bytes.NewBuffer(jsonReqBody))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	assert.NoError(t, err, "No request error")
	e := echo.New()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// setup controller/handler
	testUserController := UserController{
		router: e,
		userUC: mockUseCase,
	}

	// get the response
	testUserController.Login(c)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUseCase.AssertExpectations(t)
}

func TestLoginWrongCredentials(t *testing.T) {
	mockUseCase := mocks.NewUserUseCase(t)
	mockUseCase.On("Login", &dto.Credential{
		Email:    "sdafsfa@gmail.com",
		Password: "asdfasfas1",
	}).Return("", errors.New("credentials dont match"))

	// construct request body
	reqBody := dto.Credential{
		Email:    "sdafsfa@gmail.com",
		Password: "asdfasfas1",
	}

	jsonReqBody, err := json.Marshal(&reqBody)
	assert.NoError(t, err, "No marshaling error")

	// setup the endpoint
	req, err := http.NewRequest(http.MethodPost, "/users/login", bytes.NewBuffer(jsonReqBody))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	assert.NoError(t, err, "No request error")
	e := echo.New()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// setup controller/handler
	testUserController := UserController{
		router: e,
		userUC: mockUseCase,
	}

	// get the response
	testUserController.Login(c)
	assert.Equal(t, 403, rec.Code)
	mockUseCase.AssertExpectations(t)
}

func TestGetRecruitmentApplicationHistory(t *testing.T) {
	mockUseCase := mocks.NewUserUseCase(t)
	// setup the endpoint
	t.Run("success", func(t *testing.T) {
		mockUseCase.On("GetRecruitmentApplicationHistory", uint(1)).Return([]dto.UserRecruitmentApplicationHistory{
			{
				RecruitmentApplicationID: 1,
				RecruitmentID:            1,
				RecruitmentRole:          "Backend Engineer",
				AcceptanceStatus:         1,
			},
		}, nil).Once()
		req, err := http.NewRequest(http.MethodGet, "/users/recruitments/applications", nil)
		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		token := utils.CreateJWTToken(1, "gmail@gmail.com")
		c.Set("user", token)
		// setup controller/handler
		userController := UserController{
			router: e,
			userUC: mockUseCase,
		}

		// get the response
		userController.GetRecruitmentApplicationHistory(c)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUseCase.On("GetRecruitmentApplicationHistory", uint(1)).Return([]dto.UserRecruitmentApplicationHistory{}, errors.New("unexpected error occured")).Once()
		req, err := http.NewRequest(http.MethodGet, "/users/recruitments/applications", nil)
		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		token := utils.CreateJWTToken(1, "gmail@gmail.com")
		c.Set("user", token)
		// setup controller/handler
		userController := UserController{
			router: e,
			userUC: mockUseCase,
		}

		// get the response
		userController.GetRecruitmentApplicationHistory(c)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUseCase.AssertExpectations(t)
	})
}

func TestGetCompetitionRegistrationHistory(t *testing.T) {
	mockUseCase := mocks.NewUserUseCase(t)
	// setup the endpoint
	t.Run("success", func(t *testing.T) {
		mockUseCase.On("GetCompetitionRegistrationHistory", uint(1)).Return([]dto.UserCompetitionHistory{
			{
				CompetitionRegistrationID: 1,
				CompetitionID:             1,
				CompetitionName:           "Technoscape",
				AcceptanceStatus:          1,
			},
		}, nil).Once()
		req, err := http.NewRequest(http.MethodGet, "/users/competitions/registrations", nil)
		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		token := utils.CreateJWTToken(1, "gmail@gmail.com")
		c.Set("user", token)
		// setup controller/handler
		userController := UserController{
			router: e,
			userUC: mockUseCase,
		}

		// get the response
		userController.GetCompetitionRegistrationHistory(c)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUseCase.On("GetCompetitionRegistrationHistory", uint(1)).Return([]dto.UserCompetitionHistory{}, errors.New("unexpected error occured")).Once()
		req, err := http.NewRequest(http.MethodGet, "/users/competitions/registrations", nil)
		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		token := utils.CreateJWTToken(1, "gmail@gmail.com")
		c.Set("user", token)
		// setup controller/handler
		userController := UserController{
			router: e,
			userUC: mockUseCase,
		}

		// get the response
		userController.GetCompetitionRegistrationHistory(c)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUseCase.AssertExpectations(t)
	})
}

func TestGetCompetitionsData(t *testing.T) {
	mockUseCase := mocks.NewUserUseCase(t)
	// setup the endpoint
	t.Run("success", func(t *testing.T) {
		mockUseCase.On("GetCompetitionsData", uint(1)).Return([]compDto.CompetitionResponse{
			{
				ID:            1,
				Name:          "Technoscape",
				ContactPerson: "081234782",
				IsTeam:        1,
				Level:         "University Student",
			},
		}, nil).Once()
		req, err := http.NewRequest(http.MethodGet, "/users", nil)
		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/:id/competitions")
		c.SetParamNames("id")
		c.SetParamValues("1")
		// setup controller/handler
		userController := UserController{
			router: e,
			userUC: mockUseCase,
		}

		// get the response
		userController.GetCompetitionsData(c)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUseCase.On("GetCompetitionsData", uint(1)).Return([]compDto.CompetitionResponse{}, errors.New("unexpected error occured")).Once()
		req, err := http.NewRequest(http.MethodGet, "/users", nil)
		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/:id/competitions")
		c.SetParamNames("id")
		c.SetParamValues("1")
		// setup controller/handler
		userController := UserController{
			router: e,
			userUC: mockUseCase,
		}

		// get the response
		userController.GetCompetitionsData(c)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUseCase.AssertExpectations(t)
	})
}
