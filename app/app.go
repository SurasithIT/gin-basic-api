package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start() {
	LoadConfig()

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	group := router.Group(AppConfig.HTTPPrefix)
	group.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Ok",
		})
	})

	port := fmt.Sprintf(":%d", AppConfig.HTTPPort)
	fmt.Printf("Application is running on port %s\n", port)
	err := router.Run(port)
	if err != nil {
		fmt.Printf("Application startup has an error: %v\n", err)
	}
}
