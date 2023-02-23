package usecase

import (
	"github.com/surasithit/gin-basic-api/internal/players/dto"
	"github.com/surasithit/gin-basic-api/internal/players/repository"
)

type PlayerServiceInterface interface {
	GetPlayers() ([]*dto.Player, error)
	GetPlayerById(id string) (*dto.Player, error)
	CreatePlayer(newPlayer *dto.PlayerRequest) (*dto.Player, error)
	UpdatePlayer(id string, player *dto.PlayerRequest) (*dto.Player, error)
	DeletePlayer(id string) error
}

type Service struct {
	PlayerRepository *repository.PlayerRepository
}

func NewService(playerRepository *repository.PlayerRepository) *Service {
	return &Service{
		PlayerRepository: playerRepository,
	}
}
