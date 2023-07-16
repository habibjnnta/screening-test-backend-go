package repository

import (
	"errors"
	"log"
	"screening-test/models"
	"screening-test/paginator"
	"screening-test/player"
	"screening-test/team"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TeamRepository struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) *TeamRepository {
	return &TeamRepository{
		db: db,
	}
}

func (teamRepo *TeamRepository) GetAllTeamRepo(pagination *models.Pagination) ([]team.Team, *models.Pagination, error) {
	var result []team.Team

	log.Println(pagination.Limit, pagination.Offset)
	queryBuilder := teamRepo.db.Limit(pagination.Limit).Offset(pagination.Offset)
	data := queryBuilder.Preload(clause.Associations).Find(&result)
	// data := teamRepo.db.Preload(clause.Associations).Find(&result).Limit(pagination.Limit).Offset(pagination.Offset)
	if data.Error != nil {
		return nil, nil, data.Error
	}

	count := data.RowsAffected
	pagination.Count = int(count)

	pagination = paginator.Paging(pagination)

	return result, pagination, nil
}

func (teamRepo *TeamRepository) GetDetailTeamRepo(id int) (*team.Team, error) {
	err := teamRepo.db.First(&team.Team{}, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	var result *team.Team
	teamRepo.db.Preload(clause.Associations).Where("id = ?", id).Find(&team.Team{}).Scan(&result)

	return result, nil
}

func (teamRepo *TeamRepository) CreateTeamRepo(data *team.Team) error {
	err := teamRepo.db.Create(data).Error
	if err != nil {
		return err
	}

	return nil
}

func (teamRepo *TeamRepository) UpdateTeamRepo(data *team.Team) error {
	err := teamRepo.db.First(&team.Team{}, "id = ?", data.ID).Error
	if err != nil {
		return err
	}

	err = teamRepo.db.Model(&team.Team{}).Where("id = ?", data.ID).Updates(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func (teamRepo *TeamRepository) DeleteTeamRepo(id int) error {
	err := teamRepo.db.First(&team.Team{}, "id = ?", id).Error
	if err != nil {
		return err
	}

	var result []player.Player

	checking := teamRepo.db.Model(&result).Where("id = ?", id).RowsAffected

	if checking > 0 {
		return errors.New("team not be deleted because have the player in the team")
	}

	err = teamRepo.db.Delete(&team.Team{}, teamRepo.db.Where("id = ?", id)).Error
	if err != nil {
		return err
	}

	return nil
}
