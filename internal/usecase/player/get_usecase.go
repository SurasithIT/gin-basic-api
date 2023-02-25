package player

import (
	"fmt"

	"github.com/surasithit/gin-basic-api/domain"
)

func (s *Service) GetPlayers() ([]*domain.Player, error) {
	players, err := s.PlayerRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("error to get all players: %v", err)
	}
	playersDto := []*domain.Player{}
	for _, player := range players {
		playersDto = append(playersDto, &domain.Player{
			Id:     player.Id,
			Name:   player.Name,
			Rating: player.Rating,
		})
	}
	return playersDto, nil
}

func (s *Service) GetPlayerById(id string) (*domain.Player, error) {
	player, err := s.PlayerRepository.FindOneById(id)
	if err != nil {
		return nil, fmt.Errorf("error to get player: %v", err)
	}
	playerDto := &domain.Player{
		Id:     player.Id,
		Name:   player.Name,
		Rating: player.Rating,
	}
	return playerDto, nil
}
