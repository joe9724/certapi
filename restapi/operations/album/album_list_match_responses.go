// Code generated by go-swagger; DO NOT EDIT.

package album

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "tingtingapi/models"
)

// AlbumListMatchOKCode is the HTTP code returned for type AlbumListMatchOK
const AlbumListMatchOKCode int = 200

/*AlbumListMatchOK 登录成功，返回登录信息

swagger:response albumListMatchOK
*/
type AlbumListMatchOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse2006 `json:"body,omitempty"`
}

// NewAlbumListMatchOK creates AlbumListMatchOK with default headers values
func NewAlbumListMatchOK() *AlbumListMatchOK {
	return &AlbumListMatchOK{}
}

// WithPayload adds the payload to the album list match o k response
func (o *AlbumListMatchOK) WithPayload(payload *models.InlineResponse2006) *AlbumListMatchOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the album list match o k response
func (o *AlbumListMatchOK) SetPayload(payload *models.InlineResponse2006) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AlbumListMatchOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*AlbumListMatchDefault error

swagger:response albumListMatchDefault
*/
type AlbumListMatchDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAlbumListMatchDefault creates AlbumListMatchDefault with default headers values
func NewAlbumListMatchDefault(code int) *AlbumListMatchDefault {
	if code <= 0 {
		code = 500
	}

	return &AlbumListMatchDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the album list match default response
func (o *AlbumListMatchDefault) WithStatusCode(code int) *AlbumListMatchDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the album list match default response
func (o *AlbumListMatchDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the album list match default response
func (o *AlbumListMatchDefault) WithPayload(payload *models.Error) *AlbumListMatchDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the album list match default response
func (o *AlbumListMatchDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AlbumListMatchDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
