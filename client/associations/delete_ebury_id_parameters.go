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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteEburyIDParams creates a new DeleteEburyIDParams object
// with the default values initialized.
func NewDeleteEburyIDParams() *DeleteEburyIDParams {
	var ()
	return &DeleteEburyIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteEburyIDParamsWithTimeout creates a new DeleteEburyIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteEburyIDParamsWithTimeout(timeout time.Duration) *DeleteEburyIDParams {
	var ()
	return &DeleteEburyIDParams{

		timeout: timeout,
	}
}

// NewDeleteEburyIDParamsWithContext creates a new DeleteEburyIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteEburyIDParamsWithContext(ctx context.Context) *DeleteEburyIDParams {
	var ()
	return &DeleteEburyIDParams{

		Context: ctx,
	}
}

// NewDeleteEburyIDParamsWithHTTPClient creates a new DeleteEburyIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteEburyIDParamsWithHTTPClient(client *http.Client) *DeleteEburyIDParams {
	var ()
	return &DeleteEburyIDParams{
		HTTPClient: client,
	}
}

/*DeleteEburyIDParams contains all the parameters to send to the API endpoint
for the delete ebury ID operation typically these are written to a http.Request
*/
type DeleteEburyIDParams struct {

	/*ID
	  Association Id

	*/
	ID strfmt.UUID
	/*Version
	  Version

	*/
	Version int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete ebury ID params
func (o *DeleteEburyIDParams) WithTimeout(timeout time.Duration) *DeleteEburyIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete ebury ID params
func (o *DeleteEburyIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete ebury ID params
func (o *DeleteEburyIDParams) WithContext(ctx context.Context) *DeleteEburyIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete ebury ID params
func (o *DeleteEburyIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete ebury ID params
func (o *DeleteEburyIDParams) WithHTTPClient(client *http.Client) *DeleteEburyIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete ebury ID params
func (o *DeleteEburyIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete ebury ID params
func (o *DeleteEburyIDParams) WithID(id strfmt.UUID) *DeleteEburyIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete ebury ID params
func (o *DeleteEburyIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithVersion adds the version to the delete ebury ID params
func (o *DeleteEburyIDParams) WithVersion(version int64) *DeleteEburyIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete ebury ID params
func (o *DeleteEburyIDParams) SetVersion(version int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteEburyIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// query param version
	qrVersion := o.Version
	qVersion := swag.FormatInt64(qrVersion)
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
