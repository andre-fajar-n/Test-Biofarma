// Code generated by go-swagger; DO NOT EDIT.

package home

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// FindAllPaginationHomeURL generates an URL for the find all pagination home operation
type FindAllPaginationHomeURL struct {
	IncludeDeletedData *string
	Limit              *int64
	OrderBy            *string
	Page               *int64
	SortBy             *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *FindAllPaginationHomeURL) WithBasePath(bp string) *FindAllPaginationHomeURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *FindAllPaginationHomeURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *FindAllPaginationHomeURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/v1/home"

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var includeDeletedDataQ string
	if o.IncludeDeletedData != nil {
		includeDeletedDataQ = *o.IncludeDeletedData
	}
	if includeDeletedDataQ != "" {
		qs.Set("include_deleted_data", includeDeletedDataQ)
	}

	var limitQ string
	if o.Limit != nil {
		limitQ = swag.FormatInt64(*o.Limit)
	}
	if limitQ != "" {
		qs.Set("limit", limitQ)
	}

	var orderByQ string
	if o.OrderBy != nil {
		orderByQ = *o.OrderBy
	}
	if orderByQ != "" {
		qs.Set("order_by", orderByQ)
	}

	var pageQ string
	if o.Page != nil {
		pageQ = swag.FormatInt64(*o.Page)
	}
	if pageQ != "" {
		qs.Set("page", pageQ)
	}

	var sortByQ string
	if o.SortBy != nil {
		sortByQ = *o.SortBy
	}
	if sortByQ != "" {
		qs.Set("sort_by", sortByQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *FindAllPaginationHomeURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *FindAllPaginationHomeURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *FindAllPaginationHomeURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on FindAllPaginationHomeURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on FindAllPaginationHomeURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *FindAllPaginationHomeURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
