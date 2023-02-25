package usecase

import (
	"github.com/surasithit/gin-basic-api/config"
	"github.com/surasithit/gin-basic-api/internal/helper/utils"
	"github.com/surasithit/gin-basic-api/internal/repository"
	"github.com/surasithit/gin-basic-api/internal/usecase/player"
)

type Service struct {
	Config        *config.Config
	Utils         *utils.UtilService
	PlayerService *player.Service
}

func NewService(config *config.Config, repository *repository.Repository, utils *utils.UtilService) *Service {
	return &Service{
		Config:        config,
		Utils:         utils,
		PlayerService: player.NewService(config, repository.PlayerRepository),
	}
}
