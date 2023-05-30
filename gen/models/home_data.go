// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HomeData home data
//
// swagger:model homeData
type HomeData struct {

	// address
	Address string `json:"address,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// latitude
	Latitude float64 `json:"latitude,omitempty"`

	// longitude
	Longitude float64 `json:"longitude,omitempty"`

	// regency
	Regency string `json:"regency,omitempty"`

	// subdistrict
	Subdistrict string `json:"subdistrict,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this home data
func (m *HomeData) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this home data based on context it is used
func (m *HomeData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HomeData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HomeData) UnmarshalBinary(b []byte) error {
	var res HomeData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
