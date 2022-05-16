package repository

import (
	"database/sql/driver"
	"regexp"
	"testing"

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

// func TestUpdateRecruitment(t *testing.T) {

// }

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

// func TestGetRecruitmentByID(t *testing.T) {
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
// 	recruitmentRepo := CreateNewRecruitmentRepository(db)

// 	defer mockedDB.Close()

// 	res := sqlmock.NewRows([]string{})

// 	mockObj.ExpectBegin()
// 	mockObj.ExpectExec(regexp.QuoteMeta("UPDATE")).WithArgs(1, utils.AnyTime{}, 1).WillReturnResult(driver.RowsAffected(1))
// 	mockObj.ExpectCommit()

// 	err = recruitmentRepo.AcceptRecruitmentApplication(1)
// 	assert.NoError(t, err)
// }
