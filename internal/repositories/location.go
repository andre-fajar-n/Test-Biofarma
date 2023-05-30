package repositories

import (
	"biofarma/gen/client/bing_maps_client/location"
	"biofarma/gen/models"
	"context"
	"fmt"
)

func (r *repository) FindLocation(ctx context.Context, form *location.FindLocationParams) (*models.SuccessGetLocation, error) {
	logger := r.rt.Logger.With().Interface("form", form).Logger()

	client := r.createClientBingMaps()

	form.Context = ctx
	form.Key = &r.rt.Cfg.BingMapsKey

	resp, err := client.Location.FindLocation(form)
	if err != nil {
		logger.Error().Err(err).Msg("error FindLocation")
		return nil, err
	}

	if !resp.IsSuccess() {
		msg := resp.Error()
		logger.Error().Msg(msg)
		return nil, fmt.Errorf(msg)
	}

	return resp.Payload, nil
}
