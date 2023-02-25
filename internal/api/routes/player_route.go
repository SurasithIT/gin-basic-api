package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/surasithit/gin-basic-api/internal/api/controller"
	"github.com/surasithit/gin-basic-api/internal/usecase"
)

func initPlayerRoutes(router *gin.RouterGroup, usecase *usecase.Service) *gin.RouterGroup {
	playerController := controller.NewPlayerController(usecase.PlayerService)

	rGroup := router.Group("/players")

	rGroup.GET("", playerController.GetPlayers)
	rGroup.POST("", playerController.CreatePlayer)
	rGroup.GET("/:id", playerController.GetPlayerById)
	rGroup.PUT("/:id", playerController.UpdatePlayer)
	rGroup.DELETE("/:id", playerController.DeletePlayer)

	return rGroup
}
