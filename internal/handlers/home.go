package handlers

import (
	"biofarma/gen/models"
	"biofarma/gen/restapi/operations/home"
	"context"
	"strconv"
	"time"

	"github.com/go-openapi/strfmt"
)

func (h *handler) HomeCreate(ctx context.Context, form *home.CreateHomeParams) (*uint64, error) {
	logger := h.runtime.Logger.With().Interface("form", form).Logger()

	now := time.Now().UTC()
	nowStrfmt := strfmt.DateTime(now)
	var err error

	req := &models.Home{
		ModelTrackTime: models.ModelTrackTime{
			CreatedAt: &nowStrfmt,
		},
		HomeData: models.HomeData{
			Address: *form.Data.Address,
			Type:    strconv.Itoa(int(*form.Data.Type)),
		},
	}

	req, err = h.repo.HomeCreate(ctx, req)
	if err != nil {
		logger.Error().Err(err).Msg("error repo.HomeCreate")
		return nil, err
	}

	return req.ID, nil
}
