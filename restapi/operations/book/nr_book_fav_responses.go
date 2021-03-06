// Code generated by go-swagger; DO NOT EDIT.

package book

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "tingtingapi/models"
)

// BookFavOKCode is the HTTP code returned for type BookFavOK
const BookFavOKCode int = 200

/*BookFavOK 购买专辑返回结果

swagger:response bookFavOK
*/
type BookFavOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse2003 `json:"body,omitempty"`
}

// NewBookFavOK creates BookFavOK with default headers values
func NewBookFavOK() *BookFavOK {
	return &BookFavOK{}
}

// WithPayload adds the payload to the book fav o k response
func (o *BookFavOK) WithPayload(payload *models.InlineResponse2003) *BookFavOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the book fav o k response
func (o *BookFavOK) SetPayload(payload *models.InlineResponse2003) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BookFavOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrBookFavDefault error

swagger:response bookFavDefault
*/
type NrBookFavDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrBookFavDefault creates NrBookFavDefault with default headers values
func NewNrBookFavDefault(code int) *NrBookFavDefault {
	if code <= 0 {
		code = 500
	}

	return &NrBookFavDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the book fav default response
func (o *NrBookFavDefault) WithStatusCode(code int) *NrBookFavDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the book fav default response
func (o *NrBookFavDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the book fav default response
func (o *NrBookFavDefault) WithPayload(payload *models.Error) *NrBookFavDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the book fav default response
func (o *NrBookFavDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrBookFavDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
