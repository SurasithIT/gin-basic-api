package players

import (
	"github.com/gin-gonic/gin"
	"github.com/surasithit/gin-basic-api/internal/players/usecase"
)

type PlayerControllerInterface interface {
	GetPlayers(ctx *gin.Context) error
	GetPlayerById(ctx *gin.Context) error
	CreatePlayer(ctx *gin.Context) error
	UpdatePlayer(ctx *gin.Context) error
	DeletePlayer(ctx *gin.Context) error
}

type Controller struct {
	PlayerService *usecase.Service
}

var _ PlayerControllerInterface = (*Controller)(nil)

func NewController(playerService *usecase.Service) *Controller {
	return &Controller{
		PlayerService: playerService,
	}
}

// CreatePlayer implements PlayerControllerInterface
func (*Controller) CreatePlayer(ctx *gin.Context) error {
	panic("unimplemented")
}

// DeletePlayer implements PlayerControllerInterface
func (*Controller) DeletePlayer(ctx *gin.Context) error {
	panic("unimplemented")
}

// GetPlayerById implements PlayerControllerInterface
func (*Controller) GetPlayerById(ctx *gin.Context) error {
	panic("unimplemented")
}

// GetPlayers implements PlayerControllerInterface
func (*Controller) GetPlayers(ctx *gin.Context) error {
	panic("unimplemented")
}

// UpdatePlayer implements PlayerControllerInterface
func (*Controller) UpdatePlayer(ctx *gin.Context) error {
	panic("unimplemented")
}
