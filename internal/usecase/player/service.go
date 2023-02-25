package player

import (
	"github.com/surasithit/gin-basic-api/config"
	"github.com/surasithit/gin-basic-api/domain"
	"github.com/surasithit/gin-basic-api/internal/repository/player"
)

type Service struct {
	Config           *config.Config
	PlayerRepository *player.Repository
}

// Implement interface
var _ domain.PlayerUsecase = (*Service)(nil)

func NewService(config *config.Config, playerRepository *player.Repository) *Service {
	return &Service{
		Config:           config,
		PlayerRepository: playerRepository,
	}
}
