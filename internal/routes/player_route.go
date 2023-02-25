package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/surasithit/gin-basic-api/internal/controller/player"
)

func initPlayerRoutes(router *gin.RouterGroup, playerController *player.PlayerController) *gin.RouterGroup {
	rGroup := router.Group("/players")

	rGroup.GET("", playerController.GetPlayers)
	rGroup.POST("", playerController.CreatePlayer)
	rGroup.GET("/:id", playerController.GetPlayerById)
	rGroup.PUT("/:id", playerController.UpdatePlayer)
	rGroup.DELETE("/:id", playerController.DeletePlayer)

	return rGroup
}
