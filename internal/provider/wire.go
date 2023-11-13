//go:build wireinject
// +build wireinject

package provider

import (
	"sync"

	"bito-assignment/config"
	"bito-assignment/domain"
	repo "bito-assignment/internal/repo"
	svc "bito-assignment/internal/service"

	"github.com/google/wire"
)

var cg *config.Config
var configOnce sync.Once

func NewConfig() *config.Config {
	configOnce.Do(func() {
		cg = config.NewConfig()
	})
	return cg
}

func NewRepo() (domain.Repository, error) {
	panic(wire.Build(repo.NewRepository))
}

func NewService() (domain.Service, error) {
	panic(wire.Build(svc.NewService, NewConfig, repo.NewRepository))
}
