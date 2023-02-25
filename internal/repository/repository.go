package repository

import (
	"github.com/surasithit/gin-basic-api/config"
	"github.com/surasithit/gin-basic-api/internal/helper/utils"
	"github.com/surasithit/gin-basic-api/internal/repository/player"
)

type Repository struct {
	Utils            *utils.UtilService
	Config           *config.Config
	PlayerRepository *player.Repository
}

func NewRepository(config *config.Config, utils *utils.UtilService) *Repository {
	return &Repository{
		Config:           config,
		Utils:            utils,
		PlayerRepository: player.NewRepository(),
	}
}
