package usecase

import (
	"errors"
	"testing"
	"time"

	teamMocks "github.com/alimikegami/compnouron/internal/mocks/team/repository"
	"github.com/alimikegami/compnouron/internal/team/dto"
	"github.com/alimikegami/compnouron/internal/team/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateTeam(t *testing.T) {
	team := entity.Team{
		Name:        "Team 1",
		Description: "Team Technoscape Hackathon 2022",
		Capacity:    4,
	}
	createdTeam := entity.Team{
		ID:          1,
		Name:        "Team 1",
		Description: "Team Technoscape Hackathon 2022",
		Capacity:    4,
	}
	teamMockRepo := teamMocks.NewTeamRepository(t)
	teamMockRepo.On("CreateTeam", team).Return(createdTeam, nil)

	teamMockRepo.On("AddTeamMember", uint(1), createdTeam.ID, uint(1)).Return(nil)

	testUseCase := CreateNewTeamUseCase(teamMockRepo)
	err := testUseCase.CreateTeam(1, dto.TeamRequest{
		Name:        "Team 1",
		Description: "Team Technoscape Hackathon 2022",
		Capacity:    4,
	})

	assert.NoError(t, err)
	teamMockRepo.AssertExpectations(t)
}

func TestDeleteTeam(t *testing.T) {
	mockRepo := teamMocks.NewTeamRepository(t)
	mockRepo.On("DeleteTeam", uint(1)).Return(nil)
	testUseCase := CreateNewTeamUseCase(mockRepo)
	err := testUseCase.DeleteTeam(uint(1))
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteTeamErrorOccured(t *testing.T) {
	mockRepo := teamMocks.NewTeamRepository(t)
	mockRepo.On("DeleteTeam", uint(111)).Return(errors.New("no rows affected"))
	testUseCase := CreateNewTeamUseCase(mockRepo)
	err := testUseCase.DeleteTeam(uint(111))
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateTeam(t *testing.T) {
	mockRepo := teamMocks.NewTeamRepository(t)
	mockRepo.On("UpdateTeam", entity.Team{
		ID:          1,
		Name:        "Team 1",
		Description: "Team Technoscape Hackathon 2022",
		Capacity:    4,
	}).Return(nil)
	testUseCase := CreateNewTeamUseCase(mockRepo)
	err := testUseCase.UpdateTeam(1, dto.TeamRequest{
		Name:        "Team 1",
		Description: "Team Technoscape Hackathon 2022",
		Capacity:    4,
	}, 1)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateTeamErrorOccured(t *testing.T) {
	mockRepo := teamMocks.NewTeamRepository(t)
	mockRepo.On("UpdateTeam", entity.Team{
		ID:          9999,
		Name:        "Team 1",
		Description: "Team Technoscape Hackathon 2022",
		Capacity:    4,
	}).Return(errors.New("no affected rows"))
	testUseCase := CreateNewTeamUseCase(mockRepo)
	err := testUseCase.UpdateTeam(1, dto.TeamRequest{
		Name:        "Team 1",
		Description: "Team Technoscape Hackathon 2022",
		Capacity:    4,
	}, 9999)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetTeamsByUserID(t *testing.T) {
	mockRepo := teamMocks.NewTeamRepository(t)
	mockRepo.On("GetTeamsByUserID", uint(1)).Return([]entity.Team{
		{
			ID:          1,
			Name:        "Team 1",
			Description: "Team Technoscape Hackathon 2022",
			Capacity:    4,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          2,
			Name:        "Team 2",
			Description: "Team Invention 2022",
			Capacity:    3,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}, nil)
	testUseCase := CreateNewTeamUseCase(mockRepo)
	res, err := testUseCase.GetTeamsByUserID(1)
	assert.NoError(t, err)
	assert.Len(t, res, 2)
	mockRepo.AssertExpectations(t)
}

func TestGetTeamsByUserIDErrorOccured(t *testing.T) {
	mockRepo := teamMocks.NewTeamRepository(t)
	mockRepo.On("GetTeamsByUserID", uint(111)).Return([]entity.Team{}, nil)
	testUseCase := CreateNewTeamUseCase(mockRepo)
	res, err := testUseCase.GetTeamsByUserID(111)
	assert.NoError(t, err)
	assert.Len(t, res, 0)
	mockRepo.AssertExpectations(t)
}

func TestGetTeamDetailsByID(t *testing.T) {
	mockRepo := teamMocks.NewTeamRepository(t)
	mockRepo.On("GetTeamByID", uint(1)).Return(entity.Team{
		ID:          1,
		Name:        "Team 1",
		Description: "Team Technoscape Hackathon 2022",
		Capacity:    4,
	}, nil)
	mockRepo.On("GetTeamMembersByID", uint(1)).Return([]entity.TeamMember{
		{
			ID:        1,
			TeamID:    1,
			UserID:    1,
			IsLeader:  1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			TeamID:    1,
			UserID:    2,
			IsLeader:  0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}, nil)
	testUseCase := CreateNewTeamUseCase(mockRepo)
	res, err := testUseCase.GetTeamDetailsByID(uint(1))
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	mockRepo.AssertExpectations(t)
}

func TestGetTeamDetailsByIDErrorInGetTeamID(t *testing.T) {
	mockRepo := teamMocks.NewTeamRepository(t)
	mockRepo.On("GetTeamByID", uint(1)).Return(entity.Team{
		ID:          1,
		Name:        "Team 1",
		Description: "Team Technoscape Hackathon 2022",
		Capacity:    4,
	}, errors.New("no rows found"))
	testUseCase := CreateNewTeamUseCase(mockRepo)
	res, err := testUseCase.GetTeamDetailsByID(uint(1))
	assert.Error(t, err)
	assert.Empty(t, res)
	mockRepo.AssertExpectations(t)
}

func TestGetTeamDetailsByIDTeamMembersNotFound(t *testing.T) {
	mockRepo := teamMocks.NewTeamRepository(t)
	mockRepo.On("GetTeamByID", uint(1)).Return(entity.Team{
		ID:          1,
		Name:        "Team 1",
		Description: "Team Technoscape Hackathon 2022",
		Capacity:    4,
	}, nil)
	mockRepo.On("GetTeamMembersByID", uint(1)).Return([]entity.TeamMember{}, nil)
	testUseCase := CreateNewTeamUseCase(mockRepo)
	res, err := testUseCase.GetTeamDetailsByID(uint(1))
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	mockRepo.AssertExpectations(t)
}
