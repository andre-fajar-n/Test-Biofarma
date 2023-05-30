// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"biofarma/gen/models"
)

// FindMyUserDataOKCode is the HTTP code returned for type FindMyUserDataOK
const FindMyUserDataOKCode int = 200

/*
FindMyUserDataOK Success fetch data

swagger:response findMyUserDataOK
*/
type FindMyUserDataOK struct {

	/*
	  In: Body
	*/
	Payload *models.MyUserData `json:"body,omitempty"`
}

// NewFindMyUserDataOK creates FindMyUserDataOK with default headers values
func NewFindMyUserDataOK() *FindMyUserDataOK {

	return &FindMyUserDataOK{}
}

// WithPayload adds the payload to the find my user data o k response
func (o *FindMyUserDataOK) WithPayload(payload *models.MyUserData) *FindMyUserDataOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find my user data o k response
func (o *FindMyUserDataOK) SetPayload(payload *models.MyUserData) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindMyUserDataOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
FindMyUserDataDefault Server Error

swagger:response findMyUserDataDefault
*/
type FindMyUserDataDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewFindMyUserDataDefault creates FindMyUserDataDefault with default headers values
func NewFindMyUserDataDefault(code int) *FindMyUserDataDefault {
	if code <= 0 {
		code = 500
	}

	return &FindMyUserDataDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the find my user data default response
func (o *FindMyUserDataDefault) WithStatusCode(code int) *FindMyUserDataDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the find my user data default response
func (o *FindMyUserDataDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the find my user data default response
func (o *FindMyUserDataDefault) WithPayload(payload *models.Error) *FindMyUserDataDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find my user data default response
func (o *FindMyUserDataDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindMyUserDataDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
