package handler

import (
	"net/http"
	"screening-test/player"
	"screening-test/response"

	"github.com/gin-gonic/gin"
)

func PlayerRoute(playerUC player.PlayerUseCase, r *gin.RouterGroup) {
	uc := playerHandler{
		playerUC: playerUC,
	}

	v2 := r.Group("player")
	v2.GET("/", uc.GetAllPlayer)
	v2.GET("/:id", uc.GetDetailPlayer)
	v2.POST("/", uc.CreatePlayer)
	v2.PUT("/:id", uc.UpdatePlayer)
	v2.DELETE("/:id", uc.DeletePlayer)
}

type playerHandler struct {
	playerUC player.PlayerUseCase
}

func (playerHandler *playerHandler) GetAllPlayer(c *gin.Context) {
	result, pagination, err := playerHandler.playerUC.GetAllPlayerUC(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		Data:    result,
		Status:  "success",
		Message: "Get All Player",
		Meta:    pagination,
	})
}

func (playerHandler *playerHandler) GetDetailPlayer(c *gin.Context) {
	result, err := playerHandler.playerUC.GetDetailPlayerUC(c)
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

func (playerHandler *playerHandler) CreatePlayer(c *gin.Context) {
	err := playerHandler.playerUC.CreatePlayerUC(c)
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

func (playerHandler *playerHandler) UpdatePlayer(c *gin.Context) {
	err := playerHandler.playerUC.UpdatePlayerUC(c)
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

func (playerHandler *playerHandler) DeletePlayer(c *gin.Context) {
	err := playerHandler.playerUC.DeletePlayerUC(c)
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