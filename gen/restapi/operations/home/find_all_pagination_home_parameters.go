// Code generated by go-swagger; DO NOT EDIT.

package home

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewFindAllPaginationHomeParams creates a new FindAllPaginationHomeParams object
// with the default values initialized.
func NewFindAllPaginationHomeParams() FindAllPaginationHomeParams {

	var (
		// initialize parameters with default values

		includeDeletedDataDefault = string("false")
		limitDefault              = int64(10)
		orderByDefault            = string("created_at")
		pageDefault               = int64(1)
		sortByDefault             = string("desc")
	)

	return FindAllPaginationHomeParams{
		IncludeDeletedData: &includeDeletedDataDefault,

		Limit: &limitDefault,

		OrderBy: &orderByDefault,

		Page: &pageDefault,

		SortBy: &sortByDefault,
	}
}

// FindAllPaginationHomeParams contains all the bound params for the find all pagination home operation
// typically these are obtained from a http.Request
//
// swagger:parameters FindAllPaginationHome
type FindAllPaginationHomeParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*flag for identify include deleted data or not
	  In: query
	  Default: "false"
	*/
	IncludeDeletedData *string
	/*default parameter for limit pagination
	  In: query
	  Default: 10
	*/
	Limit *int64
	/*default parameter for order pagination
	  In: query
	  Default: "created_at"
	*/
	OrderBy *string
	/*default parameter for pagination page
	  In: query
	  Default: 1
	*/
	Page *int64
	/*default parameter for sort pagination
	  In: query
	  Default: "desc"
	*/
	SortBy *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewFindAllPaginationHomeParams() beforehand.
func (o *FindAllPaginationHomeParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qIncludeDeletedData, qhkIncludeDeletedData, _ := qs.GetOK("include_deleted_data")
	if err := o.bindIncludeDeletedData(qIncludeDeletedData, qhkIncludeDeletedData, route.Formats); err != nil {
		res = append(res, err)
	}

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	qOrderBy, qhkOrderBy, _ := qs.GetOK("order_by")
	if err := o.bindOrderBy(qOrderBy, qhkOrderBy, route.Formats); err != nil {
		res = append(res, err)
	}

	qPage, qhkPage, _ := qs.GetOK("page")
	if err := o.bindPage(qPage, qhkPage, route.Formats); err != nil {
		res = append(res, err)
	}

	qSortBy, qhkSortBy, _ := qs.GetOK("sort_by")
	if err := o.bindSortBy(qSortBy, qhkSortBy, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindIncludeDeletedData binds and validates parameter IncludeDeletedData from query.
func (o *FindAllPaginationHomeParams) bindIncludeDeletedData(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewFindAllPaginationHomeParams()
		return nil
	}
	o.IncludeDeletedData = &raw

	if err := o.validateIncludeDeletedData(formats); err != nil {
		return err
	}

	return nil
}

// validateIncludeDeletedData carries on validations for parameter IncludeDeletedData
func (o *FindAllPaginationHomeParams) validateIncludeDeletedData(formats strfmt.Registry) error {

	if err := validate.EnumCase("include_deleted_data", "query", *o.IncludeDeletedData, []interface{}{"true", "false"}, true); err != nil {
		return err
	}

	return nil
}

// bindLimit binds and validates parameter Limit from query.
func (o *FindAllPaginationHomeParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewFindAllPaginationHomeParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int64", raw)
	}
	o.Limit = &value

	return nil
}

// bindOrderBy binds and validates parameter OrderBy from query.
func (o *FindAllPaginationHomeParams) bindOrderBy(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewFindAllPaginationHomeParams()
		return nil
	}
	o.OrderBy = &raw

	return nil
}

// bindPage binds and validates parameter Page from query.
func (o *FindAllPaginationHomeParams) bindPage(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewFindAllPaginationHomeParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("page", "query", "int64", raw)
	}
	o.Page = &value

	return nil
}

// bindSortBy binds and validates parameter SortBy from query.
func (o *FindAllPaginationHomeParams) bindSortBy(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewFindAllPaginationHomeParams()
		return nil
	}
	o.SortBy = &raw

	if err := o.validateSortBy(formats); err != nil {
		return err
	}

	return nil
}

// validateSortBy carries on validations for parameter SortBy
func (o *FindAllPaginationHomeParams) validateSortBy(formats strfmt.Registry) error {

	if err := validate.EnumCase("sort_by", "query", *o.SortBy, []interface{}{"asc", "desc"}, true); err != nil {
		return err
	}

	return nil
}
