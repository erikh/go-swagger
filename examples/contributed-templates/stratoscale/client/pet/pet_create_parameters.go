// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/erikh/go-swagger/examples/contributed-templates/stratoscale/models"
)

// NewPetCreateParams creates a new PetCreateParams object
// with the default values initialized.
func NewPetCreateParams() *PetCreateParams {
	var ()
	return &PetCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPetCreateParamsWithTimeout creates a new PetCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPetCreateParamsWithTimeout(timeout time.Duration) *PetCreateParams {
	var ()
	return &PetCreateParams{

		timeout: timeout,
	}
}

// NewPetCreateParamsWithContext creates a new PetCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPetCreateParamsWithContext(ctx context.Context) *PetCreateParams {
	var ()
	return &PetCreateParams{

		Context: ctx,
	}
}

// NewPetCreateParamsWithHTTPClient creates a new PetCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPetCreateParamsWithHTTPClient(client *http.Client) *PetCreateParams {
	var ()
	return &PetCreateParams{
		HTTPClient: client,
	}
}

/*PetCreateParams contains all the parameters to send to the API endpoint
for the pet create operation typically these are written to a http.Request
*/
type PetCreateParams struct {

	/*Body
	  Pet object that needs to be added to the store

	*/
	Body *models.Pet

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pet create params
func (o *PetCreateParams) WithTimeout(timeout time.Duration) *PetCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pet create params
func (o *PetCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pet create params
func (o *PetCreateParams) WithContext(ctx context.Context) *PetCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pet create params
func (o *PetCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pet create params
func (o *PetCreateParams) WithHTTPClient(client *http.Client) *PetCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pet create params
func (o *PetCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pet create params
func (o *PetCreateParams) WithBody(body *models.Pet) *PetCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pet create params
func (o *PetCreateParams) SetBody(body *models.Pet) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PetCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
