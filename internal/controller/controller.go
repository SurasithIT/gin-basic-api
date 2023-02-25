package controller

import (
	"github.com/surasithit/gin-basic-api/internal/controller/player"
	"github.com/surasithit/gin-basic-api/internal/usecase"
)

type Controller struct {
	PlayerController *player.PlayerController
}

func NewController(usecase *usecase.Service) *Controller {
	return &Controller{
		PlayerController: player.NewController(usecase.PlayerService),
	}
}
