package utils

import "github.com/surasithit/gin-basic-api/config"

type UtilService struct {
	Config *config.Config
}

type UtilServiceInterface interface {
}

var _ UtilServiceInterface = (*UtilService)(nil)

func NewService(config *config.Config) *UtilService {
	return &UtilService{
		Config: config,
	}
}
