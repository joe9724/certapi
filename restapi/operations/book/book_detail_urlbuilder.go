// Code generated by go-swagger; DO NOT EDIT.

package book

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
)

// BookDetailURL generates an URL for the book detail operation
type BookDetailURL struct {
	Membername *string
	BookID     *string
	Client     *string
	Imei       *string
	Version    *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *BookDetailURL) WithBasePath(bp string) *BookDetailURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *BookDetailURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *BookDetailURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/book/detail"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/nanjingyouzi/TingtingApi/1.0.0"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var membername string
	if o.Membername != nil {
		membername = *o.Membername
	}
	if membername != "" {
		qs.Set("Membername", membername)
	}

	var bookID string
	if o.BookID != nil {
		bookID = *o.BookID
	}
	if bookID != "" {
		qs.Set("bookId", bookID)
	}

	var client string
	if o.Client != nil {
		client = *o.Client
	}
	if client != "" {
		qs.Set("client", client)
	}

	var imei string
	if o.Imei != nil {
		imei = *o.Imei
	}
	if imei != "" {
		qs.Set("imei", imei)
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
func (o *BookDetailURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *BookDetailURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *BookDetailURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on BookDetailURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on BookDetailURL")
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
func (o *BookDetailURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
