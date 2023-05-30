package handlers

import (
	"biofarma/gen/client/bing_maps_client/location"
	"biofarma/gen/models"
	"biofarma/gen/restapi/operations/home"
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/go-openapi/strfmt"
)

func (h *handler) HomeCreate(ctx context.Context, form *home.CreateHomeParams) (*uint64, error) {
	logger := h.runtime.Logger.With().Interface("form", form).Logger()

	respLocation, err := h.repo.FindLocation(ctx, &location.FindLocationParams{Query: form.Data.Address})
	if err != nil {
		logger.Error().Err(err).Msg("error repo.FindLocation")
		return nil, err
	}

	if respLocation.ResourceSets[0].EstimatedTotal > 1 {
		return nil, h.runtime.SetError(http.StatusBadRequest, "address that you filled founded more than 1, please more specify")
	}

	respData := respLocation.ResourceSets[0].Resources[0]

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

	req, err = h.repo.HomeCreate(ctx, req)
	if err != nil {
		logger.Error().Err(err).Msg("error repo.HomeCreate")
		return nil, err
	}

	return req.ID, nil
}
