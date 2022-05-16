package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/alimikegami/compnouron/cmd/app/docs"
	"github.com/alimikegami/compnouron/db/migration"
	competitionController "github.com/alimikegami/compnouron/internal/competition/controller"
	competitionRepository "github.com/alimikegami/compnouron/internal/competition/repository"
	competitionUseCase "github.com/alimikegami/compnouron/internal/competition/usecase"
	recruitmentController "github.com/alimikegami/compnouron/internal/recruitment/controller"
	recruitmentRepository "github.com/alimikegami/compnouron/internal/recruitment/repository"
	recruitmentUseCase "github.com/alimikegami/compnouron/internal/recruitment/usecase"
	teamController "github.com/alimikegami/compnouron/internal/team/controller"
	teamRepository "github.com/alimikegami/compnouron/internal/team/repository"
	teamUseCase "github.com/alimikegami/compnouron/internal/team/usecase"
	"github.com/alimikegami/compnouron/internal/user/controller"
	"github.com/alimikegami/compnouron/internal/user/repository"
	"github.com/alimikegami/compnouron/internal/user/usecase"
	"github.com/alimikegami/compnouron/pkg/utils"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initializeDatabaseConnection() (*gorm.DB, error) {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	address := os.Getenv("DB_ADDRESS")
	databaseName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, address, port, databaseName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

func main() {
	e := echo.New()

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := initializeDatabaseConnection()
	if err != nil {
		fmt.Println("Connection to the database has not been established")
	}

	config := middleware.JWTConfig{
		Claims:     &utils.JwtCustomClaims{},
		SigningKey: []byte(os.Getenv("SIGNING_KEY")),
	}

	migration.Migrate(db)

	cr := competitionRepository.CreateNewCompetitionRepository(db)
	cuc := competitionUseCase.CreateNewCompetitionUseCase(cr)
	cc := competitionController.CreateNewCompetitionController(e, cuc)
	cc.InitializeCompetitionRoute(config)

	userRepository := repository.CreateNewUserRepository(db)
	userUseCase := usecase.CreateNewUserUseCase(userRepository)
	userController := controller.CreateNewUserController(e, userUseCase)
	userController.InitializeUserRoute(config)

	tr := teamRepository.CreateNewTeamRepository(db)
	tuc := teamUseCase.CreateNewTeamUseCase(tr)
	tc := teamController.CreateNewTeamController(e, tuc)
	tc.InitializeTeamRoute(config)

	rr := recruitmentRepository.CreateNewRecruitmentRepository(db)
	ruc := recruitmentUseCase.CreateNewRecruitmentUseCase(rr, tr)
	rc := recruitmentController.CreateNewRecruitmentController(e, ruc)
	rc.InitializeRecruitmentRoute(config)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
