package repository

import "github.com/surasithit/gin-basic-api/internal/players/models"

type PlayerRepositoryInterface interface {
	FindAll() ([]*models.Player, error)
	FindOneById(id string) (*models.Player, error)
	CreateOne(newPlayer *models.Player) (*models.Player, error)
	UpdateById(id string, player *models.Player) (*models.Player, error)
	DeleteById(id string) error
}

type PlayerRepository struct {
}
