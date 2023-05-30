package repositories

import (
	"biofarma/gen/models"
	"biofarma/runtime"
	"context"
)

type (
	repository struct {
		rt runtime.Runtime
	}

	RepositoryInterface interface {
		homeInterface
	}

	homeInterface interface {
		HomeCreate(ctx context.Context, data *models.Home) (*models.Home, error)
	}
)

func New(
	rt runtime.Runtime,
) RepositoryInterface {
	return &repository{
		rt,
	}
}
