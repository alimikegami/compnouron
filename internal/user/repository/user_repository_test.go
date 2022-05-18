package repository

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alimikegami/compnouron/internal/user/entity"
	"github.com/alimikegami/compnouron/pkg/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
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
	userRepo := CreateNewUserRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectBegin()
	mockObj.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`name`,`email`,`phone_number`,`password`,`school_institution`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?)")).WithArgs("Alim Ikegami", "sdafsfa@gmail.com", "081111111111", "asdfasfas", "Udayana University", utils.AnyTime{}, utils.AnyTime{}).WillReturnResult(sqlmock.NewResult(1, 1))
	mockObj.ExpectCommit()

	userID, err := userRepo.CreateUser(entity.User{
		Name:              "Alim Ikegami",
		Email:             "sdafsfa@gmail.com",
		PhoneNumber:       "081111111111",
		Password:          "asdfasfas",
		SchoolInstitution: "Udayana University",
	})
	assert.NoError(t, err)
	assert.Equal(t, userID, uint(1))
}

func TestCreateUserUnexpectedDBError(t *testing.T) {
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
	userRepo := CreateNewUserRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectBegin()
	mockObj.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`name`,`email`,`phone_number`,`password`,`school_institution`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?)")).WithArgs("Alim Ikegami", "sdafsfa@gmail.com", "081111111111", "asdfasfas", "Udayana University", utils.AnyTime{}, utils.AnyTime{}).WillReturnError(errors.New("unexpected DB error"))
	mockObj.ExpectCommit()

	userID, err := userRepo.CreateUser(entity.User{
		Name:              "Alim Ikegami",
		Email:             "sdafsfa@gmail.com",
		PhoneNumber:       "081111111111",
		Password:          "asdfasfas",
		SchoolInstitution: "Udayana University",
	})
	assert.Error(t, err)
	assert.Equal(t, userID, uint(0))
}

func TestGetUserByEmail(t *testing.T) {
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
	userRepo := CreateNewUserRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectQuery("SELECT \\* FROM `users` WHERE email = \\? ORDER BY `users`\\.`id` LIMIT 1").WithArgs("alimikegami1@gmail.com").WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "phone_number", "password", "school_institution", "created_at", "updated_at"}).AddRow(1, "Alim Ikegami", "alimikegami1@gmail.com", "081239990127", "asdfasdfasdfsdf", "Udayana University", time.Now(), time.Now()))

	entity := userRepo.GetUserByEmail("alimikegami1@gmail.com")
	assert.NotEmpty(t, entity)
}

func TestGetUserByEmailNotFound(t *testing.T) {
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
	userRepo := CreateNewUserRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectQuery("SELECT \\* FROM `users` WHERE email = \\? ORDER BY `users`\\.`id` LIMIT 1").WithArgs("alimikegami11@gmail.com").WillReturnRows(sqlmock.NewRows(nil))
	entity := userRepo.GetUserByEmail("alimikegami1@gmail.com")
	assert.Empty(t, entity)
}

func TestAddUserSkills(t *testing.T) {
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
	userRepo := CreateNewUserRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectBegin()
	mockObj.ExpectExec(regexp.QuoteMeta("INSERT INTO `skills` (`name`,`user_id`,`created_at`,`updated_at`) VALUES (?,?,?,?)")).WithArgs("Node.Js", 1, utils.AnyTime{}, utils.AnyTime{}).WillReturnResult(sqlmock.NewResult(1, 1))
	mockObj.ExpectCommit()

	err = userRepo.AddUserSkills([]entity.Skill{
		{
			Name:   "Node.Js",
			UserID: 1,
		},
	})
	assert.NoError(t, err)
}

func TestAddUserSkillsErrorDB(t *testing.T) {
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
	userRepo := CreateNewUserRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectBegin()
	mockObj.ExpectExec(regexp.QuoteMeta("INSERT INTO `skills` (`name`,`user_id`,`created_at`,`updated_at`) VALUES (?,?,?,?)")).WithArgs("Node.Js", 1, utils.AnyTime{}, utils.AnyTime{}).WillReturnError(errors.New("error occured"))
	mockObj.ExpectCommit()

	err = userRepo.AddUserSkills([]entity.Skill{
		{
			Name:   "Node.Js",
			UserID: 1,
		},
	})
	assert.Error(t, err)
}
