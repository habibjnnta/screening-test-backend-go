package repository

import (
	"log"
	"screening-test/models"
	"screening-test/paginator"
	"screening-test/player"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PlayerRepository struct {
	db *gorm.DB
}

func NewPlayerRepository(db *gorm.DB) *PlayerRepository {
	return &PlayerRepository{
		db: db,
	}
}

func (playerRepo *PlayerRepository) GetAllPlayerRepo(pagination *models.Pagination) ([]player.Player, *models.Pagination, error) {
	var result []player.Player

	log.Println(pagination.Limit, pagination.Offset)
	queryBuilder := playerRepo.db.Limit(pagination.Limit).Offset(pagination.Offset)
	data := queryBuilder.Preload(clause.Associations).Find(&result)
	// data := playerRepo.db.Preload(clause.Associations).Find(&result).Limit(pagination.Limit).Offset(pagination.Offset)
	if data.Error != nil {
		return nil, nil, data.Error
	}

	count := data.RowsAffected
	pagination.Count = int(count)

	pagination = paginator.Paging(pagination)

	return result, pagination, nil
}

func (playerRepo *PlayerRepository) GetDetailPlayerRepo(id int) (*player.Player, error) {
	err := playerRepo.db.First(&player.Player{}, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	var result *player.Player
	playerRepo.db.Preload(clause.Associations).Where("id = ?", id).Find(&player.Player{}).Scan(&result)

	return result, nil
}

func (playerRepo *PlayerRepository) CreatePlayerRepo(data *player.Player) error {
	err := playerRepo.db.Create(data).Error
	if err != nil {
		return err
	}

	return nil
}

func (playerRepo *PlayerRepository) UpdatePlayerRepo(data *player.Player) error {
	err := playerRepo.db.First(&player.Player{}, "id = ?", data.ID).Error
	if err != nil {
		return err
	}

	err = playerRepo.db.Model(&player.Player{}).Where("id = ?", data.ID).Updates(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func (playerRepo *PlayerRepository) DeletePlayerRepo(id int) error {
	err := playerRepo.db.First(&player.Player{}, "id = ?", id).Error
	if err != nil {
		return err
	}

	err = playerRepo.db.Delete(&player.Player{}, playerRepo.db.Where("id = ?", id)).Error
	if err != nil {
		return err
	}

	return nil
}
