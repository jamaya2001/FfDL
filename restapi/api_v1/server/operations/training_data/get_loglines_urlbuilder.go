// Code generated by go-swagger; DO NOT EDIT.

package training_data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// GetLoglinesURL generates an URL for the get loglines operation
type GetLoglinesURL struct {
	ModelID string

	Pagesize   *int32
	Pos        *int64
	SearchType *string
	SinceTime  *string
	Version    *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetLoglinesURL) WithBasePath(bp string) *GetLoglinesURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetLoglinesURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetLoglinesURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/v1/logs/{model_id}/loglines"

	modelID := o.ModelID
	if modelID != "" {
		_path = strings.Replace(_path, "{model_id}", modelID, -1)
	} else {
		return nil, errors.New("ModelID is required on GetLoglinesURL")
	}
	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var pagesize string
	if o.Pagesize != nil {
		pagesize = swag.FormatInt32(*o.Pagesize)
	}
	if pagesize != "" {
		qs.Set("pagesize", pagesize)
	}

	var pos string
	if o.Pos != nil {
		pos = swag.FormatInt64(*o.Pos)
	}
	if pos != "" {
		qs.Set("pos", pos)
	}

	var searchType string
	if o.SearchType != nil {
		searchType = *o.SearchType
	}
	if searchType != "" {
		qs.Set("searchType", searchType)
	}

	var sinceTime string
	if o.SinceTime != nil {
		sinceTime = *o.SinceTime
	}
	if sinceTime != "" {
		qs.Set("since_time", sinceTime)
	}

	var version string
	if o.Version != nil {
		version = *o.Version
	}
	if version != "" {
		qs.Set("version", version)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetLoglinesURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetLoglinesURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetLoglinesURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetLoglinesURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetLoglinesURL")
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
func (o *GetLoglinesURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
