package usecase

import (
	"screening-test/models"
	"screening-test/request"
	"screening-test/team"
	"strconv"

	"github.com/gin-gonic/gin"
)

// NewTeamUseCase return TeamUseCase
func NewTeamUseCase(teamRepo team.TeamRepository) *TeamUseCase {
	return &TeamUseCase{
		teamRepo: teamRepo,
	}
}

type TeamUseCase struct {
	teamRepo team.TeamRepository
}

func (teamUC *TeamUseCase) GetAllTeamUC(c *gin.Context) ([]team.Team, *models.Pagination, error) {
	pagination, err := request.Paginate(c)
	if err != nil {
		return nil, nil, err
	}

	result, pagination, err := teamUC.teamRepo.GetAllTeamRepo(pagination)
	if err != nil {
		return nil, nil, err
	}

	return result, pagination, nil

}

func (teamUC *TeamUseCase) GetDetailTeamUC(c *gin.Context) (*team.Team, error) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		return nil, err
	}

	result, err := teamUC.teamRepo.GetDetailTeamRepo(ID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (teamUC *TeamUseCase) CreateTeamUC(c *gin.Context) error {
	var result team.Team
	err := c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = teamUC.teamRepo.CreateTeamRepo(&result)
	if err != nil {
		return err
	}

	return nil
}

func (teamUC *TeamUseCase) UpdateTeamUC(c *gin.Context) error {
	var result team.Team
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		return err
	}

	err = c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	result.ID = uint(ID)

	err = teamUC.teamRepo.UpdateTeamRepo(&result)
	if err != nil {
		return err
	}

	return nil
}

func (teamUC *TeamUseCase) DeleteTeamUC(c *gin.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		return err
	}

	err = teamUC.teamRepo.DeleteTeamRepo(ID)
	if err != nil {
		return err
	}

	return nil
}
