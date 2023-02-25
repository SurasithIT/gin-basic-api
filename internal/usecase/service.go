package usecase

import (
	"github.com/surasithit/gin-basic-api/internal/repository"
	"github.com/surasithit/gin-basic-api/internal/usecase/player"
)

type Service struct {
	PlayerService *player.Service
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		PlayerService: player.NewService(repository.PlayerRepository),
	}
}
