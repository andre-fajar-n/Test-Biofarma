// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HomeWithoutTrackTime home without track time
//
// swagger:model homeWithoutTrackTime
type HomeWithoutTrackTime struct {
	ModelIdentifier

	HomeData
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HomeWithoutTrackTime) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ModelIdentifier
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ModelIdentifier = aO0

	// AO1
	var aO1 HomeData
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.HomeData = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HomeWithoutTrackTime) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ModelIdentifier)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.HomeData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this home without track time
func (m *HomeWithoutTrackTime) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ModelIdentifier
	if err := m.ModelIdentifier.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with HomeData
	if err := m.HomeData.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this home without track time based on the context it is used
func (m *HomeWithoutTrackTime) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ModelIdentifier
	if err := m.ModelIdentifier.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with HomeData
	if err := m.HomeData.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *HomeWithoutTrackTime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HomeWithoutTrackTime) UnmarshalBinary(b []byte) error {
	var res HomeWithoutTrackTime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
