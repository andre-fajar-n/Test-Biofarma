package repositories

import (
	"biofarma/gen/client/bing_maps_client"

	"github.com/go-openapi/strfmt"
)

func (r *repository) createClientBingMaps() *bing_maps_client.BingMaps {
	return bing_maps_client.NewHTTPClientWithConfig(strfmt.Default, &bing_maps_client.TransportConfig{
		Host:     r.rt.Cfg.BingMapsHost,
		BasePath: bing_maps_client.DefaultBasePath,
		Schemes:  nil,
	})
}
