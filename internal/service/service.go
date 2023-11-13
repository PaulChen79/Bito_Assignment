package service

import (
	"bito-assignment/config"
	"bito-assignment/domain"
)

type service struct {
	config *config.Config
	repo   domain.Repository
}

func NewService(config *config.Config, repo domain.Repository) domain.Service {
	return &service{
		config: config,
		repo:   repo,
	}
}

func NewServiceForTest(config *config.Config, repo domain.Repository) *service {
	return &service{
		config: config,
		repo:   repo,
	}
}
