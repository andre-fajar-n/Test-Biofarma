package handlers

import (
	"biofarma/gen/models"
	"biofarma/gen/restapi/operations/route"
	"biofarma/internal/utils"
	"context"
	"net/http"
)

func (h *handler) FindRoute(ctx context.Context, form *route.FindRouteParams) (*models.SuccessFindRouteAllOf1, error) {
	logger := h.runtime.Logger.With().
		Interface("form", form).
		Logger()

	totalData, err := h.repo.CountByIDsHome(ctx, form.HomeIds, false)
	if err != nil {
		logger.Error().Err(err).Msg("error repo.CountByIDsHome")
		return nil, err
	}

	if len(form.HomeIds) != int(*totalData) {
		return nil, h.runtime.SetError(http.StatusBadRequest, utils.ERR_THERE_IS_DATA_DELETED)
	}

	data, err := h.repo.FindByIDsHome(ctx, form.HomeIds, false)
	if err != nil {
		logger.Error().Err(err).Msg("repo.FindByIDsHome")
		return nil, err
	}

	output := models.SuccessFindRouteAllOf1{
		Data: make([]*models.SuccessFindRouteAllOf1DataItems, len(data)-1),
	}

	for i := 0; i < len(data)-1; i++ {
		firstData := data[i]
		secondData := data[i+1]

		output.Data[i] = &models.SuccessFindRouteAllOf1DataItems{
			Destination: &models.HomeWithoutTrackTime{
				ModelIdentifier: firstData.ModelIdentifier,
				HomeData:        firstData.HomeData,
			},
			Origin: &models.HomeWithoutTrackTime{
				ModelIdentifier: secondData.ModelIdentifier,
				HomeData:        secondData.HomeData,
			},
			Distance: utils.DistanceByLatLon(firstData.Latitude, firstData.Longitude, secondData.Latitude, secondData.Longitude),
		}
	}

	return &output, nil
}
