// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (

)

// Icon icon
// swagger:model Icon
type SendSms struct {

	// cover
	Id int64 `json:"id,primary_key"`

	// icon
	Phone string `json:"icon,omitempty"`

	// id
	Code string `json:"code,omitempty"`

	// name
	Ts int64 `json:"ts,omitempty"`

	// order
	Type int64 `json:"type,omitempty"`  //0=注册 =找回密码
}
