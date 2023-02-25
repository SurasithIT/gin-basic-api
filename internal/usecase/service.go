package usecase

import (
	"github.com/surasithit/gin-basic-api/config"
	"github.com/surasithit/gin-basic-api/internal/repository"
	"github.com/surasithit/gin-basic-api/internal/usecase/player"
)

type Service struct {
	Config        *config.Config
	PlayerService *player.Service
}

func NewService(config *config.Config, repository *repository.Repository) *Service {
	return &Service{
		Config:        config,
		PlayerService: player.NewService(repository.PlayerRepository),
	}
}
