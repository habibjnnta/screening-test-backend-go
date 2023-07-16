package team

import (
	migrtations "screening-test/migrations"
	"screening-test/models"

	"github.com/gin-gonic/gin"
)

type Team struct {
	ID       uint                 `json:"id"`
	NameTeam string               `json:"name_team"`
	Players  []migrtations.Player `gorm:"foreignKey:team_id;references:id"`
}

type TeamRepository interface {
	GetAllTeamRepo(pagination *models.Pagination) ([]Team, *models.Pagination, error)
	GetDetailTeamRepo(id int) (*Team, error)
	CreateTeamRepo(*Team) error
	UpdateTeamRepo(*Team) error
	DeleteTeamRepo(id int) error
}

type TeamUseCase interface {
	GetAllTeamUC(*gin.Context) ([]Team, *models.Pagination, error)
	GetDetailTeamUC(*gin.Context) (*Team, error)
	CreateTeamUC(*gin.Context) error
	UpdateTeamUC(*gin.Context) error
	DeleteTeamUC(*gin.Context) error
}
