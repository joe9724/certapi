// Code generated by go-swagger; DO NOT EDIT.

package album

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// AlbumDetailURL generates an URL for the album detail operation
type AlbumDetailURL struct {
	AlbumID  *string
	Imei     *string
	MemberID *int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *AlbumDetailURL) WithBasePath(bp string) *AlbumDetailURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *AlbumDetailURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *AlbumDetailURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/album/detail"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/nanjingyouzi/TingtingApi/1.0.0"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var albumID string
	if o.AlbumID != nil {
		albumID = *o.AlbumID
	}
	if albumID != "" {
		qs.Set("albumId", albumID)
	}

	var imei string
	if o.Imei != nil {
		imei = *o.Imei
	}
	if imei != "" {
		qs.Set("imei", imei)
	}

	var memberID string
	if o.MemberID != nil {
		memberID = swag.FormatInt64(*o.MemberID)
	}
	if memberID != "" {
		qs.Set("memberId", memberID)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *AlbumDetailURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *AlbumDetailURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *AlbumDetailURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on AlbumDetailURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on AlbumDetailURL")
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
func (o *AlbumDetailURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
