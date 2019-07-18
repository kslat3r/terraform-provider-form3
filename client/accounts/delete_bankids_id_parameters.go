// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteBankidsIDParams creates a new DeleteBankidsIDParams object
// with the default values initialized.
func NewDeleteBankidsIDParams() *DeleteBankidsIDParams {
	var ()
	return &DeleteBankidsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBankidsIDParamsWithTimeout creates a new DeleteBankidsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteBankidsIDParamsWithTimeout(timeout time.Duration) *DeleteBankidsIDParams {
	var ()
	return &DeleteBankidsIDParams{

		timeout: timeout,
	}
}

// NewDeleteBankidsIDParamsWithContext creates a new DeleteBankidsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteBankidsIDParamsWithContext(ctx context.Context) *DeleteBankidsIDParams {
	var ()
	return &DeleteBankidsIDParams{

		Context: ctx,
	}
}

// NewDeleteBankidsIDParamsWithHTTPClient creates a new DeleteBankidsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteBankidsIDParamsWithHTTPClient(client *http.Client) *DeleteBankidsIDParams {
	var ()
	return &DeleteBankidsIDParams{
		HTTPClient: client,
	}
}

/*DeleteBankidsIDParams contains all the parameters to send to the API endpoint
for the delete bankids ID operation typically these are written to a http.Request
*/
type DeleteBankidsIDParams struct {

	/*ID
	  BankId Id

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

// WithTimeout adds the timeout to the delete bankids ID params
func (o *DeleteBankidsIDParams) WithTimeout(timeout time.Duration) *DeleteBankidsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete bankids ID params
func (o *DeleteBankidsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete bankids ID params
func (o *DeleteBankidsIDParams) WithContext(ctx context.Context) *DeleteBankidsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete bankids ID params
func (o *DeleteBankidsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete bankids ID params
func (o *DeleteBankidsIDParams) WithHTTPClient(client *http.Client) *DeleteBankidsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete bankids ID params
func (o *DeleteBankidsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete bankids ID params
func (o *DeleteBankidsIDParams) WithID(id strfmt.UUID) *DeleteBankidsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete bankids ID params
func (o *DeleteBankidsIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithVersion adds the version to the delete bankids ID params
func (o *DeleteBankidsIDParams) WithVersion(version int64) *DeleteBankidsIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete bankids ID params
func (o *DeleteBankidsIDParams) SetVersion(version int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBankidsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
