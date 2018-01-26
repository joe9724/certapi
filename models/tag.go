// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Msg msg
// swagger:model Msg
type Tag struct {

	// create time
	Id int64 `json:"id"`

	// from
	Icon string `json:"icon"`

	// sub title
	Type string `json:"type"`

	// title
	Status string `json:"status"`

	Name string `json:"name"`
}

// Validate validates this msg
func (m *Msg) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Msg) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Msg) UnmarshalBinary(b []byte) error {
	var res Msg
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
