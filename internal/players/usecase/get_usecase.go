package usecase

import (
	"fmt"

	"github.com/surasithit/gin-basic-api/internal/players/dto"
)

func (s *Service) GetPlayers() ([]*dto.Player, error) {
	players, err := s.PlayerRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("error to get all players: %v", err)
	}
	playersDto := []*dto.Player{}
	for _, player := range players {
		playersDto = append(playersDto, &dto.Player{
			Id:     player.Id,
			Name:   player.Name,
			Rating: player.Rating,
		})
	}
	return playersDto, nil
}

func (s *Service) GetPlayersId(id string) (*dto.Player, error) {
	player, err := s.PlayerRepository.FindOneById(id)
	if err != nil {
		return nil, fmt.Errorf("error to get player: %v", err)
	}
	playerDto := &dto.Player{
		Id:     player.Id,
		Name:   player.Name,
		Rating: player.Rating,
	}
	return playerDto, nil
}
