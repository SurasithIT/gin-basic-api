package player

import (
	"fmt"

	"github.com/surasithit/gin-basic-api/internal/dto"
)

func (s *Service) UpdatePlayer(id string, player *dto.PlayerRequest) (*dto.Player, error) {
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
	playerDto := &dto.Player{
		Id:     updatedPlayer.Id,
		Name:   updatedPlayer.Name,
		Rating: updatedPlayer.Rating,
	}

	return playerDto, nil
}
