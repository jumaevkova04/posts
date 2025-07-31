package service

import (
	"github.com/jumaevkova/posts/internal/config"
	"github.com/jumaevkova/posts/internal/repository"
	"github.com/jumaevkova/posts/pkg/redis"
	"github.com/jumaevkova/posts/pkg/smtp"
)

type Service struct {
	Config     *config.Config
	Repository *repository.Repository
	RedisDB    *redis.DB
	Mailer     *smtp.Mailer
}

func NewService(
	config *config.Config,
	repository *repository.Repository,
	redisDB *redis.DB,
	mailer *smtp.Mailer,
) *Service {
	return &Service{
		Config:     config,
		Repository: repository,
		RedisDB:    redisDB,
		Mailer:     mailer,
	}
}
