package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/surasithit/gin-basic-api/domain"
	"github.com/surasithit/gin-basic-api/internal/usecase/player"
)

type PlayerControllerInterface interface {
	GetPlayers(ctx *gin.Context)
	GetPlayerById(ctx *gin.Context)
	CreatePlayer(ctx *gin.Context)
	UpdatePlayer(ctx *gin.Context)
	DeletePlayer(ctx *gin.Context)
}

type PlayerController struct {
	PlayerService *player.Service
}

// Implement interface
var _ PlayerControllerInterface = (*PlayerController)(nil)

func NewPlayerController(playerService *player.Service) *PlayerController {
	return &PlayerController{
		PlayerService: playerService,
	}
}

// CreatePlayer implements PlayerControllerInterface
func (c *PlayerController) CreatePlayer(ctx *gin.Context) {
	newPlayer := &domain.PlayerRequest{}
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
func (c *PlayerController) DeletePlayer(ctx *gin.Context) {
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
func (c *PlayerController) GetPlayerById(ctx *gin.Context) {
	id, success := ctx.Params.Get("id")
	if !success {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	player, err := c.PlayerService.GetPlayerById(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, player)
}

// GetPlayers implements PlayerControllerInterface
func (c *PlayerController) GetPlayers(ctx *gin.Context) {
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
func (c *PlayerController) UpdatePlayer(ctx *gin.Context) {
	id, success := ctx.Params.Get("id")
	if !success {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	player := &domain.PlayerRequest{}
	if err := ctx.ShouldBindJSON(player); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	updatedPlayer, err := c.PlayerService.UpdatePlayer(id, player)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, updatedPlayer)
}
