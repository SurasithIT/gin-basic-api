package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/surasithit/gin-basic-api/config"
	"github.com/surasithit/gin-basic-api/internal/controller"
	"github.com/surasithit/gin-basic-api/internal/repository"
	"github.com/surasithit/gin-basic-api/internal/routes"
	"github.com/surasithit/gin-basic-api/internal/usecase"
)

func Start() {
	AppConfig, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "PATCH"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	initialApp(router, AppConfig)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", AppConfig.HTTPPort),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

func initialApp(router *gin.Engine, config *config.Config) {
	rGroup := router.Group(config.HTTPPrefix)
	rGroup.GET("/health", HealthCheckHandler())

	repo := repository.NewRepository()
	usecases := usecase.NewService(repo)
	controllers := controller.NewController(usecases)

	routes.InitialRoutes(rGroup, controllers)
}

func HealthCheckHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Ok",
		})
	}
}
