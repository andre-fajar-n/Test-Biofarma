package handlers

import (
	"biofarma/gen/restapi/operations/home"
	"biofarma/internal/repositories"
	"biofarma/runtime"
	"context"
)

type (
	handler struct {
		runtime runtime.Runtime
		repo    repositories.RepositoryInterface
	}

	HandlerInterface interface {
		homeInterface
	}

	homeInterface interface {
		HomeCreate(ctx context.Context, form *home.CreateHomeParams) (*uint64, error)
	}
)

func NewHandler(
	rt runtime.Runtime,
	repo repositories.RepositoryInterface,
) HandlerInterface {
	return &handler{
		rt,
		repo,
	}
}
