package transport

import (
	"net/http"
	"url-shortener/internal/config"
	"url-shortener/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Init(cfg *config.Config) *gin.Engine {
	switch cfg.Environment {
	case config.EnvLocal:
		gin.SetMode(gin.DebugMode)
	case config.EnvProd:
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return router
}
