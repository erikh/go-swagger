// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/erikh/go-swagger/examples/generated/models"
)

// FindPetsByTagsOKCode is the HTTP code returned for type FindPetsByTagsOK
const FindPetsByTagsOKCode int = 200

/*FindPetsByTagsOK successful operation

swagger:response findPetsByTagsOK
*/
type FindPetsByTagsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Pet `json:"body,omitempty"`
}

// NewFindPetsByTagsOK creates FindPetsByTagsOK with default headers values
func NewFindPetsByTagsOK() *FindPetsByTagsOK {

	return &FindPetsByTagsOK{}
}

// WithPayload adds the payload to the find pets by tags o k response
func (o *FindPetsByTagsOK) WithPayload(payload []*models.Pet) *FindPetsByTagsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find pets by tags o k response
func (o *FindPetsByTagsOK) SetPayload(payload []*models.Pet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindPetsByTagsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Pet, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// FindPetsByTagsBadRequestCode is the HTTP code returned for type FindPetsByTagsBadRequest
const FindPetsByTagsBadRequestCode int = 400

/*FindPetsByTagsBadRequest Invalid tag value

swagger:response findPetsByTagsBadRequest
*/
type FindPetsByTagsBadRequest struct {
}

// NewFindPetsByTagsBadRequest creates FindPetsByTagsBadRequest with default headers values
func NewFindPetsByTagsBadRequest() *FindPetsByTagsBadRequest {

	return &FindPetsByTagsBadRequest{}
}

// WriteResponse to the client
func (o *FindPetsByTagsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}
