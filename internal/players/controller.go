package players

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/surasithit/gin-basic-api/internal/players/dto"
	"github.com/surasithit/gin-basic-api/internal/players/usecase"
)

type PlayerControllerInterface interface {
	GetPlayers(ctx *gin.Context)
	GetPlayerById(ctx *gin.Context)
	CreatePlayer(ctx *gin.Context)
	UpdatePlayer(ctx *gin.Context)
	DeletePlayer(ctx *gin.Context)
}

type Controller struct {
	PlayerService *usecase.Service
}

// Implement interface
var _ PlayerControllerInterface = (*Controller)(nil)

func NewController(playerService *usecase.Service) *Controller {
	return &Controller{
		PlayerService: playerService,
	}
}

// CreatePlayer implements PlayerControllerInterface
func (c *Controller) CreatePlayer(ctx *gin.Context) {
	newPlayer := &dto.PlayerRequest{}
	if err := ctx.ShouldBindJSON(newPlayer); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	createdPlayer, err := c.PlayerService.CreatePlayer(newPlayer)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, createdPlayer)
}

// DeletePlayer implements PlayerControllerInterface
func (c *Controller) DeletePlayer(ctx *gin.Context) {
	id, success := ctx.Params.Get("id")
	if !success {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err := c.PlayerService.DeletePlayer(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.Status(http.StatusOK)
}

// GetPlayerById implements PlayerControllerInterface
func (c *Controller) GetPlayerById(ctx *gin.Context) {
	id, success := ctx.Params.Get("id")
	if !success {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	player, err := c.PlayerService.GetPlayersId(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, player)
}

// GetPlayers implements PlayerControllerInterface
func (c *Controller) GetPlayers(ctx *gin.Context) {
	player, err := c.PlayerService.GetPlayers()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, player)
}

// UpdatePlayer implements PlayerControllerInterface
func (c *Controller) UpdatePlayer(ctx *gin.Context) {
	id, success := ctx.Params.Get("id")
	if !success {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	player := &dto.PlayerRequest{}
	updatedPlayer, err := c.PlayerService.UpdatePlayer(id, player)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, updatedPlayer)
}
