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
