package handlers

import (
	"biofarma/gen/models"
	"biofarma/gen/restapi/operations/home"
	"context"
	"strconv"
	"time"

	"github.com/go-openapi/strfmt"
)

func (h *handler) CreateHome(ctx context.Context, form *home.CreateHomeParams) (*uint64, error) {
	logger := h.runtime.Logger.With().Interface("form", form).Logger()

	respData, err := h.validateLocationByAddress(ctx, *form.Data.Address)
	if err != nil {
		logger.Error().Err(err).Msg("error validateLocationByAddress")
		return nil, err
	}

	now := time.Now().UTC()
	nowStrfmt := strfmt.DateTime(now)

	req := &models.Home{
		ModelTrackTime: models.ModelTrackTime{
			CreatedAt: &nowStrfmt,
		},
		HomeData: models.HomeData{
			Address:          *form.Data.Address,
			Type:             strconv.Itoa(int(*form.Data.Type)),
			Latitude:         respData.Point.Coordinates[0],
			Longitude:        respData.Point.Coordinates[1],
			Country:          respData.Address.CountryRegion,
			Regency:          respData.Address.AdminDistrict2,
			Subdistrict:      respData.Address.Locality,
			AddressLine:      respData.Address.AddressLine,
			FormattedAddress: respData.Address.FormattedAddress,
			PostalCode:       respData.Address.PostalCode,
		},
	}

	req, err = h.repo.CreateHome(ctx, req)
	if err != nil {
		logger.Error().Err(err).Msg("error repo.HomeCreate")
		return nil, err
	}

	return req.ID, nil
}

func (h *handler) UpdateHome(ctx context.Context, form *home.UpdateHomeParams) error {
	logger := h.runtime.Logger.With().Interface("form", form).Logger()

	existingData, err := h.repo.FindOneHome(ctx, form.HomeID, false)
	if err != nil {
		logger.Error().Err(err).Msg("error repo.FindOneHome")
		return err
	}

	respData, err := h.validateLocationByAddress(ctx, *form.Data.Address)
	if err != nil {
		logger.Error().Err(err).Msg("error validateLocationByAddress")
		return err
	}

	now := time.Now().UTC()
	nowStrfmt := strfmt.DateTime(now)

	req := &models.Home{
		ModelIdentifier: models.ModelIdentifier{ID: existingData.ID},
		ModelTrackTime: models.ModelTrackTime{
			UpdatedAt: &nowStrfmt,
		},
		HomeData: models.HomeData{
			Address:          *form.Data.Address,
			Type:             strconv.Itoa(int(*form.Data.Type)),
			Latitude:         respData.Point.Coordinates[0],
			Longitude:        respData.Point.Coordinates[1],
			Country:          respData.Address.CountryRegion,
			Regency:          respData.Address.AdminDistrict2,
			Subdistrict:      respData.Address.Locality,
			AddressLine:      respData.Address.AddressLine,
			FormattedAddress: respData.Address.FormattedAddress,
			PostalCode:       respData.Address.PostalCode,
		},
	}

	err = h.repo.UpdateHome(ctx, req)
	if err != nil {
		logger.Error().Err(err).Msg("error repo.UpdateHome")
		return err
	}

	return nil
}
