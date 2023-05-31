// Code generated by go-swagger; DO NOT EDIT.

package home

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteHomeHandlerFunc turns a function with the right signature into a delete home handler
type DeleteHomeHandlerFunc func(DeleteHomeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteHomeHandlerFunc) Handle(params DeleteHomeParams) middleware.Responder {
	return fn(params)
}

// DeleteHomeHandler interface for that can handle valid delete home params
type DeleteHomeHandler interface {
	Handle(DeleteHomeParams) middleware.Responder
}

// NewDeleteHome creates a new http.Handler for the delete home operation
func NewDeleteHome(ctx *middleware.Context, handler DeleteHomeHandler) *DeleteHome {
	return &DeleteHome{Context: ctx, Handler: handler}
}

/*
	DeleteHome swagger:route DELETE /v1/home/{home_id} home deleteHome

# Delete

Delete home data
*/
type DeleteHome struct {
	Context *middleware.Context
	Handler DeleteHomeHandler
}

func (o *DeleteHome) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteHomeParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}