// Code generated by go-swagger; DO NOT EDIT.

package route

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"biofarma/gen/models"
)

// FindRouteOKCode is the HTTP code returned for type FindRouteOK
const FindRouteOKCode int = 200

/*
FindRouteOK Success get data

swagger:response findRouteOK
*/
type FindRouteOK struct {

	/*
	  In: Body
	*/
	Payload *models.SuccessFindRoute `json:"body,omitempty"`
}

// NewFindRouteOK creates FindRouteOK with default headers values
func NewFindRouteOK() *FindRouteOK {

	return &FindRouteOK{}
}

// WithPayload adds the payload to the find route o k response
func (o *FindRouteOK) WithPayload(payload *models.SuccessFindRoute) *FindRouteOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find route o k response
func (o *FindRouteOK) SetPayload(payload *models.SuccessFindRoute) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindRouteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
FindRouteDefault Server Error

swagger:response findRouteDefault
*/
type FindRouteDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewFindRouteDefault creates FindRouteDefault with default headers values
func NewFindRouteDefault(code int) *FindRouteDefault {
	if code <= 0 {
		code = 500
	}

	return &FindRouteDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the find route default response
func (o *FindRouteDefault) WithStatusCode(code int) *FindRouteDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the find route default response
func (o *FindRouteDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the find route default response
func (o *FindRouteDefault) WithPayload(payload *models.Error) *FindRouteDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find route default response
func (o *FindRouteDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindRouteDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}