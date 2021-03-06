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

// NewGetVocalinkreportParams creates a new GetVocalinkreportParams object
// with the default values initialized.
func NewGetVocalinkreportParams() *GetVocalinkreportParams {
	var ()
	return &GetVocalinkreportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVocalinkreportParamsWithTimeout creates a new GetVocalinkreportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVocalinkreportParamsWithTimeout(timeout time.Duration) *GetVocalinkreportParams {
	var ()
	return &GetVocalinkreportParams{

		timeout: timeout,
	}
}

// NewGetVocalinkreportParamsWithContext creates a new GetVocalinkreportParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVocalinkreportParamsWithContext(ctx context.Context) *GetVocalinkreportParams {
	var ()
	return &GetVocalinkreportParams{

		Context: ctx,
	}
}

// NewGetVocalinkreportParamsWithHTTPClient creates a new GetVocalinkreportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVocalinkreportParamsWithHTTPClient(client *http.Client) *GetVocalinkreportParams {
	var ()
	return &GetVocalinkreportParams{
		HTTPClient: client,
	}
}

/*GetVocalinkreportParams contains all the parameters to send to the API endpoint
for the get vocalinkreport operation typically these are written to a http.Request
*/
type GetVocalinkreportParams struct {

	/*FilterOrganisationID
	  Organisation id

	*/
	FilterOrganisationID []strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get vocalinkreport params
func (o *GetVocalinkreportParams) WithTimeout(timeout time.Duration) *GetVocalinkreportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vocalinkreport params
func (o *GetVocalinkreportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vocalinkreport params
func (o *GetVocalinkreportParams) WithContext(ctx context.Context) *GetVocalinkreportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vocalinkreport params
func (o *GetVocalinkreportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vocalinkreport params
func (o *GetVocalinkreportParams) WithHTTPClient(client *http.Client) *GetVocalinkreportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vocalinkreport params
func (o *GetVocalinkreportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilterOrganisationID adds the filterOrganisationID to the get vocalinkreport params
func (o *GetVocalinkreportParams) WithFilterOrganisationID(filterOrganisationID []strfmt.UUID) *GetVocalinkreportParams {
	o.SetFilterOrganisationID(filterOrganisationID)
	return o
}

// SetFilterOrganisationID adds the filterOrganisationId to the get vocalinkreport params
func (o *GetVocalinkreportParams) SetFilterOrganisationID(filterOrganisationID []strfmt.UUID) {
	o.FilterOrganisationID = filterOrganisationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetVocalinkreportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
