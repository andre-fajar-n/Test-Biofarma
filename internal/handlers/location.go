package handlers

import (
	"biofarma/gen/client/bing_maps_client/location"
	"biofarma/gen/models"
	"biofarma/internal/utils"
	"context"
	"net/http"
)

func (h *handler) validateLocationByAddress(ctx context.Context, address string) (*models.SuccessGetLocationResourceSetsItemsResourcesItems, error) {
	logger := h.runtime.Logger.With().Str("address", address).Logger()

	respLocation, err := h.repo.FindLocation(ctx, &location.FindLocationParams{Query: &address})
	if err != nil {
		logger.Error().Err(err).Msg("error repo.FindLocation")
		return nil, err
	}

	if respLocation.ResourceSets[0].EstimatedTotal > 1 {
		return nil, h.runtime.SetError(http.StatusBadRequest, utils.ERR_ADDRESS_FOUND_MORE_THAN_ONE)
	}

	return respLocation.ResourceSets[0].Resources[0], nil
}
