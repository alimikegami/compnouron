package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	mocks "github.com/alimikegami/compnouron/internal/mocks/user/usecase"
	"github.com/alimikegami/compnouron/internal/user/dto"
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
