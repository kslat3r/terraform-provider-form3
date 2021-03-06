// Code generated by go-swagger; DO NOT EDIT.

package accounts

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

// NewGetBicsParams creates a new GetBicsParams object
// with the default values initialized.
func NewGetBicsParams() *GetBicsParams {
	var ()
	return &GetBicsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBicsParamsWithTimeout creates a new GetBicsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBicsParamsWithTimeout(timeout time.Duration) *GetBicsParams {
	var ()
	return &GetBicsParams{

		timeout: timeout,
	}
}

// NewGetBicsParamsWithContext creates a new GetBicsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBicsParamsWithContext(ctx context.Context) *GetBicsParams {
	var ()
	return &GetBicsParams{

		Context: ctx,
	}
}

// NewGetBicsParamsWithHTTPClient creates a new GetBicsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBicsParamsWithHTTPClient(client *http.Client) *GetBicsParams {
	var ()
	return &GetBicsParams{
		HTTPClient: client,
	}
}

/*GetBicsParams contains all the parameters to send to the API endpoint
for the get bics operation typically these are written to a http.Request
*/
type GetBicsParams struct {

	/*FilterBic
	  Filter by specific bic

	*/
	FilterBic *string
	/*PageNumber
	  Which page to select

	*/
	PageNumber *string
	/*PageSize
	  Number of items to select

	*/
	PageSize *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get bics params
func (o *GetBicsParams) WithTimeout(timeout time.Duration) *GetBicsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bics params
func (o *GetBicsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bics params
func (o *GetBicsParams) WithContext(ctx context.Context) *GetBicsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bics params
func (o *GetBicsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bics params
func (o *GetBicsParams) WithHTTPClient(client *http.Client) *GetBicsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bics params
func (o *GetBicsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilterBic adds the filterBic to the get bics params
func (o *GetBicsParams) WithFilterBic(filterBic *string) *GetBicsParams {
	o.SetFilterBic(filterBic)
	return o
}

// SetFilterBic adds the filterBic to the get bics params
func (o *GetBicsParams) SetFilterBic(filterBic *string) {
	o.FilterBic = filterBic
}

// WithPageNumber adds the pageNumber to the get bics params
func (o *GetBicsParams) WithPageNumber(pageNumber *string) *GetBicsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get bics params
func (o *GetBicsParams) SetPageNumber(pageNumber *string) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get bics params
func (o *GetBicsParams) WithPageSize(pageSize *int64) *GetBicsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get bics params
func (o *GetBicsParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetBicsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FilterBic != nil {

		// query param filter[bic]
		var qrFilterBic string
		if o.FilterBic != nil {
			qrFilterBic = *o.FilterBic
		}
		qFilterBic := qrFilterBic
		if qFilterBic != "" {
			if err := r.SetQueryParam("filter[bic]", qFilterBic); err != nil {
				return err
			}
		}

	}

	if o.PageNumber != nil {

		// query param page[number]
		var qrPageNumber string
		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := qrPageNumber
		if qPageNumber != "" {
			if err := r.SetQueryParam("page[number]", qPageNumber); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param page[size]
		var qrPageSize int64
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("page[size]", qPageSize); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
