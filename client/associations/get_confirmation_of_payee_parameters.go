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

// NewGetConfirmationOfPayeeParams creates a new GetConfirmationOfPayeeParams object
// with the default values initialized.
func NewGetConfirmationOfPayeeParams() *GetConfirmationOfPayeeParams {
	var ()
	return &GetConfirmationOfPayeeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetConfirmationOfPayeeParamsWithTimeout creates a new GetConfirmationOfPayeeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetConfirmationOfPayeeParamsWithTimeout(timeout time.Duration) *GetConfirmationOfPayeeParams {
	var ()
	return &GetConfirmationOfPayeeParams{

		timeout: timeout,
	}
}

// NewGetConfirmationOfPayeeParamsWithContext creates a new GetConfirmationOfPayeeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetConfirmationOfPayeeParamsWithContext(ctx context.Context) *GetConfirmationOfPayeeParams {
	var ()
	return &GetConfirmationOfPayeeParams{

		Context: ctx,
	}
}

// NewGetConfirmationOfPayeeParamsWithHTTPClient creates a new GetConfirmationOfPayeeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetConfirmationOfPayeeParamsWithHTTPClient(client *http.Client) *GetConfirmationOfPayeeParams {
	var ()
	return &GetConfirmationOfPayeeParams{
		HTTPClient: client,
	}
}

/*GetConfirmationOfPayeeParams contains all the parameters to send to the API endpoint
for the get confirmation of payee operation typically these are written to a http.Request
*/
type GetConfirmationOfPayeeParams struct {

	/*FilterOrganisationID
	  Organisation id

	*/
	FilterOrganisationID []strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get confirmation of payee params
func (o *GetConfirmationOfPayeeParams) WithTimeout(timeout time.Duration) *GetConfirmationOfPayeeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get confirmation of payee params
func (o *GetConfirmationOfPayeeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get confirmation of payee params
func (o *GetConfirmationOfPayeeParams) WithContext(ctx context.Context) *GetConfirmationOfPayeeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get confirmation of payee params
func (o *GetConfirmationOfPayeeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get confirmation of payee params
func (o *GetConfirmationOfPayeeParams) WithHTTPClient(client *http.Client) *GetConfirmationOfPayeeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get confirmation of payee params
func (o *GetConfirmationOfPayeeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilterOrganisationID adds the filterOrganisationID to the get confirmation of payee params
func (o *GetConfirmationOfPayeeParams) WithFilterOrganisationID(filterOrganisationID []strfmt.UUID) *GetConfirmationOfPayeeParams {
	o.SetFilterOrganisationID(filterOrganisationID)
	return o
}

// SetFilterOrganisationID adds the filterOrganisationId to the get confirmation of payee params
func (o *GetConfirmationOfPayeeParams) SetFilterOrganisationID(filterOrganisationID []strfmt.UUID) {
	o.FilterOrganisationID = filterOrganisationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetConfirmationOfPayeeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	var valuesFilterOrganisationID []string
	for _, v := range o.FilterOrganisationID {
		valuesFilterOrganisationID = append(valuesFilterOrganisationID, v.String())
	}

	joinedFilterOrganisationID := swag.JoinByFormat(valuesFilterOrganisationID, "")
	// query array param filter[organisation_id]
	if err := r.SetQueryParam("filter[organisation_id]", joinedFilterOrganisationID...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
