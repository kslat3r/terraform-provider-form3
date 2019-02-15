// Code generated by go-swagger; DO NOT EDIT.

package payments

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

	models "github.com/form3tech-oss/go-form3/models"
)

// NewPatchPaymentsIDSubmissionsSubmissionIDParams creates a new PatchPaymentsIDSubmissionsSubmissionIDParams object
// with the default values initialized.
func NewPatchPaymentsIDSubmissionsSubmissionIDParams() *PatchPaymentsIDSubmissionsSubmissionIDParams {
	var ()
	return &PatchPaymentsIDSubmissionsSubmissionIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchPaymentsIDSubmissionsSubmissionIDParamsWithTimeout creates a new PatchPaymentsIDSubmissionsSubmissionIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchPaymentsIDSubmissionsSubmissionIDParamsWithTimeout(timeout time.Duration) *PatchPaymentsIDSubmissionsSubmissionIDParams {
	var ()
	return &PatchPaymentsIDSubmissionsSubmissionIDParams{

		timeout: timeout,
	}
}

// NewPatchPaymentsIDSubmissionsSubmissionIDParamsWithContext creates a new PatchPaymentsIDSubmissionsSubmissionIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchPaymentsIDSubmissionsSubmissionIDParamsWithContext(ctx context.Context) *PatchPaymentsIDSubmissionsSubmissionIDParams {
	var ()
	return &PatchPaymentsIDSubmissionsSubmissionIDParams{

		Context: ctx,
	}
}

// NewPatchPaymentsIDSubmissionsSubmissionIDParamsWithHTTPClient creates a new PatchPaymentsIDSubmissionsSubmissionIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchPaymentsIDSubmissionsSubmissionIDParamsWithHTTPClient(client *http.Client) *PatchPaymentsIDSubmissionsSubmissionIDParams {
	var ()
	return &PatchPaymentsIDSubmissionsSubmissionIDParams{
		HTTPClient: client,
	}
}

/*PatchPaymentsIDSubmissionsSubmissionIDParams contains all the parameters to send to the API endpoint
for the patch payments ID submissions submission ID operation typically these are written to a http.Request
*/
type PatchPaymentsIDSubmissionsSubmissionIDParams struct {

	/*SubmissionUpdateRequest*/
	SubmissionUpdateRequest *models.PaymentSubmissionAmendment
	/*ID
	  Payment Id

	*/
	ID strfmt.UUID
	/*SubmissionID
	  Submission Id

	*/
	SubmissionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch payments ID submissions submission ID params
func (o *PatchPaymentsIDSubmissionsSubmissionIDParams) WithTimeout(timeout time.Duration) *PatchPaymentsIDSubmissionsSubmissionIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch payments ID submissions submission ID params
func (o *PatchPaymentsIDSubmissionsSubmissionIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch payments ID submissions submission ID params
func (o *PatchPaymentsIDSubmissionsSubmissionIDParams) WithContext(ctx context.Context) *PatchPaymentsIDSubmissionsSubmissionIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch payments ID submissions submission ID params
func (o *PatchPaymentsIDSubmissionsSubmissionIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch payments ID submissions submission ID params
func (o *PatchPaymentsIDSubmissionsSubmissionIDParams) WithHTTPClient(client *http.Client) *PatchPaymentsIDSubmissionsSubmissionIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch payments ID submissions submission ID params
func (o *PatchPaymentsIDSubmissionsSubmissionIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSubmissionUpdateRequest adds the submissionUpdateRequest to the patch payments ID submissions submission ID params
func (o *PatchPaymentsIDSubmissionsSubmissionIDParams) WithSubmissionUpdateRequest(submissionUpdateRequest *models.PaymentSubmissionAmendment) *PatchPaymentsIDSubmissionsSubmissionIDParams {
	o.SetSubmissionUpdateRequest(submissionUpdateRequest)
	return o
}

// SetSubmissionUpdateRequest adds the submissionUpdateRequest to the patch payments ID submissions submission ID params
func (o *PatchPaymentsIDSubmissionsSubmissionIDParams) SetSubmissionUpdateRequest(submissionUpdateRequest *models.PaymentSubmissionAmendment) {
	o.SubmissionUpdateRequest = submissionUpdateRequest
}

// WithID adds the id to the patch payments ID submissions submission ID params
func (o *PatchPaymentsIDSubmissionsSubmissionIDParams) WithID(id strfmt.UUID) *PatchPaymentsIDSubmissionsSubmissionIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch payments ID submissions submission ID params
func (o *PatchPaymentsIDSubmissionsSubmissionIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithSubmissionID adds the submissionID to the patch payments ID submissions submission ID params
func (o *PatchPaymentsIDSubmissionsSubmissionIDParams) WithSubmissionID(submissionID strfmt.UUID) *PatchPaymentsIDSubmissionsSubmissionIDParams {
	o.SetSubmissionID(submissionID)
	return o
}

// SetSubmissionID adds the submissionId to the patch payments ID submissions submission ID params
func (o *PatchPaymentsIDSubmissionsSubmissionIDParams) SetSubmissionID(submissionID strfmt.UUID) {
	o.SubmissionID = submissionID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchPaymentsIDSubmissionsSubmissionIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SubmissionUpdateRequest != nil {
		if err := r.SetBodyParam(o.SubmissionUpdateRequest); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param submissionId
	if err := r.SetPathParam("submissionId", o.SubmissionID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
