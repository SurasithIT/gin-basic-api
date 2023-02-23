package players

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.RouterGroup, controller *Controller) *gin.RouterGroup {
	rGroup := router.Group("/players")

	rGroup.GET("", func(c *gin.Context) { controller.GetPlayers(c) })
	rGroup.POST("", func(c *gin.Context) { controller.CreatePlayer(c) })
	rGroup.GET("/:id", func(c *gin.Context) { controller.GetPlayerById(c) })
	rGroup.PUT("/:id", func(c *gin.Context) { controller.UpdatePlayer(c) })
	rGroup.DELETE("/:id", func(c *gin.Context) { controller.DeletePlayer(c) })

	return rGroup
}
