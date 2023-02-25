package player

import (
	"github.com/surasithit/gin-basic-api/domain"
	"github.com/surasithit/gin-basic-api/internal/repository/player"
)

type Service struct {
	PlayerRepository *player.Repository
}

// Implement interface
var _ domain.PlayerUsecase = (*Service)(nil)

func NewService(playerRepository *player.Repository) *Service {
	return &Service{
		PlayerRepository: playerRepository,
	}
}
