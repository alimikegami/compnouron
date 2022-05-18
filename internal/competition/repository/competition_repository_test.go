package repository

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alimikegami/compnouron/internal/competition/entity"
	"github.com/alimikegami/compnouron/pkg/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestCreateCompetition(t *testing.T) {
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
	compRepo := CreateNewCompetitionRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectBegin()
	mockObj.ExpectExec(regexp.QuoteMeta("INSERT INTO `competitions` (`name`,`description`,`contact_person`,`is_team`,`is_the_same_institution`,`registration_period_status`,`team_capacity`,`level`,`created_at`,`updated_at`,`user_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?)")).WithArgs("Technoscape Hackathon 2022", "Hackathon dengan peserta sebanyak 4 orang per tim", "081239990128", 1, 1, 0, 4, "University Student", utils.AnyTime{}, utils.AnyTime{}, 1).WillReturnResult(sqlmock.NewResult(1, 1))
	mockObj.ExpectCommit()

	err = compRepo.CreateCompetition(&entity.Competition{
		Name:                 "Technoscape Hackathon 2022",
		Description:          "Hackathon dengan peserta sebanyak 4 orang per tim",
		ContactPerson:        "081239990128",
		IsTheSameInstitution: 1,
		IsTeam:               1,
		TeamCapacity:         4,
		Level:                "University Student",
		UserID:               1,
	})
	assert.NoError(t, err)
}

func TestCreateCompetitionUnexpectedDBError(t *testing.T) {
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
	compRepo := CreateNewCompetitionRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectBegin()
	mockObj.ExpectExec(regexp.QuoteMeta("INSERT INTO `competitions` (`name`,`description`,`contact_person`,`is_team`,`is_the_same_institution`,`registration_period_status`,`team_capacity`,`level`,`created_at`,`updated_at`,`user_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?)")).WithArgs("Technoscape Hackathon 2022", "Hackathon dengan peserta sebanyak 4 orang per tim", "081239990128", 1, 1, 0, 4, "University Student", utils.AnyTime{}, utils.AnyTime{}, 1).WillReturnError(errors.New("unexpected DB error"))
	mockObj.ExpectCommit()

	err = compRepo.CreateCompetition(&entity.Competition{
		Name:                 "Technoscape Hackathon 2022",
		Description:          "Hackathon dengan peserta sebanyak 4 orang per tim",
		ContactPerson:        "081239990128",
		IsTheSameInstitution: 1,
		IsTeam:               1,
		TeamCapacity:         4,
		Level:                "University Student",
		UserID:               1,
	})
	assert.Error(t, err)
}

func TestDeleteCompetition(t *testing.T) {
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
	compRepo := CreateNewCompetitionRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectBegin()
	mockObj.ExpectExec(regexp.QuoteMeta("DELETE")).WithArgs(uint(1)).WillReturnResult(sqlmock.NewResult(0, 1))
	mockObj.ExpectCommit()

	err = compRepo.DeleteCompetition(uint(1))
	assert.NoError(t, err)
}

func TestDeleteCompetitionNoRowsAffected(t *testing.T) {
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
	compRepo := CreateNewCompetitionRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectBegin()
	mockObj.ExpectExec(regexp.QuoteMeta("DELETE")).WithArgs(uint(1111)).WillReturnResult(sqlmock.NewResult(0, 0))
	mockObj.ExpectCommit()

	err = compRepo.DeleteCompetition(uint(1111))
	assert.NoError(t, err)
}

func TestGetCompetitionByID(t *testing.T) {
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
	compRepo := CreateNewCompetitionRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `competitions` WHERE `competitions`.`id` = ? ORDER BY `competitions`.`id` LIMIT 1")).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "contact_person", "is_team", "is_the_same_institution", "registration_period_status", "team_capacity", "level", "created_at", "updated_at", "user_id"}).AddRow(1, "Technoscape", "Hackathon", "081239990129", 1, 1, 0, 4, "University Student", time.Now(), time.Now(), uint(1)))

	entity, err := compRepo.GetCompetitionByID(1)
	assert.NoError(t, err)
	assert.NotEmpty(t, entity)
}

func TestGetCompetitionByIDNotFound(t *testing.T) {
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
	compRepo := CreateNewCompetitionRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `competitions` WHERE `competitions`.`id` = ? ORDER BY `competitions`.`id` LIMIT 1")).WithArgs(1).WillReturnRows(sqlmock.NewRows(nil))

	entity, err := compRepo.GetCompetitionByID(1)
	assert.Error(t, gorm.ErrRecordNotFound)
	assert.Empty(t, entity)
}

func TestUpdateCompetition(t *testing.T) {
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
	compRepo := CreateNewCompetitionRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectBegin()
	mockObj.ExpectExec(regexp.QuoteMeta("UPDATE `competitions` SET `id`=?,`name`=?,`description`=?,`contact_person`=?,`is_team`=?,`is_the_same_institution`=?,`team_capacity`=?,`level`=?,`updated_at`=? WHERE id = ? AND `id` = ?")).WithArgs(1, "Technoscape Hackathon 2022", "Hackathon dengan peserta sebanyak 4 orang per tim", "081239990128", 1, 1, 4, "University Student", utils.AnyTime{}, 1, 1).WillReturnResult(sqlmock.NewResult(0, 1))
	mockObj.ExpectCommit()

	err = compRepo.UpdateCompetition(entity.Competition{
		ID:                   1,
		Name:                 "Technoscape Hackathon 2022",
		Description:          "Hackathon dengan peserta sebanyak 4 orang per tim",
		ContactPerson:        "081239990128",
		IsTeam:               1,
		IsTheSameInstitution: 1,
		TeamCapacity:         4,
		Level:                "University Student",
	})
	assert.NoError(t, err)
}

func TestUpdateCompetitionNoAffectedRows(t *testing.T) {
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
	compRepo := CreateNewCompetitionRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectBegin()
	mockObj.ExpectExec(regexp.QuoteMeta("UPDATE `competitions` SET `id`=?,`name`=?,`description`=?,`contact_person`=?,`is_team`=?,`is_the_same_institution`=?,`team_capacity`=?,`level`=?,`updated_at`=? WHERE id = ? AND `id` = ?")).WithArgs(1, "Technoscape Hackathon 2022", "Hackathon dengan peserta sebanyak 4 orang per tim", "081239990128", 1, 1, 4, "University Student", utils.AnyTime{}, 1, 1).WillReturnError(errors.New("no affected rows"))
	mockObj.ExpectCommit()

	err = compRepo.UpdateCompetition(entity.Competition{
		ID:            1,
		Name:          "Technoscape Hackathon 2022",
		Description:   "Hackathon dengan peserta sebanyak 4 orang per tim",
		ContactPerson: "081239990128",
		IsTeam:        1,
		TeamCapacity:  4,
		Level:         "University Student",
	})
	assert.Error(t, err)
}

// func TestRegister(t *testing.T) {

// }
