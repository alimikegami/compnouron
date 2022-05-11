package migration

import (
	compEntity "github.com/alimikegami/compnouron/internal/competition/entity"
	recruitmentEntity "github.com/alimikegami/compnouron/internal/recruitment/entity"
	teamEntity "github.com/alimikegami/compnouron/internal/team/entity"
	"github.com/alimikegami/compnouron/internal/user/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	if (!db.Migrator().HasTable(&entity.User{})) {
		db.Migrator().CreateTable(&entity.User{})
	}

	if !db.Migrator().HasTable(&compEntity.Competition{}) {
		db.Migrator().CreateTable(&compEntity.Competition{})
	}

	if !db.Migrator().HasTable(&teamEntity.Team{}) {
		db.Migrator().CreateTable(&teamEntity.Team{})
	}

	if !db.Migrator().HasTable(&recruitmentEntity.Recruitment{}) {
		db.Migrator().CreateTable(&recruitmentEntity.Recruitment{})
	}

	if !db.Migrator().HasTable(&recruitmentEntity.RecruitmentApplication{}) {
		db.Migrator().CreateTable(&recruitmentEntity.RecruitmentApplication{})
	}

	if !db.Migrator().HasTable(&teamEntity.TeamMember{}) {
		db.Migrator().CreateTable(&teamEntity.TeamMember{})
	}
}
