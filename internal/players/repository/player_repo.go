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

var _ PlayerRepositoryInterface = (*PlayerRepository)(nil)

func NewRepository() *PlayerRepository {
	return &PlayerRepository{}
}

// CreateOne implements PlayerRepositoryInterface
func (*PlayerRepository) CreateOne(newPlayer *models.Player) (*models.Player, error) {
	panic("unimplemented")
}

// DeleteById implements PlayerRepositoryInterface
func (*PlayerRepository) DeleteById(id string) error {
	panic("unimplemented")
}

// FindAll implements PlayerRepositoryInterface
func (*PlayerRepository) FindAll() ([]*models.Player, error) {
	panic("unimplemented")
}

// FindOneById implements PlayerRepositoryInterface
func (*PlayerRepository) FindOneById(id string) (*models.Player, error) {
	panic("unimplemented")
}

// UpdateById implements PlayerRepositoryInterface
func (*PlayerRepository) UpdateById(id string, player *models.Player) (*models.Player, error) {
	panic("unimplemented")
}