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
		CreateHome(ctx context.Context, data *models.Home) (*models.Home, error)
		FindOneHome(ctx context.Context, homeID uint64, includeDeletedData bool) (*models.Home, error)
		UpdateHome(ctx context.Context, data *models.Home) error
		SoftDeleteHome(ctx context.Context, data *models.Home) error
		FindAllPaginationHome(
			ctx context.Context,
			limit,
			offset int64,
			sortBy,
			orderBy string,
			includeDeletedData bool,
		) ([]models.Home, *int64, error)
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
