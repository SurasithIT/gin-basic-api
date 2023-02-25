package domain

import "github.com/surasithit/gin-basic-api/internal/model"

// Player defines model for Player.
type Player struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Rating float64 `json:"rating"`
}

// PlayerRequest defines model for PlayerRequest.
type PlayerRequest struct {
	Name   string  `json:"name"  binding:"required"`
	Rating float64 `json:"rating" binding:"required,gte=0,lte=10"`
}

// PlayerUsecase represent player usecase
type PlayerUsecase interface {
	GetPlayers() ([]*Player, error)
	GetPlayerById(id string) (*Player, error)
	CreatePlayer(newPlayer *PlayerRequest) (*Player, error)
	UpdatePlayer(id string, player *PlayerRequest) (*Player, error)
	DeletePlayer(id string) error
}

// PlayerUsecase represent player usecase
type PlayerRepository interface {
	FindAll() ([]*model.Player, error)
	FindOneById(id string) (*model.Player, error)
	CreateOne(newPlayer *model.Player) (*model.Player, error)
	UpdateById(id string, player *model.Player) (*model.Player, error)
	DeleteById(id string) error
}
