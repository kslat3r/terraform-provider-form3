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

// NewGetSepaReconciliationParams creates a new GetSepaReconciliationParams object
// with the default values initialized.
func NewGetSepaReconciliationParams() *GetSepaReconciliationParams {
	var ()
	return &GetSepaReconciliationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSepaReconciliationParamsWithTimeout creates a new GetSepaReconciliationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSepaReconciliationParamsWithTimeout(timeout time.Duration) *GetSepaReconciliationParams {
	var ()
	return &GetSepaReconciliationParams{

		timeout: timeout,
	}
}

// NewGetSepaReconciliationParamsWithContext creates a new GetSepaReconciliationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSepaReconciliationParamsWithContext(ctx context.Context) *GetSepaReconciliationParams {
	var ()
	return &GetSepaReconciliationParams{

		Context: ctx,
	}
}

// NewGetSepaReconciliationParamsWithHTTPClient creates a new GetSepaReconciliationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSepaReconciliationParamsWithHTTPClient(client *http.Client) *GetSepaReconciliationParams {
	var ()
	return &GetSepaReconciliationParams{
		HTTPClient: client,
	}
}

/*GetSepaReconciliationParams contains all the parameters to send to the API endpoint
for the get sepa reconciliation operation typically these are written to a http.Request
*/
type GetSepaReconciliationParams struct {

	/*FilterOrganisationID
	  Organisation id

	*/
	FilterOrganisationID []strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get sepa reconciliation params
func (o *GetSepaReconciliationParams) WithTimeout(timeout time.Duration) *GetSepaReconciliationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sepa reconciliation params
func (o *GetSepaReconciliationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sepa reconciliation params
func (o *GetSepaReconciliationParams) WithContext(ctx context.Context) *GetSepaReconciliationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sepa reconciliation params
func (o *GetSepaReconciliationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sepa reconciliation params
func (o *GetSepaReconciliationParams) WithHTTPClient(client *http.Client) *GetSepaReconciliationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sepa reconciliation params
func (o *GetSepaReconciliationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilterOrganisationID adds the filterOrganisationID to the get sepa reconciliation params
func (o *GetSepaReconciliationParams) WithFilterOrganisationID(filterOrganisationID []strfmt.UUID) *GetSepaReconciliationParams {
	o.SetFilterOrganisationID(filterOrganisationID)
	return o
}

// SetFilterOrganisationID adds the filterOrganisationId to the get sepa reconciliation params
func (o *GetSepaReconciliationParams) SetFilterOrganisationID(filterOrganisationID []strfmt.UUID) {
	o.FilterOrganisationID = filterOrganisationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSepaReconciliationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
