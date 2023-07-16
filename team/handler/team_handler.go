package handler

import (
	"net/http"
	"screening-test/team"
	"screening-test/response"

	"github.com/gin-gonic/gin"
)

func TeamRoute(teamUC team.TeamUseCase, r *gin.RouterGroup) {
	uc := teamHandler{
		teamUC: teamUC,
	}

	v2 := r.Group("team")
	v2.GET("/", uc.GetAllTeam)
	v2.GET("/:id", uc.GetDetailTeam)
	v2.POST("/", uc.CreateTeam)
	v2.PUT("/:id", uc.UpdateTeam)
	v2.DELETE("/:id", uc.DeleteTeam)
}

type teamHandler struct {
	teamUC team.TeamUseCase
}

func (teamHandler *teamHandler) GetAllTeam(c *gin.Context) {
	result, pagination, err := teamHandler.teamUC.GetAllTeamUC(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		Data:    result,
		Status:  "success",
		Message: "Get All Team",
		Meta:    pagination,
	})
}

func (teamHandler *teamHandler) GetDetailTeam(c *gin.Context) {
	result, err := teamHandler.teamUC.GetDetailTeamUC(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		Data:    result,
		Status:  "success",
		Message: "success",
	})
}

func (teamHandler *teamHandler) CreateTeam(c *gin.Context) {
	err := teamHandler.teamUC.CreateTeamUC(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, response.Response{
		Status:  "success",
		Message: "success",
	})
}

func (teamHandler *teamHandler) UpdateTeam(c *gin.Context) {
	err := teamHandler.teamUC.UpdateTeamUC(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusAccepted, response.Response{
		Status:  "success",
		Message: "success",
	})
}

func (teamHandler *teamHandler) DeleteTeam(c *gin.Context) {
	err := teamHandler.teamUC.DeleteTeamUC(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusAccepted, response.Response{
		Status:  "success",
		Message: "success",
	})
}