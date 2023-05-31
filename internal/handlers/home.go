package handlers

import (
	"biofarma/gen/models"
	"biofarma/gen/restapi/operations/home"
	"context"
	"math"
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

func (h *handler) FindOneHome(ctx context.Context, form *home.FindOneHomeParams) (*models.SuccessFindOneAllOf1, error) {
	logger := h.runtime.Logger.With().Interface("form", form).Logger()

	includeDeletedData, _ := strconv.ParseBool(*form.IncludeDeletedData)
	data, err := h.repo.FindOneHome(ctx, form.HomeID, includeDeletedData)
	if err != nil {
		logger.Error().Err(err).Msg("error repo.FindOneHome")
		return nil, err
	}

	output := models.SuccessFindOneAllOf1Data(*data)
	return &models.SuccessFindOneAllOf1{Data: &output}, nil
}

func (h *handler) DeleteHome(ctx context.Context, form *home.DeleteHomeParams) error {
	logger := h.runtime.Logger.With().Interface("form", form).Logger()

	existingData, err := h.repo.FindOneHome(ctx, form.HomeID, false)
	if err != nil {
		logger.Error().Err(err).Msg("error repo.FindOneHome")
		return err
	}

	now := time.Now().UTC()
	nowStrfmt := strfmt.DateTime(now)

	req := &models.Home{
		ModelIdentifier: models.ModelIdentifier{ID: existingData.ID},
		ModelTrackTime: models.ModelTrackTime{
			DeletedAt: &nowStrfmt,
		},
	}

	err = h.repo.SoftDeleteHome(ctx, req)
	if err != nil {
		logger.Error().Err(err).Msg("error repo.SoftDeleteHome")
		return err
	}

	return nil
}

func (h *handler) FindAllPaginationHome(ctx context.Context, form *home.FindAllPaginationHomeParams) (*models.SuccessFindAllPaginationAllOf1, error) {
	logger := h.runtime.Logger.With().Interface("form", form).Logger()

	includeDeletedData, _ := strconv.ParseBool(*form.IncludeDeletedData)
	limit := *form.Limit
	offset := (*form.Page - 1) * *form.Limit

	data, totalData, err := h.repo.FindAllPaginationHome(ctx, limit, offset, *form.SortBy, *form.OrderBy, includeDeletedData)
	if err != nil {
		logger.Error().Err(err).Msg("error repo.FindAllPaginationHome")
		return nil, err
	}

	totalPage := math.Ceil((float64(*totalData) / float64(*form.Limit)))
	output := models.SuccessFindAllPaginationAllOf1{
		Data: make([]*models.SuccessFindOneAllOf1Data, len(data)),
		Metadata: &models.SuccessFindAllPaginationAllOf1Metadata{
			Metadata: models.Metadata{
				Page:      *form.Page,
				PerPage:   *form.Limit,
				TotalPage: int64(totalPage),
				TotalRow:  *totalData,
			},
			CustomFields: models.CustomFields{
				"include_deleted_data": form.IncludeDeletedData,
			},
		},
	}

	for i, v := range data {
		temp := models.SuccessFindOneAllOf1Data(v)
		output.Data[i] = &temp
	}

	return &output, nil
}
