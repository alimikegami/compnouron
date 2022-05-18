package repository

import (
	"database/sql/driver"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alimikegami/compnouron/internal/recruitment/entity"
	"github.com/alimikegami/compnouron/pkg/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestCloseRecruitmentApplicationPeriod(t *testing.T) {
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
	recruitmentRepo := CreateNewRecruitmentRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectBegin()
	mockObj.ExpectExec(regexp.QuoteMeta("UPDATE")).WithArgs(2, utils.AnyTime{}, 1).WillReturnResult(driver.RowsAffected(1))
	mockObj.ExpectCommit()

	err = recruitmentRepo.CloseRecruitmentApplicationPeriod(1)
	assert.NoError(t, err)
}

func TestCloseRecruitmentApplicationPeriodNoRowsAffected(t *testing.T) {
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
	recruitmentRepo := CreateNewRecruitmentRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectBegin()
	mockObj.ExpectExec(regexp.QuoteMeta("UPDATE")).WithArgs(2, utils.AnyTime{}, 1).WillReturnError(errors.New("no rows affected"))
	mockObj.ExpectCommit()

	err = recruitmentRepo.CloseRecruitmentApplicationPeriod(1)
	assert.Error(t, err)
}

func TestOpenRecruitmentApplicationPeriod(t *testing.T) {
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
	recruitmentRepo := CreateNewRecruitmentRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectBegin()
	mockObj.ExpectExec(regexp.QuoteMeta("UPDATE")).WithArgs(1, utils.AnyTime{}, 1).WillReturnResult(driver.RowsAffected(1))
	mockObj.ExpectCommit()

	err = recruitmentRepo.OpenRecruitmentApplicationPeriod(1)
	assert.NoError(t, err)
}

func TestOpenRecruitmentApplicationPeriodNoRowsAffected(t *testing.T) {
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
	recruitmentRepo := CreateNewRecruitmentRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectBegin()
	mockObj.ExpectExec(regexp.QuoteMeta("UPDATE")).WithArgs(1, utils.AnyTime{}, 1).WillReturnError(errors.New("no rows affected"))
	mockObj.ExpectCommit()

	err = recruitmentRepo.OpenRecruitmentApplicationPeriod(1)
	assert.Error(t, err)
}

func TestDeleteRecruitmentByID(t *testing.T) {
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
	recruitmentRepo := CreateNewRecruitmentRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectBegin()
	mockObj.ExpectExec(regexp.QuoteMeta("DELETE")).WithArgs(1).WillReturnResult(driver.RowsAffected(1))
	mockObj.ExpectCommit()

	err = recruitmentRepo.DeleteRecruitmentByID(1)
	assert.NoError(t, err)
}

func TestDeleteRecruitmentByIDNoRowsAffected(t *testing.T) {
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
	recruitmentRepo := CreateNewRecruitmentRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectBegin()
	mockObj.ExpectExec(regexp.QuoteMeta("DELETE")).WithArgs(1).WillReturnError(errors.New("no rows affected"))
	mockObj.ExpectCommit()

	err = recruitmentRepo.DeleteRecruitmentByID(1)
	assert.Error(t, err)
}

func TestCreateRecruitment(t *testing.T) {
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
	recruitmentRepo := CreateNewRecruitmentRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectBegin()
	mockObj.ExpectExec(regexp.QuoteMeta("INSERT")).WithArgs("frontend engineer", "We need frontend engineer that can use React.Js", 1, 0, utils.AnyTime{}, utils.AnyTime{}).WillReturnResult(sqlmock.NewResult(2, 1))
	mockObj.ExpectCommit()

	err = recruitmentRepo.CreateRecruitment(entity.Recruitment{
		Role:                        "frontend engineer",
		Description:                 "We need frontend engineer that can use React.Js",
		TeamID:                      1,
		ApplicationAcceptanceStatus: 0,
	})
	assert.NoError(t, err)
}

func TestCreateRecruitmentUnexpectedDBError(t *testing.T) {
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
	recruitmentRepo := CreateNewRecruitmentRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectBegin()
	mockObj.ExpectExec(regexp.QuoteMeta("INSERT")).WithArgs("frontend engineer", "We need frontend engineer that can use React.Js", 1, 0, utils.AnyTime{}, utils.AnyTime{}).WillReturnError(errors.New("unexpected DB error"))
	mockObj.ExpectCommit()

	err = recruitmentRepo.CreateRecruitment(entity.Recruitment{
		Role:                        "frontend engineer",
		Description:                 "We need frontend engineer that can use React.Js",
		TeamID:                      1,
		ApplicationAcceptanceStatus: 0,
	})
	assert.Error(t, err)
}

func TestCreateRecruitmentApplication(t *testing.T) {
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
	recruitmentRepo := CreateNewRecruitmentRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectBegin()
	mockObj.ExpectExec(regexp.QuoteMeta("INSERT")).WithArgs(1, 1, 0, utils.AnyTime{}, utils.AnyTime{}).WillReturnResult(sqlmock.NewResult(2, 1))
	mockObj.ExpectCommit()

	err = recruitmentRepo.CreateRecruitmentApplication(entity.RecruitmentApplication{
		UserID:           1,
		RecruitmentID:    1,
		AcceptanceStatus: 0,
	})
	assert.NoError(t, err)
}

func TestCreateRecruitmentApplicationUnexpectedDBError(t *testing.T) {
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
	recruitmentRepo := CreateNewRecruitmentRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectBegin()
	mockObj.ExpectExec(regexp.QuoteMeta("INSERT")).WithArgs(1, 1, 0, utils.AnyTime{}, utils.AnyTime{}).WillReturnError(errors.New("unexpected DB error"))
	mockObj.ExpectCommit()

	err = recruitmentRepo.CreateRecruitmentApplication(entity.RecruitmentApplication{
		UserID:           1,
		RecruitmentID:    1,
		AcceptanceStatus: 0,
	})
	assert.Error(t, err)
}

func TestRejectRecruitmentApplication(t *testing.T) {
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
	recruitmentRepo := CreateNewRecruitmentRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectBegin()
	mockObj.ExpectExec(regexp.QuoteMeta("UPDATE")).WithArgs(2, utils.AnyTime{}, 1).WillReturnResult(driver.RowsAffected(1))
	mockObj.ExpectCommit()

	err = recruitmentRepo.RejectRecruitmentApplication(1)
	assert.NoError(t, err)
}

func TestRejectRecruitmentApplicationUnexpectedDBError(t *testing.T) {
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
	recruitmentRepo := CreateNewRecruitmentRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectBegin()
	mockObj.ExpectExec(regexp.QuoteMeta("UPDATE")).WithArgs(2, utils.AnyTime{}, 1).WillReturnError(errors.New("unexpected DB error"))
	mockObj.ExpectCommit()

	err = recruitmentRepo.RejectRecruitmentApplication(1)
	assert.Error(t, err)
}

func TestAcceptRecruitmentApplication(t *testing.T) {
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
	recruitmentRepo := CreateNewRecruitmentRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectBegin()
	mockObj.ExpectExec(regexp.QuoteMeta("UPDATE")).WithArgs(1, utils.AnyTime{}, 1).WillReturnResult(driver.RowsAffected(1))
	mockObj.ExpectCommit()

	err = recruitmentRepo.AcceptRecruitmentApplication(1)
	assert.NoError(t, err)
}

func TestAcceptRecruitmentApplicationUnexpectedDBError(t *testing.T) {
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
	recruitmentRepo := CreateNewRecruitmentRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectBegin()
	mockObj.ExpectExec(regexp.QuoteMeta("UPDATE")).WithArgs(1, utils.AnyTime{}, 1).WillReturnError(errors.New("unexpected database error"))
	mockObj.ExpectCommit()

	err = recruitmentRepo.AcceptRecruitmentApplication(1)
	assert.Error(t, err)
}

func TestGetRecruitmentByID(t *testing.T) {
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
	recruitmentRepo := CreateNewRecruitmentRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectQuery(regexp.QuoteMeta("SELECT `recruitments`.`id`,`recruitments`.`role`,`recruitments`.`description`,`recruitments`.`team_id`,`recruitments`.`application_acceptance_status`,`recruitments`.`created_at`,`recruitments`.`updated_at`,`Team`.`id` AS `Team__id`,`Team`.`name` AS `Team__name`,`Team`.`description` AS `Team__description`,`Team`.`capacity` AS `Team__capacity`,`Team`.`created_at` AS `Team__created_at`,`Team`.`updated_at` AS `Team__updated_at` FROM `recruitments` LEFT JOIN `teams` `Team` ON `recruitments`.`team_id` = `Team`.`id` WHERE recruitments.id = ? ORDER BY `recruitments`.`id` LIMIT 1")).WithArgs(uint(1)).WillReturnRows(sqlmock.NewRows([]string{"recruitments.id", "recruitments.role", "recruitments.description", "recruitments.team_id", "recruitments.application_acceptance_status", "recruitments.created_at", "recruitments.updated_at", "Team__id", "Team__name", "Team__description", "Team__Team__capacity", "Team__created_at", "Team__updated_at"}).AddRow(1, "Backend Engineer", "asdfasdf", uint(1), 0, time.Now(), time.Now(), uint(1), "Team 1", "Team hackahton", 4, time.Now(), time.Now()))
	entity, err := recruitmentRepo.GetRecruitmentByID(uint(1))
	assert.NotEmpty(t, entity)
	assert.NoError(t, err)
}

func TestGetRecruitmentByIDNoRows(t *testing.T) {
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
	recruitmentRepo := CreateNewRecruitmentRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectQuery(regexp.QuoteMeta("SELECT `recruitments`.`id`,`recruitments`.`role`,`recruitments`.`description`,`recruitments`.`team_id`,`recruitments`.`application_acceptance_status`,`recruitments`.`created_at`,`recruitments`.`updated_at`,`Team`.`id` AS `Team__id`,`Team`.`name` AS `Team__name`,`Team`.`description` AS `Team__description`,`Team`.`capacity` AS `Team__capacity`,`Team`.`created_at` AS `Team__created_at`,`Team`.`updated_at` AS `Team__updated_at` FROM `recruitments` LEFT JOIN `teams` `Team` ON `recruitments`.`team_id` = `Team`.`id` WHERE recruitments.id = ? ORDER BY `recruitments`.`id` LIMIT 1")).WithArgs(uint(1)).WillReturnRows(sqlmock.NewRows(nil))
	entity, err := recruitmentRepo.GetRecruitmentByID(uint(1))
	assert.Empty(t, entity)
	assert.Error(t, err)
}
