package players

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.RouterGroup, controller *Controller) *gin.RouterGroup {
	rGroup := router.Group("/players")

	rGroup.GET("", controller.GetPlayers)
	rGroup.POST("", controller.CreatePlayer)
	rGroup.GET("/:id", controller.GetPlayerById)
	rGroup.PUT("/:id", controller.UpdatePlayer)
	rGroup.DELETE("/:id", controller.DeletePlayer)

	return rGroup
}
