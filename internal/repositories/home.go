package repositories

import (
	"biofarma/gen/models"
	"context"
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
			return nil, r.rt.SetError(http.StatusNotFound, "home not found")
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
