// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SuccessGetLocationResourceSetsItemsResourcesItemsPoint success get location resource sets items resources items point
//
// swagger:model successGetLocationResourceSetsItemsResourcesItemsPoint
type SuccessGetLocationResourceSetsItemsResourcesItemsPoint struct {

	// coordinates
	Coordinates []float64 `json:"coordinates"`
}

// Validate validates this success get location resource sets items resources items point
func (m *SuccessGetLocationResourceSetsItemsResourcesItemsPoint) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this success get location resource sets items resources items point based on context it is used
func (m *SuccessGetLocationResourceSetsItemsResourcesItemsPoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SuccessGetLocationResourceSetsItemsResourcesItemsPoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SuccessGetLocationResourceSetsItemsResourcesItemsPoint) UnmarshalBinary(b []byte) error {
	var res SuccessGetLocationResourceSetsItemsResourcesItemsPoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
