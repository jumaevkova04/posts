package main

import (
	"fmt"
	"github.com/jumaevkova04/posts/internal/config"
	"github.com/jumaevkova04/posts/internal/handlers"
	"github.com/jumaevkova04/posts/internal/repository"
	"github.com/jumaevkova04/posts/internal/service"
	"github.com/jumaevkova04/posts/pkg/leveledlog"
	"github.com/jumaevkova04/posts/pkg/redis"
	"github.com/jumaevkova04/posts/pkg/server"
	"github.com/jumaevkova04/posts/pkg/smtp"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	cfg, err := config.NewConfig("config.json")
	if err != nil {
		fmt.Printf("config: %v\n", err)
		return
	}
	cfg.DefaultIfNotSet()

	logger := leveledlog.NewLogger(&lumberjack.Logger{
		Filename:   "logs/posts.log",
		MaxSize:    2,
		MaxBackups: 30,
		MaxAge:     40,
		Compress:   true,
	}, leveledlog.LevelAll)

	db, err := repository.NewRepository(cfg.Storage.Dsn)
	if err != nil {
		logger.Fatal(err)
	}

	redisDB, err := redis.NewDB(cfg.Redis.Addr, cfg.Redis.Password, cfg.Redis.DB)
	if err != nil {
		logger.Fatal(err)
	}

	mailer := smtp.NewMailer(cfg.Smtp.Host, cfg.Smtp.Port, cfg.Smtp.Username, cfg.Smtp.Password, cfg.Smtp.From)

	srv := service.NewService(cfg, db, redisDB, mailer)

	app := &handlers.Handler{
		Config:  cfg,
		Logger:  logger,
		Service: srv,
	}

	logger.Info("starting server on %s", cfg.Addr)
	fmt.Println("server is listening...")

	if err = server.Run(cfg.Addr, app.Routes()); err != nil {
		logger.Fatal(err)
	}

	logger.Info("server stopped")
}
