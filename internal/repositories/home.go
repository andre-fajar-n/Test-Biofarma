package repositories

import (
	"biofarma/gen/models"
	"context"
)

func (r *repository) HomeCreate(ctx context.Context, data *models.Home) (*models.Home, error) {
	logger := r.rt.Logger.With().Interface("data", data).Logger()

	err := r.rt.Db.Model(&data).Create(&data).Error
	if err != nil {
		logger.Error().Err(err).Msg("error query")
		return nil, err
	}

	return data, nil
}
