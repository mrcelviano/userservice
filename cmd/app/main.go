package main

import (
	"context"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mrcelviano/userservice/internal/config"
	"github.com/mrcelviano/userservice/internal/delivery"
	"github.com/mrcelviano/userservice/internal/repository"
	"github.com/mrcelviano/userservice/internal/service"
	"github.com/mrcelviano/userservice/pkg/database/postgres"
	"github.com/mrcelviano/userservice/pkg/logger"
	"github.com/mrcelviano/userservice/pkg/notification"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	configsDirectory = "configs"

	contextTimeoutValue = 5 * time.Second
)

func main() {
	cfg, err := config.Init(configsDirectory)
	if err != nil {
		logger.Error(err)
		return
	}

	postgresConnection, err := postgres.NewGoCraftDBConnectionPG(cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User,
		cfg.Postgres.Password, cfg.Postgres.DBName)
	if err != nil {
		logger.Error(err)
		return
	}

	//rpc
	notificationClient, err := notification.NewNotificationClient(cfg.Services)
	if err != nil {
		logger.Error(err)
		return
	}

	//repo
	var (
		userRepo = repository.NewUserRepositoryPG(postgresConnection)
	)

	//service
	var (
		userLogic = service.NewUserService(userRepo, notificationClient)
	)

	//delivery
	e := echo.New()
	e.Pre(
		middleware.AddTrailingSlash(),
	)
	delivery.NewUserHandlers(e.Group("api"), userLogic)

	logger.Info("server start")

	go func() {
		err := e.Start(":" + cfg.HTTP.Port)
		if err != nil {
			logger.Errorf("can`t run http server: %s\n", err.Error())
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), contextTimeoutValue)
	defer shutdown()

	err = e.Shutdown(ctx)
	if err != nil {
		logger.Errorf("can`t stop server: %v", err.Error())
		return
	}
}
