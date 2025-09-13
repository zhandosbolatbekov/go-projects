package app

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"url-shortener/internal/config"
	"url-shortener/internal/repository"
	"url-shortener/internal/server"
	"url-shortener/internal/service"
	"url-shortener/internal/transport"
	"url-shortener/pkg"
)

func Run(configPath string) {
	// Configs
	cfg, err := config.Init(configPath)
	if err != nil {
		log.Fatal("error loading config: ", err)
		return
	}

	// Dependencies
	var logLevel slog.Level
	switch cfg.Environment {
	case config.EnvLocal:
		logLevel = slog.LevelDebug
	case config.EnvProd:
		logLevel = slog.LevelInfo
	}
	logger := pkg.NewLogger(pkg.LoggerConfig{Level: logLevel})

	// Services, Repos & API Handlers
	repos := repository.NewRepositories()
	services := service.NewServices(repos)
	handlers := transport.NewHandler(services)

	// HTTP Server
	srv := server.NewServer(cfg, handlers.Init(cfg))
	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Error("error occurred while running http server: " + err.Error())
		}
	}()
	logger.Info("Server started")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.Error("failed to stop server: " + err.Error())
	}

	//if err := mongoClient.Disconnect(context.Background()); err != nil {
	//	logger.Error(err.Error())
	//}
}
