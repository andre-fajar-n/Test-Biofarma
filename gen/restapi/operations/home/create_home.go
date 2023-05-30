// Code generated by go-swagger; DO NOT EDIT.

package home

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateHomeHandlerFunc turns a function with the right signature into a create home handler
type CreateHomeHandlerFunc func(CreateHomeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateHomeHandlerFunc) Handle(params CreateHomeParams) middleware.Responder {
	return fn(params)
}

// CreateHomeHandler interface for that can handle valid create home params
type CreateHomeHandler interface {
	Handle(CreateHomeParams) middleware.Responder
}

// NewCreateHome creates a new http.Handler for the create home operation
func NewCreateHome(ctx *middleware.Context, handler CreateHomeHandler) *CreateHome {
	return &CreateHome{Context: ctx, Handler: handler}
}

/*
	CreateHome swagger:route POST /v1/home home createHome

# Create

Create new home data
*/
type CreateHome struct {
	Context *middleware.Context
	Handler CreateHomeHandler
}

func (o *CreateHome) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateHomeParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
