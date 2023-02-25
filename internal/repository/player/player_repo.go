package player

import (
	"github.com/google/uuid"
	"github.com/surasithit/gin-basic-api/domain"
	"github.com/surasithit/gin-basic-api/internal/model"
)

type Repository struct {
}

// Implement interface
var _ domain.PlayerRepository = (*Repository)(nil)

func NewRepository() *Repository {
	return &Repository{}
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

// CreateOne implements RepositoryInterface
func (r *Repository) CreateOne(newPlayer *model.Player) (*model.Player, error) {
	newPlayer.Id = uuid.NewString()
	mockPlayers = append(mockPlayers, newPlayer)
	return newPlayer, nil
}

// DeleteById implements RepositoryInterface
func (r *Repository) DeleteById(id string) error {
	for i, player := range mockPlayers {
		if player.Id == id {
			mockPlayers = append(mockPlayers[:i], mockPlayers[i+1:]...)
		}
	}
	return nil
}

// FindAll implements RepositoryInterface
func (r *Repository) FindAll() ([]*model.Player, error) {
	return mockPlayers, nil
}

// FindOneById implements RepositoryInterface
func (r *Repository) FindOneById(id string) (*model.Player, error) {
	res := &model.Player{}
	for _, player := range mockPlayers {
		if player.Id == id {
			res = player
			break
		}
	}
	return res, nil
}

// UpdateById implements RepositoryInterface
func (r *Repository) UpdateById(id string, player *model.Player) (*model.Player, error) {
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
