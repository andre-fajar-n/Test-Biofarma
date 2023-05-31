package rest

import (
	"biofarma/gen/models"
	"biofarma/gen/restapi/operations"
	"biofarma/gen/restapi/operations/health"
	"biofarma/gen/restapi/operations/home"
	"biofarma/internal/handlers"
	"biofarma/runtime"
	"context"

	"github.com/go-openapi/runtime/middleware"
)

func Route(rt *runtime.Runtime, api *operations.ServerAPI, apiHandler handlers.HandlerInterface) {
	//  health
	api.HealthHealthHandler = health.HealthHandlerFunc(func(hp health.HealthParams) middleware.Responder {
		return health.NewHealthOK().WithPayload(&models.Success{
			Message: "Server up",
		})
	})

	// home
	{
		api.HomeCreateHomeHandler = home.CreateHomeHandlerFunc(func(chp home.CreateHomeParams) middleware.Responder {
			homeID, err := apiHandler.CreateHome(context.Background(), &chp)
			if err != nil {
				errRes := rt.GetError(err)
				return home.NewCreateHomeDefault(int(errRes.Code())).WithPayload(&models.Error{
					Code:    int64(errRes.Code()),
					Message: errRes.Error(),
				})
			}

			return home.NewCreateHomeCreated().WithPayload(&models.SuccessCreate{
				SuccessCreateAllOf0: models.SuccessCreateAllOf0{
					Message: "Success create home",
				},
				SuccessCreateAllOf1: models.SuccessCreateAllOf1{
					HomeID: *homeID,
				},
			})
		})

		api.HomeUpdateHomeHandler = home.UpdateHomeHandlerFunc(func(uhp home.UpdateHomeParams) middleware.Responder {
			err := apiHandler.UpdateHome(context.Background(), &uhp)
			if err != nil {
				errRes := rt.GetError(err)
				return home.NewUpdateHomeDefault(int(errRes.Code())).WithPayload(&models.Error{
					Code:    int64(errRes.Code()),
					Message: errRes.Error(),
				})
			}

			return home.NewUpdateHomeOK().WithPayload(&models.Success{
				Message: "Success update home",
			})
		})
	}
}
