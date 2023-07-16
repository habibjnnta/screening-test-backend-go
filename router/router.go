package router

import (
	"context"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"screening-test/middleware"
	playerHandler "screening-test/player/handler"
	playerRepo "screening-test/player/repository"
	playerUC "screening-test/player/usecase"
	teamHandler "screening-test/team/handler"
	teamRepo "screening-test/team/repository"
	teamUC "screening-test/team/usecase"
)

type Handlers struct {
	Ctx context.Context
	DB  *gorm.DB
	R   *gin.Engine
}

func (h *Handlers) Routes() {
	// Repository
	TeamRepo := teamRepo.NewTeamRepository(h.DB)
	PlayerRepo := playerRepo.NewPlayerRepository(h.DB)

	// Usecase
	TeamUseCase := teamUC.NewTeamUseCase(TeamRepo)
	PlayerUseCase := playerUC.NewPlayerUseCase(PlayerRepo)

	// Add Middleware CORS
	middleware.Add(h.R, middleware.CORSMiddleware())

	// Handlers
	v1 := h.R.Group("api")
	teamHandler.TeamRoute(TeamUseCase, v1)
	playerHandler.PlayerRoute(PlayerUseCase, v1)
}
