package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/surasithit/gin-basic-api/internal/controller"
)

func InitialRoutes(rGroup *gin.RouterGroup, c *controller.Controller) *gin.RouterGroup {
	initPlayerRoutes(rGroup, c.PlayerController)

	return rGroup
}
