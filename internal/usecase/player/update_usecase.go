package player

import (
	"fmt"

	"github.com/surasithit/gin-basic-api/domain"
)

func (s *Service) UpdatePlayer(id string, player *domain.PlayerRequest) (*domain.Player, error) {
	existPlayer, err := s.PlayerRepository.FindOneById(id)
	if err != nil {
		return nil, fmt.Errorf("error to get player: %v", err)
	}

	existPlayer.Name = player.Name
	existPlayer.Rating = player.Rating

	updatedPlayer, err := s.PlayerRepository.UpdateById(id, existPlayer)
	if err != nil {
		return nil, fmt.Errorf("error to get player: %v", err)
	}
	playerDto := &domain.Player{
		Id:     updatedPlayer.Id,
		Name:   updatedPlayer.Name,
		Rating: updatedPlayer.Rating,
	}

	return playerDto, nil
}
