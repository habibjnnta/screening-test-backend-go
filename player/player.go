package player

import (
	migrtations "screening-test/migrations"
	"screening-test/models"

	"github.com/gin-gonic/gin"
)

type Player struct {
	ID     uint             `json:"id"`
	Name   string           `json:"name"`
	TeamId uint             `json:"team_id"`
	Teams  migrtations.Team `gorm:"foreignKey:ID;references:team_id"`
}

type PlayerRepository interface {
	GetAllPlayerRepo(pagination *models.Pagination) ([]Player, *models.Pagination, error)
	GetDetailPlayerRepo(id int) (*Player, error)
	CreatePlayerRepo(*Player) error
	UpdatePlayerRepo(*Player) error
	DeletePlayerRepo(id int) error
}

type PlayerUseCase interface {
	GetAllPlayerUC(*gin.Context) ([]Player, *models.Pagination, error)
	GetDetailPlayerUC(*gin.Context) (*Player, error)
	CreatePlayerUC(*gin.Context) error
	UpdatePlayerUC(*gin.Context) error
	DeletePlayerUC(*gin.Context) error
}
