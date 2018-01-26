// Code generated by go-swagger; DO NOT EDIT.

package member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "tingtingapi/models"
)

// MemberUploadRecordOKCode is the HTTP code returned for type MemberUploadRecordOK
const MemberUploadRecordOKCode int = 200

/*MemberUploadRecordOK 上传成功，返回成功信息

swagger:response memberUploadRecordOK
*/
type MemberUploadRecordOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse2008 `json:"body,omitempty"`
}

// NewMemberUploadRecordOK creates MemberUploadRecordOK with default headers values
func NewMemberUploadRecordOK() *MemberUploadRecordOK {
	return &MemberUploadRecordOK{}
}

// WithPayload adds the payload to the member upload record o k response
func (o *MemberUploadRecordOK) WithPayload(payload *models.InlineResponse2008) *MemberUploadRecordOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the member upload record o k response
func (o *MemberUploadRecordOK) SetPayload(payload *models.InlineResponse2008) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *MemberUploadRecordOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrMemberUploadRecordDefault error

swagger:response memberUploadRecordDefault
*/
type NrMemberUploadRecordDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrMemberUploadRecordDefault creates NrMemberUploadRecordDefault with default headers values
func NewNrMemberUploadRecordDefault(code int) *NrMemberUploadRecordDefault {
	if code <= 0 {
		code = 500
	}

	return &NrMemberUploadRecordDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the member upload record default response
func (o *NrMemberUploadRecordDefault) WithStatusCode(code int) *NrMemberUploadRecordDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the member upload record default response
func (o *NrMemberUploadRecordDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the member upload record default response
func (o *NrMemberUploadRecordDefault) WithPayload(payload *models.Error) *NrMemberUploadRecordDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the member upload record default response
func (o *NrMemberUploadRecordDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrMemberUploadRecordDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}