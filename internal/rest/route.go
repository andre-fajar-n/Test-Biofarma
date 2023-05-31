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

		api.HomeFindOneHomeHandler = home.FindOneHomeHandlerFunc(func(fohp home.FindOneHomeParams) middleware.Responder {
			data, err := apiHandler.FindOneHome(context.Background(), &fohp)
			if err != nil {
				errRes := rt.GetError(err)
				return home.NewFindOneHomeDefault(int(errRes.Code())).WithPayload(&models.Error{
					Code:    int64(errRes.Code()),
					Message: errRes.Error(),
				})
			}

			return home.NewFindOneHomeOK().WithPayload(&models.SuccessFindOne{
				SuccessCreateAllOf0: models.SuccessCreateAllOf0{
					Message: "Success get home data",
				},
				SuccessFindOneAllOf1: *data,
			})
		})

		api.HomeDeleteHomeHandler = home.DeleteHomeHandlerFunc(func(dhp home.DeleteHomeParams) middleware.Responder {
			err := apiHandler.DeleteHome(context.Background(), &dhp)
			if err != nil {
				errRes := rt.GetError(err)
				return home.NewDeleteHomeDefault(int(errRes.Code())).WithPayload(&models.Error{
					Code:    int64(errRes.Code()),
					Message: errRes.Error(),
				})
			}

			return home.NewDeleteHomeOK().WithPayload(&models.Success{
				Message: "Success delete home",
			})
		})
	}
}
