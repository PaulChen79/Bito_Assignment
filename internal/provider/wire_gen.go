// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package provider

import (
	"bito-assignment/config"
	"bito-assignment/domain"
	"bito-assignment/internal/repo"
	"bito-assignment/internal/service"
	"sync"
)

// Injectors from wire.go:

func NewRepo() (domain.Repository, error) {
	repository := repo.NewRepository()
	return repository, nil
}

func NewService() (domain.Service, error) {
	config := NewConfig()
	repository := repo.NewRepository()
	domainService := service.NewService(config, repository)
	return domainService, nil
}

// wire.go:

var cg *config.Config

var configOnce sync.Once

func NewConfig() *config.Config {
	configOnce.Do(func() {
		cg = config.NewConfig()
	})
	return cg
}
