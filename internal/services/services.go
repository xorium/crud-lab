package services

import (
	"crud-lab/pkg/domain/repository"
	"go.uber.org/zap"
)

var DefaultLogger, _ = zap.NewProduction()

type ServicesDeps struct {
	repo   repository.Repository
	logger *zap.Logger
}

type Services struct {
	deps  *ServicesDeps
	Users *UsersService
}

func NewServices(repo repository.Repository) *Services {
	deps := &ServicesDeps{
		repo:   repo,
		logger: DefaultLogger,
	}
	return &Services{
		deps:  deps,
		Users: &UsersService{deps},
	}
}

func (s *Services) SetLogger(logger *zap.Logger) {
	s.deps.logger = logger
}
