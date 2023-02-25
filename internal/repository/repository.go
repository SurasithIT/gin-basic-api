package repository

import (
	"github.com/surasithit/gin-basic-api/internal/repository/player"
)

type Repository struct {
	PlayerRepository *player.PlayerRepository
}

func NewRepository() *Repository {
	return &Repository{}
}
