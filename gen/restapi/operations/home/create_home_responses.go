// Code generated by go-swagger; DO NOT EDIT.

package home

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"biofarma/gen/models"
)

// CreateHomeCreatedCode is the HTTP code returned for type CreateHomeCreated
const CreateHomeCreatedCode int = 201

/*
CreateHomeCreated Success create

swagger:response createHomeCreated
*/
type CreateHomeCreated struct {

	/*
	  In: Body
	*/
	Payload *models.SuccessCreate `json:"body,omitempty"`
}

// NewCreateHomeCreated creates CreateHomeCreated with default headers values
func NewCreateHomeCreated() *CreateHomeCreated {

	return &CreateHomeCreated{}
}

// WithPayload adds the payload to the create home created response
func (o *CreateHomeCreated) WithPayload(payload *models.SuccessCreate) *CreateHomeCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create home created response
func (o *CreateHomeCreated) SetPayload(payload *models.SuccessCreate) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateHomeCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
CreateHomeDefault Server Error

swagger:response createHomeDefault
*/
type CreateHomeDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateHomeDefault creates CreateHomeDefault with default headers values
func NewCreateHomeDefault(code int) *CreateHomeDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateHomeDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create home default response
func (o *CreateHomeDefault) WithStatusCode(code int) *CreateHomeDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create home default response
func (o *CreateHomeDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create home default response
func (o *CreateHomeDefault) WithPayload(payload *models.Error) *CreateHomeDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create home default response
func (o *CreateHomeDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateHomeDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
