package handlers

import (
	"biofarma/gen/models"
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
		CreateHome(ctx context.Context, form *home.CreateHomeParams) (*uint64, error)
		UpdateHome(ctx context.Context, form *home.UpdateHomeParams) error
		FindOneHome(ctx context.Context, form *home.FindOneHomeParams) (*models.SuccessFindOneAllOf1, error)
		DeleteHome(ctx context.Context, form *home.DeleteHomeParams) error
		FindAllPaginationHome(ctx context.Context, form *home.FindAllPaginationHomeParams) (*models.SuccessFindAllPaginationAllOf1, error)
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
