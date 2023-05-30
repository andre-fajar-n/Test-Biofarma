package repositories

import (
	"biofarma/gen/client/bing_maps_client/location"
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
		locationInterface
	}

	homeInterface interface {
		HomeCreate(ctx context.Context, data *models.Home) (*models.Home, error)
	}

	locationInterface interface {
		FindLocation(ctx context.Context, form *location.FindLocationParams) (*models.SuccessGetLocation, error)
	}
)

func New(
	rt runtime.Runtime,
) RepositoryInterface {
	return &repository{
		rt,
	}
}
