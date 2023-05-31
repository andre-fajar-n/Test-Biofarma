// Code generated by go-swagger; DO NOT EDIT.

package home

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// FindOneHomeURL generates an URL for the find one home operation
type FindOneHomeURL struct {
	HomeID uint64

	IncludeDeletedData *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *FindOneHomeURL) WithBasePath(bp string) *FindOneHomeURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *FindOneHomeURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *FindOneHomeURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/v1/home/{home_id}"

	homeID := swag.FormatUint64(o.HomeID)
	if homeID != "" {
		_path = strings.Replace(_path, "{home_id}", homeID, -1)
	} else {
		return nil, errors.New("homeId is required on FindOneHomeURL")
	}

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

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *FindOneHomeURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *FindOneHomeURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *FindOneHomeURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on FindOneHomeURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on FindOneHomeURL")
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
func (o *FindOneHomeURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
