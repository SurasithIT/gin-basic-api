package player

import (
	"github.com/google/uuid"
	"github.com/surasithit/gin-basic-api/internal/model"
)

type PlayerRepositoryInterface interface {
	FindAll() ([]*model.Player, error)
	FindOneById(id string) (*model.Player, error)
	CreateOne(newPlayer *model.Player) (*model.Player, error)
	UpdateById(id string, player *model.Player) (*model.Player, error)
	DeleteById(id string) error
}

type PlayerRepository struct {
}

// Implement interface
var _ PlayerRepositoryInterface = (*PlayerRepository)(nil)

func NewRepository() *PlayerRepository {
	return &PlayerRepository{}
}

var mockPlayers = []*model.Player{
	{
		Id:     "e9f3538e-d322-4c83-b794-310592598f56",
		Name:   "Lionel Messi",
		Rating: 9.9,
	},
	{
		Id:     "3d6ee7d6-2934-400d-97ad-7da2fb41f7c5",
		Name:   "Cristiano Ronaldo",
		Rating: 9.9,
	},
	{
		Id:     "a190b9fd-f5e9-47e7-bf2e-7bb9c32e64dc",
		Name:   "Neymar Jr.",
		Rating: 9.7,
	},
}

// CreateOne implements PlayerRepositoryInterface
func (r *PlayerRepository) CreateOne(newPlayer *model.Player) (*model.Player, error) {
	newPlayer.Id = uuid.NewString()
	mockPlayers = append(mockPlayers, newPlayer)
	return newPlayer, nil
}

// DeleteById implements PlayerRepositoryInterface
func (r *PlayerRepository) DeleteById(id string) error {
	for i, player := range mockPlayers {
		if player.Id == id {
			mockPlayers = append(mockPlayers[:i], mockPlayers[i+1:]...)
		}
	}
	return nil
}

// FindAll implements PlayerRepositoryInterface
func (r *PlayerRepository) FindAll() ([]*model.Player, error) {
	return mockPlayers, nil
}

// FindOneById implements PlayerRepositoryInterface
func (r *PlayerRepository) FindOneById(id string) (*model.Player, error) {
	res := &model.Player{}
	for _, player := range mockPlayers {
		if player.Id == id {
			res = player
			break
		}
	}
	return res, nil
}

// UpdateById implements PlayerRepositoryInterface
func (r *PlayerRepository) UpdateById(id string, player *model.Player) (*model.Player, error) {
	res := &model.Player{}
	for _, p := range mockPlayers {
		if p.Id == id {
			p = player
			res = p
			break
		}
	}
	return res, nil
}
