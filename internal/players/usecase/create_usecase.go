package usecase

import (
	"fmt"

	"github.com/surasithit/gin-basic-api/internal/players/dto"
	"github.com/surasithit/gin-basic-api/internal/players/models"
)

func (s *Service) CreatePlayer(newPlayer *dto.PlayerRequest) (*dto.Player, error) {
	player := &models.Player{
		Name:   newPlayer.Name,
		Rating: newPlayer.Rating,
	}
	createdPlayer, err := s.PlayerRepository.CreateOne(player)
	if err != nil {
		return nil, fmt.Errorf("error to create new player: %v", err.Error())
	}
	createdPlayerDto := &dto.Player{
		Id:     createdPlayer.Id,
		Name:   createdPlayer.Name,
		Rating: createdPlayer.Rating,
	}
	return createdPlayerDto, nil
}
