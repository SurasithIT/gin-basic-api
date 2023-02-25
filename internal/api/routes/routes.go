package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/surasithit/gin-basic-api/internal/usecase"
)

func InitialRoutes(rGroup *gin.RouterGroup, usecase *usecase.Service) *gin.RouterGroup {
	initPlayerRoutes(rGroup, usecase)

	return rGroup
}
