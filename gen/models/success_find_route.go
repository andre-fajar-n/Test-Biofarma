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

// SuccessFindRoute success find route
//
// swagger:model successFindRoute
type SuccessFindRoute struct {
	SuccessCreateAllOf0

	SuccessFindRouteAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SuccessFindRoute) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 SuccessCreateAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.SuccessCreateAllOf0 = aO0

	// AO1
	var aO1 SuccessFindRouteAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.SuccessFindRouteAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SuccessFindRoute) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.SuccessCreateAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.SuccessFindRouteAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this success find route
func (m *SuccessFindRoute) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SuccessCreateAllOf0
	if err := m.SuccessCreateAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with SuccessFindRouteAllOf1
	if err := m.SuccessFindRouteAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this success find route based on the context it is used
func (m *SuccessFindRoute) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SuccessCreateAllOf0
	if err := m.SuccessCreateAllOf0.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with SuccessFindRouteAllOf1
	if err := m.SuccessFindRouteAllOf1.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SuccessFindRoute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SuccessFindRoute) UnmarshalBinary(b []byte) error {
	var res SuccessFindRoute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}