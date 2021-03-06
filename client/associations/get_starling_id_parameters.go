// Code generated by go-swagger; DO NOT EDIT.

package associations

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
)

// NewGetStarlingIDParams creates a new GetStarlingIDParams object
// with the default values initialized.
func NewGetStarlingIDParams() *GetStarlingIDParams {
	var ()
	return &GetStarlingIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStarlingIDParamsWithTimeout creates a new GetStarlingIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStarlingIDParamsWithTimeout(timeout time.Duration) *GetStarlingIDParams {
	var ()
	return &GetStarlingIDParams{

		timeout: timeout,
	}
}

// NewGetStarlingIDParamsWithContext creates a new GetStarlingIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStarlingIDParamsWithContext(ctx context.Context) *GetStarlingIDParams {
	var ()
	return &GetStarlingIDParams{

		Context: ctx,
	}
}

// NewGetStarlingIDParamsWithHTTPClient creates a new GetStarlingIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStarlingIDParamsWithHTTPClient(client *http.Client) *GetStarlingIDParams {
	var ()
	return &GetStarlingIDParams{
		HTTPClient: client,
	}
}

/*GetStarlingIDParams contains all the parameters to send to the API endpoint
for the get starling ID operation typically these are written to a http.Request
*/
type GetStarlingIDParams struct {

	/*ID
	  Association Id

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get starling ID params
func (o *GetStarlingIDParams) WithTimeout(timeout time.Duration) *GetStarlingIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get starling ID params
func (o *GetStarlingIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get starling ID params
func (o *GetStarlingIDParams) WithContext(ctx context.Context) *GetStarlingIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get starling ID params
func (o *GetStarlingIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get starling ID params
func (o *GetStarlingIDParams) WithHTTPClient(client *http.Client) *GetStarlingIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get starling ID params
func (o *GetStarlingIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get starling ID params
func (o *GetStarlingIDParams) WithID(id strfmt.UUID) *GetStarlingIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get starling ID params
func (o *GetStarlingIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetStarlingIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
