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

// SuccessFindAllPaginationAllOf1Metadata success find all pagination all of1 metadata
//
// swagger:model successFindAllPaginationAllOf1Metadata
type SuccessFindAllPaginationAllOf1Metadata struct {
	Metadata

	CustomFields
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SuccessFindAllPaginationAllOf1Metadata) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Metadata
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Metadata = aO0

	// AO1
	var aO1 CustomFields
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.CustomFields = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SuccessFindAllPaginationAllOf1Metadata) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Metadata)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.CustomFields)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this success find all pagination all of1 metadata
func (m *SuccessFindAllPaginationAllOf1Metadata) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Metadata
	if err := m.Metadata.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with CustomFields
	if err := m.CustomFields.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this success find all pagination all of1 metadata based on the context it is used
func (m *SuccessFindAllPaginationAllOf1Metadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Metadata
	if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with CustomFields
	if err := m.CustomFields.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SuccessFindAllPaginationAllOf1Metadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SuccessFindAllPaginationAllOf1Metadata) UnmarshalBinary(b []byte) error {
	var res SuccessFindAllPaginationAllOf1Metadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
