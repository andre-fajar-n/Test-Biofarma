package repositories

import (
	"biofarma/gen/models"
	"biofarma/internal/utils"
	"context"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

func (r *repository) CreateHome(ctx context.Context, data *models.Home) (*models.Home, error) {
	logger := r.rt.Logger.With().Interface("data", data).Logger()

	err := r.rt.Db.Model(&data).Create(&data).Error
	if err != nil {
		logger.Error().Err(err).Msg("error query")
		return nil, err
	}

	return data, nil
}

func (r *repository) FindOneHome(ctx context.Context, homeID uint64, includeDeletedData bool) (*models.Home, error) {
	logger := r.rt.Logger.With().
		Uint64("homeID", homeID).
		Bool("includeDeletedData", includeDeletedData).
		Logger()

	var output models.Home
	db := r.rt.Db.Model(&output).Where("id", homeID)

	if !includeDeletedData {
		db = db.Where("deleted_at IS NULL")
	}

	err := db.First(&output).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, r.rt.SetError(http.StatusNotFound, utils.ERR_HOME_NOT_FOUND)
		}
		logger.Error().Err(err).Msg("error query")
		return nil, err
	}

	return &output, nil
}

func (r *repository) UpdateHome(ctx context.Context, data *models.Home) error {
	logger := r.rt.Logger.With().Interface("data", data).Logger()

	err := r.rt.Db.Model(&data).Omit("id", "created_at", "deleted_at").Updates(&data).Error
	if err != nil {
		logger.Error().Err(err).Msg("error query")
		return err
	}

	return nil
}

func (r *repository) SoftDeleteHome(ctx context.Context, data *models.Home) error {
	logger := r.rt.Logger.With().
		Interface("data", data).
		Logger()

	err := r.rt.Db.Model(&data).Select("deleted_at").Updates(&data).Error
	if err != nil {
		logger.Error().Err(err).Msg("error query")
		return err
	}

	return nil
}

func (r *repository) FindAllPaginationHome(
	ctx context.Context,
	limit,
	offset int64,
	sortBy,
	orderBy string,
	includeDeletedData bool,
) ([]models.Home, *int64, error) {
	logger := r.rt.Logger.With().
		Int64("limit", limit).
		Int64("offset", offset).
		Str("sortBy", sortBy).
		Str("orderBy", orderBy).
		Bool("includeDeletedData", includeDeletedData).
		Logger()

	var output []models.Home
	var count int64
	db := r.rt.Db.Model(&output)

	if !includeDeletedData {
		db = db.Where("deleted_at IS NULL")
	}

	err := db.Count(&count).Error
	if err != nil {
		logger.Error().Err(err).Msg("error query count")
		return nil, nil, err
	}

	err = db.
		Order(fmt.Sprintf("%s %s", orderBy, sortBy)).
		Limit(int(limit)).
		Offset(int(offset)).
		Find(&output).Error
	if err != nil {
		logger.Error().Err(err).Msg("error query find")
		return nil, nil, err
	}

	return output, &count, nil
}
