package repository

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alimikegami/compnouron/internal/team/entity"
	"github.com/alimikegami/compnouron/pkg/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestCreateTeam(t *testing.T) {
	mockedDB, mockObj, err := sqlmock.New()
	db, err := gorm.Open(mysql.Dialector{
		&mysql.Config{
			Conn:                      mockedDB,
			SkipInitializeWithVersion: true,
		},
	}, &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	teamRepo := CreateNewTeamRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectBegin()
	mockObj.ExpectExec(regexp.QuoteMeta("INSERT")).WithArgs("Team 1", "Software engineering team for Technoscape Hackathon 2022", 4, utils.AnyTime{}, utils.AnyTime{}).WillReturnResult(sqlmock.NewResult(2, 1))
	mockObj.ExpectCommit()

	team, err := teamRepo.CreateTeam(entity.Team{
		Name:        "Team 1",
		Description: "Software engineering team for Technoscape Hackathon 2022",
		Capacity:    4,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, team)
}

func TestAddTeamMember(t *testing.T) {
	mockedDB, mockObj, err := sqlmock.New()
	db, err := gorm.Open(mysql.Dialector{
		&mysql.Config{
			Conn:                      mockedDB,
			SkipInitializeWithVersion: true,
		},
	}, &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	teamRepo := CreateNewTeamRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectBegin()
	mockObj.ExpectExec(regexp.QuoteMeta("INSERT INTO `team_members` (`team_id`,`user_id`,`is_leader`,`created_at`,`updated_at`) VALUES (?,?,?,?,?)")).WithArgs(1, 1, 1, utils.AnyTime{}, utils.AnyTime{}).WillReturnResult(sqlmock.NewResult(2, 1))
	mockObj.ExpectCommit()

	err = teamRepo.AddTeamMember(1, 1, 1)
	assert.NoError(t, err)
}

func TestDeleteTeam(t *testing.T) {
	mockedDB, mockObj, err := sqlmock.New()
	db, err := gorm.Open(mysql.Dialector{
		&mysql.Config{
			Conn:                      mockedDB,
			SkipInitializeWithVersion: true,
		},
	}, &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	teamRepo := CreateNewTeamRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectBegin()
	mockObj.ExpectExec(regexp.QuoteMeta("DELETE")).WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))
	mockObj.ExpectCommit()

	err = teamRepo.DeleteTeam(1)
	assert.NoError(t, err)
}

func TestDeleteTeamNoRowsAffected(t *testing.T) {
	mockedDB, mockObj, err := sqlmock.New()
	db, err := gorm.Open(mysql.Dialector{
		&mysql.Config{
			Conn:                      mockedDB,
			SkipInitializeWithVersion: true,
		},
	}, &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	teamRepo := CreateNewTeamRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectBegin()
	mockObj.ExpectExec(regexp.QuoteMeta("DELETE")).WithArgs(99).WillReturnResult(sqlmock.NewResult(0, 0))
	mockObj.ExpectCommit()

	err = teamRepo.DeleteTeam(99)
	assert.Error(t, err)
}

// func TestDeleteTeamUnexpectedDatabaseError(t *testing.T) {
// 	mockedDB, mockObj, err := sqlmock.New()
// 	db, err := gorm.Open(mysql.Dialector{
// 		&mysql.Config{
// 			Conn:                      mockedDB,
// 			SkipInitializeWithVersion: true,
// 		},
// 	}, &gorm.Config{})
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	teamRepo := CreateNewTeamRepository(db)

// 	defer mockedDB.Close()

// 	mockObj.ExpectBegin()
// 	mockObj.ExpectExec(regexp.QuoteMeta("DELETE")).WithArgs(99).WillReturnResult(sqlmock.NewErrorResult(driver.ErrBadConn))
// 	mockObj.ExpectCommit()

// 	err = teamRepo.DeleteTeam(99)
// 	assert.NoError(t, err)
// }

func TestUpdateTeam(t *testing.T) {
	mockedDB, mockObj, err := sqlmock.New()
	db, err := gorm.Open(mysql.Dialector{
		&mysql.Config{
			Conn:                      mockedDB,
			SkipInitializeWithVersion: true,
		},
	}, &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	teamRepo := CreateNewTeamRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectBegin()
	mockObj.ExpectExec(regexp.QuoteMeta("UPDATE")).WithArgs(1, "Team 2", "Team Hackathon Technoscape 2023", 5, utils.AnyTime{}, 1).WillReturnResult(sqlmock.NewResult(0, 1))
	mockObj.ExpectCommit()

	err = teamRepo.UpdateTeam(entity.Team{
		ID:          1,
		Name:        "Team 2",
		Description: "Team Hackathon Technoscape 2023",
		Capacity:    5,
	})
	assert.NoError(t, err)
}

func TestUpdateTeamNoRowsAffected(t *testing.T) {
	mockedDB, mockObj, err := sqlmock.New()
	db, err := gorm.Open(mysql.Dialector{
		&mysql.Config{
			Conn:                      mockedDB,
			SkipInitializeWithVersion: true,
		},
	}, &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	teamRepo := CreateNewTeamRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectBegin()
	mockObj.ExpectExec(regexp.QuoteMeta("UPDATE")).WithArgs(1, "Team 1000", "Team Hackathon Technoscape 2023", 5, utils.AnyTime{}, 1).WillReturnResult(sqlmock.NewResult(0, 0))
	mockObj.ExpectCommit()

	err = teamRepo.UpdateTeam(entity.Team{
		ID:          1000,
		Name:        "Team 1000",
		Description: "Team Hackathon Technoscape 2023",
		Capacity:    5,
	})
	assert.Error(t, err)
}

func TestGetTeamByID(t *testing.T) {
	mockedDB, mockObj, err := sqlmock.New()
	db, err := gorm.Open(mysql.Dialector{
		&mysql.Config{
			Conn:                      mockedDB,
			SkipInitializeWithVersion: true,
		},
	}, &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	teamRepo := CreateNewTeamRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectQuery("SELECT \\* FROM `teams` WHERE `teams`.`id` = \\? ORDER BY `teams`\\.`id` LIMIT 1").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "capacity", "created_at", "updated_at"}).AddRow(1, "Team 1", "Team Hackathon Technoscape 2023", 4, time.Now(), time.Now()))

	entity, err := teamRepo.GetTeamByID(1)
	assert.NoError(t, err)
	assert.NotEmpty(t, entity)
}

func TestGetTeamByIDNoRows(t *testing.T) {
	mockedDB, mockObj, err := sqlmock.New()
	db, err := gorm.Open(mysql.Dialector{
		&mysql.Config{
			Conn:                      mockedDB,
			SkipInitializeWithVersion: true,
		},
	}, &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	teamRepo := CreateNewTeamRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectQuery("SELECT \\* FROM `teams` WHERE `teams`.`id` = \\? ORDER BY `teams`\\.`id` LIMIT 1").WithArgs(1).WillReturnRows(sqlmock.NewRows(nil))

	entity, err := teamRepo.GetTeamByID(1)
	assert.Error(t, err)
	assert.Empty(t, entity)
}

func TestGetTeamsByUserID(t *testing.T) {
	mockedDB, mockObj, err := sqlmock.New()
	db, err := gorm.Open(mysql.Dialector{
		&mysql.Config{
			Conn:                      mockedDB,
			SkipInitializeWithVersion: true,
		},
	}, &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	teamRepo := CreateNewTeamRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectQuery("SELECT `teams`.`id`,`teams`.`name`,`teams`.`description`,`teams`.`capacity`,`teams`.`created_at`,`teams`.`updated_at` FROM `teams` JOIN team_members ON team_members.team_id = teams.id WHERE team_members.user_id = \\?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"teams.id", "teams.name", "teams.description", "teams.capacity", "teams.created_at", "teams.updated_at"}).AddRow(1, "Team 1", "Team Hackathon Technoscape 2023", 4, time.Now(), time.Now()))

	entity, err := teamRepo.GetTeamsByUserID(1)
	assert.NoError(t, err)
	assert.NotNil(t, entity)
}

func TestGetTeamsByUserIDNoRows(t *testing.T) {
	mockedDB, mockObj, err := sqlmock.New()
	db, err := gorm.Open(mysql.Dialector{
		&mysql.Config{
			Conn:                      mockedDB,
			SkipInitializeWithVersion: true,
		},
	}, &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	teamRepo := CreateNewTeamRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectQuery("SELECT `teams`.`id`,`teams`.`name`,`teams`.`description`,`teams`.`capacity`,`teams`.`created_at`,`teams`.`updated_at` FROM `teams` JOIN team_members ON team_members.team_id = teams.id WHERE team_members.user_id = \\?").WithArgs(1).WillReturnRows(sqlmock.NewRows(nil))

	entity, err := teamRepo.GetTeamsByUserID(1)
	assert.NoError(t, err)
	assert.Len(t, entity, 0)
}
