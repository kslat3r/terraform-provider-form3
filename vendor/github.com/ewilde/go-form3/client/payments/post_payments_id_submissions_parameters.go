// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/ewilde/go-form3/models"
)

// NewPostPaymentsIDSubmissionsParams creates a new PostPaymentsIDSubmissionsParams object
// with the default values initialized.
func NewPostPaymentsIDSubmissionsParams() *PostPaymentsIDSubmissionsParams {
	var ()
	return &PostPaymentsIDSubmissionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPaymentsIDSubmissionsParamsWithTimeout creates a new PostPaymentsIDSubmissionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPaymentsIDSubmissionsParamsWithTimeout(timeout time.Duration) *PostPaymentsIDSubmissionsParams {
	var ()
	return &PostPaymentsIDSubmissionsParams{

		timeout: timeout,
	}
}

// NewPostPaymentsIDSubmissionsParamsWithContext creates a new PostPaymentsIDSubmissionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPaymentsIDSubmissionsParamsWithContext(ctx context.Context) *PostPaymentsIDSubmissionsParams {
	var ()
	return &PostPaymentsIDSubmissionsParams{

		Context: ctx,
	}
}

// NewPostPaymentsIDSubmissionsParamsWithHTTPClient creates a new PostPaymentsIDSubmissionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPaymentsIDSubmissionsParamsWithHTTPClient(client *http.Client) *PostPaymentsIDSubmissionsParams {
	var ()
	return &PostPaymentsIDSubmissionsParams{
		HTTPClient: client,
	}
}

/*PostPaymentsIDSubmissionsParams contains all the parameters to send to the API endpoint
for the post payments ID submissions operation typically these are written to a http.Request
*/
type PostPaymentsIDSubmissionsParams struct {

	/*SubmissionCreationRequest*/
	SubmissionCreationRequest *models.PaymentSubmissionCreation
	/*ID
	  Payment Id

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post payments ID submissions params
func (o *PostPaymentsIDSubmissionsParams) WithTimeout(timeout time.Duration) *PostPaymentsIDSubmissionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post payments ID submissions params
func (o *PostPaymentsIDSubmissionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post payments ID submissions params
func (o *PostPaymentsIDSubmissionsParams) WithContext(ctx context.Context) *PostPaymentsIDSubmissionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post payments ID submissions params
func (o *PostPaymentsIDSubmissionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post payments ID submissions params
func (o *PostPaymentsIDSubmissionsParams) WithHTTPClient(client *http.Client) *PostPaymentsIDSubmissionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post payments ID submissions params
func (o *PostPaymentsIDSubmissionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSubmissionCreationRequest adds the submissionCreationRequest to the post payments ID submissions params
func (o *PostPaymentsIDSubmissionsParams) WithSubmissionCreationRequest(submissionCreationRequest *models.PaymentSubmissionCreation) *PostPaymentsIDSubmissionsParams {
	o.SetSubmissionCreationRequest(submissionCreationRequest)
	return o
}

// SetSubmissionCreationRequest adds the submissionCreationRequest to the post payments ID submissions params
func (o *PostPaymentsIDSubmissionsParams) SetSubmissionCreationRequest(submissionCreationRequest *models.PaymentSubmissionCreation) {
	o.SubmissionCreationRequest = submissionCreationRequest
}

// WithID adds the id to the post payments ID submissions params
func (o *PostPaymentsIDSubmissionsParams) WithID(id strfmt.UUID) *PostPaymentsIDSubmissionsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post payments ID submissions params
func (o *PostPaymentsIDSubmissionsParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostPaymentsIDSubmissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SubmissionCreationRequest != nil {
		if err := r.SetBodyParam(o.SubmissionCreationRequest); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
