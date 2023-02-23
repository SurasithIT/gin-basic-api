package players

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.RouterGroup) *gin.RouterGroup {
	rGroup := router.Group("/players")

	rGroup.GET("", mockSuccessHandler())
	rGroup.POST("", mockSuccessHandler())
	rGroup.GET("/:id", mockSuccessHandler())
	rGroup.PUT("/:id", mockSuccessHandler())
	rGroup.DELETE("/:id", mockSuccessHandler())

	return rGroup
}

func mockSuccessHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Ok",
		})
	}
}
