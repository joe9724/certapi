// Code generated by go-swagger; DO NOT EDIT.

package chapter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "tingtingapi/models"
)

// ChapterFavListOKCode is the HTTP code returned for type ChapterFavListOK
const ChapterFavListOKCode int = 200

/*ChapterFavListOK 登录成功，返回登录信息

swagger:response chapterFavListOK
*/
type ChapterFavListOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse20015 `json:"body,omitempty"`
}

// NewChapterFavListOK creates ChapterFavListOK with default headers values
func NewChapterFavListOK() *ChapterFavListOK {
	return &ChapterFavListOK{}
}

// WithPayload adds the payload to the chapter fav list o k response
func (o *ChapterFavListOK) WithPayload(payload *models.InlineResponse20015) *ChapterFavListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the chapter fav list o k response
func (o *ChapterFavListOK) SetPayload(payload *models.InlineResponse20015) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ChapterFavListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ChapterFavListDefault error

swagger:response chapterFavListDefault
*/
type ChapterFavListDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewChapterFavListDefault creates ChapterFavListDefault with default headers values
func NewChapterFavListDefault(code int) *ChapterFavListDefault {
	if code <= 0 {
		code = 500
	}

	return &ChapterFavListDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the chapter fav list default response
func (o *ChapterFavListDefault) WithStatusCode(code int) *ChapterFavListDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the chapter fav list default response
func (o *ChapterFavListDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the chapter fav list default response
func (o *ChapterFavListDefault) WithPayload(payload *models.Error) *ChapterFavListDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the chapter fav list default response
func (o *ChapterFavListDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ChapterFavListDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}