package player

import (
	"github.com/surasithit/gin-basic-api/internal/dto"
	"github.com/surasithit/gin-basic-api/internal/repository/player"
)

type PlayerServiceInterface interface {
	GetPlayers() ([]*dto.Player, error)
	GetPlayerById(id string) (*dto.Player, error)
	CreatePlayer(newPlayer *dto.PlayerRequest) (*dto.Player, error)
	UpdatePlayer(id string, player *dto.PlayerRequest) (*dto.Player, error)
	DeletePlayer(id string) error
}

type Service struct {
	PlayerRepository *player.PlayerRepository
}

func NewService(playerRepository *player.PlayerRepository) *Service {
	return &Service{
		PlayerRepository: playerRepository,
	}
}
