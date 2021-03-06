// Code generated by go-swagger; DO NOT EDIT.

package album

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "tingtingapi/models"
)

// AlbumBuyOKCode is the HTTP code returned for type AlbumBuyOK
const AlbumBuyOKCode int = 200

/*AlbumBuyOK 购买专辑返回结果

swagger:response albumBuyOK
*/
type AlbumBuyOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse2007 `json:"body,omitempty"`
}

// NewAlbumBuyOK creates AlbumBuyOK with default headers values
func NewAlbumBuyOK() *AlbumBuyOK {
	return &AlbumBuyOK{}
}

// WithPayload adds the payload to the album buy o k response
func (o *AlbumBuyOK) WithPayload(payload *models.InlineResponse2007) *AlbumBuyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the album buy o k response
func (o *AlbumBuyOK) SetPayload(payload *models.InlineResponse2007) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AlbumBuyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*AlbumBuyDefault error

swagger:response albumBuyDefault
*/
type AlbumBuyDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAlbumBuyDefault creates AlbumBuyDefault with default headers values
func NewAlbumBuyDefault(code int) *AlbumBuyDefault {
	if code <= 0 {
		code = 500
	}

	return &AlbumBuyDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the album buy default response
func (o *AlbumBuyDefault) WithStatusCode(code int) *AlbumBuyDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the album buy default response
func (o *AlbumBuyDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the album buy default response
func (o *AlbumBuyDefault) WithPayload(payload *models.Error) *AlbumBuyDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the album buy default response
func (o *AlbumBuyDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AlbumBuyDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
