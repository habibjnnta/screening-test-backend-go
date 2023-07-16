package usecase

import (
	"screening-test/models"
	"screening-test/player"
	"screening-test/request"
	"strconv"

	"github.com/gin-gonic/gin"
)

// NewPlayerUseCase return PlayerUseCase
func NewPlayerUseCase(playerRepo player.PlayerRepository) *PlayerUseCase {
	return &PlayerUseCase{
		playerRepo: playerRepo,
	}
}

type PlayerUseCase struct {
	playerRepo player.PlayerRepository
}

func (playerUC *PlayerUseCase) GetAllPlayerUC(c *gin.Context) ([]player.Player, *models.Pagination, error) {
	pagination, err := request.Paginate(c)
	if err != nil {
		return nil, nil, err
	}

	result, pagination, err := playerUC.playerRepo.GetAllPlayerRepo(pagination)
	if err != nil {
		return nil, nil, err
	}

	return result, pagination, nil

}

func (playerUC *PlayerUseCase) GetDetailPlayerUC(c *gin.Context) (*player.Player, error) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		return nil, err
	}

	result, err := playerUC.playerRepo.GetDetailPlayerRepo(ID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (playerUC *PlayerUseCase) CreatePlayerUC(c *gin.Context) error {
	var result player.Player
	err := c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = playerUC.playerRepo.CreatePlayerRepo(&result)
	if err != nil {
		return err
	}

	return nil
}

func (playerUC *PlayerUseCase) UpdatePlayerUC(c *gin.Context) error {
	var result player.Player
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		return err
	}

	err = c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	result.ID = uint(ID)

	err = playerUC.playerRepo.UpdatePlayerRepo(&result)
	if err != nil {
		return err
	}

	return nil
}

func (playerUC *PlayerUseCase) DeletePlayerUC(c *gin.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		return err
	}

	err = playerUC.playerRepo.DeletePlayerRepo(ID)
	if err != nil {
		return err
	}

	return nil
}
