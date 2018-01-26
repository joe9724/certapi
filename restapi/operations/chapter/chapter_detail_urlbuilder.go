// Code generated by go-swagger; DO NOT EDIT.

package chapter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// ChapterDetailURL generates an URL for the chapter detail operation
type ChapterDetailURL struct {
	ChapterID *string
	Userid    *int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ChapterDetailURL) WithBasePath(bp string) *ChapterDetailURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ChapterDetailURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *ChapterDetailURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/chapter/detail"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/nanjingyouzi/TingtingApi/1.0.0"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var chapterID string
	if o.ChapterID != nil {
		chapterID = *o.ChapterID
	}
	if chapterID != "" {
		qs.Set("chapterId", chapterID)
	}

	var userid string
	if o.Userid != nil {
		userid = swag.FormatInt64(*o.Userid)
	}
	if userid != "" {
		qs.Set("userid", userid)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *ChapterDetailURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *ChapterDetailURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *ChapterDetailURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on ChapterDetailURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on ChapterDetailURL")
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
func (o *ChapterDetailURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}