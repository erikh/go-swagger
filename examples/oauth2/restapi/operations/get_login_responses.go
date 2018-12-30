// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/erikh/go-swagger/examples/oauth2/models"
)

// GetLoginOKCode is the HTTP code returned for type GetLoginOK
const GetLoginOKCode int = 200

/*GetLoginOK login

swagger:response getLoginOK
*/
type GetLoginOK struct {

	/*
	  In: Body
	*/
	Payload *GetLoginOKBody `json:"body,omitempty"`
}

// NewGetLoginOK creates GetLoginOK with default headers values
func NewGetLoginOK() *GetLoginOK {

	return &GetLoginOK{}
}

// WithPayload adds the payload to the get login o k response
func (o *GetLoginOK) WithPayload(payload *GetLoginOKBody) *GetLoginOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get login o k response
func (o *GetLoginOK) SetPayload(payload *GetLoginOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLoginOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetLoginDefault error

swagger:response getLoginDefault
*/
type GetLoginDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetLoginDefault creates GetLoginDefault with default headers values
func NewGetLoginDefault(code int) *GetLoginDefault {
	if code <= 0 {
		code = 500
	}

	return &GetLoginDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get login default response
func (o *GetLoginDefault) WithStatusCode(code int) *GetLoginDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get login default response
func (o *GetLoginDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get login default response
func (o *GetLoginDefault) WithPayload(payload *models.Error) *GetLoginDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get login default response
func (o *GetLoginDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLoginDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
