package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	mocks "github.com/alimikegami/compnouron/internal/mocks/team/usecase"

	"github.com/alimikegami/compnouron/internal/team/dto"
	"github.com/alimikegami/compnouron/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateTeam(t *testing.T) {
	mockUseCase := mocks.NewTeamUseCase(t)

	// construct request body
	reqBody := dto.TeamRequest{
		Name:        "Team 1",
		Description: "Team Technoscape Hackathon 2022",
		Capacity:    4,
	}

	jsonReqBody, err := json.Marshal(&reqBody)
	assert.NoError(t, err, "No marshaling error")

	t.Run("success", func(t *testing.T) {
		mockUseCase.On("CreateTeam", uint(1), dto.TeamRequest{
			Name:        "Team 1",
			Description: "Team Technoscape Hackathon 2022",
			Capacity:    4,
		}).Return(nil).Once()
		req, err := http.NewRequest(http.MethodPost, "/teams", bytes.NewBuffer(jsonReqBody))
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")

		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		token := utils.CreateJWTToken(1, "gmail@gmail.com")
		c.Set("user", token)
		// setup controller/handler
		testTeamController := TeamController{
			router: e,
			teamUC: mockUseCase,
		}

		// get the response
		testTeamController.CreateTeam(c)
		assert.Equal(t, http.StatusCreated, rec.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUseCase.On("CreateTeam", uint(1), dto.TeamRequest{
			Name:        "Team 1",
			Description: "Team Technoscape Hackathon 2022",
			Capacity:    4,
		}).Return(errors.New("internal server error")).Once()
		req, err := http.NewRequest(http.MethodPost, "/teams", bytes.NewBuffer(jsonReqBody))
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")

		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		token := utils.CreateJWTToken(1, "gmail@gmail.com")
		c.Set("user", token)
		// setup controller/handler
		testTeamController := TeamController{
			router: e,
			teamUC: mockUseCase,
		}

		// get the response
		testTeamController.CreateTeam(c)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUseCase.AssertExpectations(t)
	})
}

func TestDeleteTeam(t *testing.T) {
	mockUseCase := mocks.NewTeamUseCase(t)

	t.Run("success", func(t *testing.T) {
		mockUseCase.On("DeleteTeam", uint(1), uint(1)).Return(nil).Once()

		// setup the endpoint
		req, err := http.NewRequest(http.MethodDelete, "/teams", nil)

		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		token := utils.CreateJWTToken(1, "gmail@gmail.com")
		c.Set("user", token)
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		// setup controller/handler
		testTeamController := TeamController{
			router: e,
			teamUC: mockUseCase,
		}

		// get the response
		testTeamController.DeleteTeam(c)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockUseCase.On("DeleteTeam", uint(1111), uint(1)).Return(errors.New("no affected rows")).Once()

		// setup the endpoint
		req, err := http.NewRequest(http.MethodDelete, "/teams", nil)

		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		token := utils.CreateJWTToken(1, "gmail@gmail.com")
		c.Set("user", token)
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("1111")
		// setup controller/handler
		testTeamController := TeamController{
			router: e,
			teamUC: mockUseCase,
		}

		// get the response
		testTeamController.DeleteTeam(c)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("action-unauthorized", func(t *testing.T) {
		mockUseCase.On("DeleteTeam", uint(1111), uint(2)).Return(errors.New("action unauthorized")).Once()

		// setup the endpoint
		req, err := http.NewRequest(http.MethodDelete, "/teams", nil)

		assert.NoError(t, err, "No request error")
		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		token := utils.CreateJWTToken(2, "gmail@gmail.com")
		c.Set("user", token)
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("1111")
		// setup controller/handler
		testTeamController := TeamController{
			router: e,
			teamUC: mockUseCase,
		}

		// get the response
		testTeamController.DeleteTeam(c)
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		mockUseCase.AssertExpectations(t)
	})
}

func TestUpdateTeam(t *testing.T) {
	mockUseCase := mocks.NewTeamUseCase(t)
	mockUseCase.On("UpdateTeam", uint(1), dto.TeamRequest{
		Name:        "Team 1",
		Description: "Team Technoscape Hackathon 2022",
		Capacity:    4,
	}, uint(1)).Return(nil)

	// construct request body
	reqBody := dto.TeamRequest{
		Name:        "Team 1",
		Description: "Team Technoscape Hackathon 2022",
		Capacity:    4,
	}

	jsonReqBody, err := json.Marshal(&reqBody)
	assert.NoError(t, err, "No marshaling error")

	// setup the endpoint
	req, err := http.NewRequest(http.MethodPut, "/teams", bytes.NewBuffer(jsonReqBody))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	assert.NoError(t, err, "No request error")
	e := echo.New()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	token := utils.CreateJWTToken(1, "gmail@gmail.com")
	c.Set("user", token)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// setup controller/handler
	testTeamController := TeamController{
		router: e,
		teamUC: mockUseCase,
	}

	// get the response
	testTeamController.UpdateTeam(c)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUseCase.AssertExpectations(t)
}

func TestUpdateTeamError(t *testing.T) {
	mockUseCase := mocks.NewTeamUseCase(t)
	mockUseCase.On("UpdateTeam", uint(1), dto.TeamRequest{
		Name:        "Team 1",
		Description: "Team Technoscape Hackathon 2022",
		Capacity:    4,
	}, uint(999)).Return(errors.New("no rows affected"))

	// construct request body
	reqBody := dto.TeamRequest{
		Name:        "Team 1",
		Description: "Team Technoscape Hackathon 2022",
		Capacity:    4,
	}

	jsonReqBody, err := json.Marshal(&reqBody)
	assert.NoError(t, err, "No marshaling error")

	// setup the endpoint
	req, err := http.NewRequest(http.MethodPut, "/teams", bytes.NewBuffer(jsonReqBody))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	assert.NoError(t, err, "No request error")
	e := echo.New()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	token := utils.CreateJWTToken(1, "gmail@gmail.com")
	c.Set("user", token)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("999")

	// setup controller/handler
	testTeamController := TeamController{
		router: e,
		teamUC: mockUseCase,
	}

	// get the response
	testTeamController.UpdateTeam(c)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUseCase.AssertExpectations(t)
}

func TestGetTeamsByUserID(t *testing.T) {
	mockUseCase := mocks.NewTeamUseCase(t)
	mockUseCase.On("GetTeamsByUserID", uint(1)).Return([]dto.BriefTeamResponse{
		{
			ID:   1,
			Name: "Team 1",
		},
		{
			ID:   2,
			Name: "Team 2",
		},
	}, nil)

	// setup the endpoint
	req, err := http.NewRequest(http.MethodGet, "/teams/users", nil)

	assert.NoError(t, err, "No request error")
	e := echo.New()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	token := utils.CreateJWTToken(1, "gmail@gmail.com")
	c.Set("user", token)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// setup controller/handler
	testTeamController := TeamController{
		router: e,
		teamUC: mockUseCase,
	}

	// get the response
	testTeamController.GetTeamsByUserID(c)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUseCase.AssertExpectations(t)
}

func TestGetTeamsByUserIDNoRowsFound(t *testing.T) {
	mockUseCase := mocks.NewTeamUseCase(t)
	mockUseCase.On("GetTeamsByUserID", uint(1)).Return([]dto.BriefTeamResponse{}, nil)

	// setup the endpoint
	req, err := http.NewRequest(http.MethodGet, "/teams/users", nil)

	assert.NoError(t, err, "No request error")
	e := echo.New()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	token := utils.CreateJWTToken(1, "gmail@gmail.com")
	c.Set("user", token)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// setup controller/handler
	testTeamController := TeamController{
		router: e,
		teamUC: mockUseCase,
	}

	// get the response
	testTeamController.GetTeamsByUserID(c)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUseCase.AssertExpectations(t)
}

func TestGetTeamDetailsByID(t *testing.T) {
	mockUseCase := mocks.NewTeamUseCase(t)
	mockUseCase.On("GetTeamDetailsByID", uint(1)).Return(dto.TeamDetailsResponse{
		Name:        "Team 1",
		Description: "Team Hackathon Technoscape 2022",
		Capacity:    4,
		TeamMembers: []dto.TeamMemberResponse{
			{
				UserID:   1,
				Name:     "Alim Ikegami",
				IsLeader: 1,
			},
		},
	}, nil)

	// setup the endpoint
	req, err := http.NewRequest(http.MethodGet, "/teams", nil)

	assert.NoError(t, err, "No request error")
	e := echo.New()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	token := utils.CreateJWTToken(1, "gmail@gmail.com")
	c.Set("user", token)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// setup controller/handler
	testTeamController := TeamController{
		router: e,
		teamUC: mockUseCase,
	}

	// get the response
	testTeamController.GetTeamDetailsByID(c)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUseCase.AssertExpectations(t)
}

func TestGetTeamDetailsByIDNoRowsFound(t *testing.T) {
	mockUseCase := mocks.NewTeamUseCase(t)
	mockUseCase.On("GetTeamDetailsByID", uint(1)).Return(dto.TeamDetailsResponse{}, errors.New("no rows found"))

	// setup the endpoint
	req, err := http.NewRequest(http.MethodGet, "/teams", nil)

	assert.NoError(t, err, "No request error")
	e := echo.New()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	token := utils.CreateJWTToken(1, "gmail@gmail.com")
	c.Set("user", token)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// setup controller/handler
	testTeamController := TeamController{
		router: e,
		teamUC: mockUseCase,
	}

	// get the response
	testTeamController.GetTeamDetailsByID(c)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUseCase.AssertExpectations(t)
}
