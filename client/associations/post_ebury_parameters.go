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

	models "github.com/form3tech-oss/terraform-provider-form3/models"
)

// NewPostEburyParams creates a new PostEburyParams object
// with the default values initialized.
func NewPostEburyParams() *PostEburyParams {
	var ()
	return &PostEburyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostEburyParamsWithTimeout creates a new PostEburyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostEburyParamsWithTimeout(timeout time.Duration) *PostEburyParams {
	var ()
	return &PostEburyParams{

		timeout: timeout,
	}
}

// NewPostEburyParamsWithContext creates a new PostEburyParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostEburyParamsWithContext(ctx context.Context) *PostEburyParams {
	var ()
	return &PostEburyParams{

		Context: ctx,
	}
}

// NewPostEburyParamsWithHTTPClient creates a new PostEburyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostEburyParamsWithHTTPClient(client *http.Client) *PostEburyParams {
	var ()
	return &PostEburyParams{
		HTTPClient: client,
	}
}

/*PostEburyParams contains all the parameters to send to the API endpoint
for the post ebury operation typically these are written to a http.Request
*/
type PostEburyParams struct {

	/*CreationRequest*/
	CreationRequest *models.EburyAssociationCreation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post ebury params
func (o *PostEburyParams) WithTimeout(timeout time.Duration) *PostEburyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post ebury params
func (o *PostEburyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post ebury params
func (o *PostEburyParams) WithContext(ctx context.Context) *PostEburyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post ebury params
func (o *PostEburyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post ebury params
func (o *PostEburyParams) WithHTTPClient(client *http.Client) *PostEburyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post ebury params
func (o *PostEburyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreationRequest adds the creationRequest to the post ebury params
func (o *PostEburyParams) WithCreationRequest(creationRequest *models.EburyAssociationCreation) *PostEburyParams {
	o.SetCreationRequest(creationRequest)
	return o
}

// SetCreationRequest adds the creationRequest to the post ebury params
func (o *PostEburyParams) SetCreationRequest(creationRequest *models.EburyAssociationCreation) {
	o.CreationRequest = creationRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostEburyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreationRequest != nil {
		if err := r.SetBodyParam(o.CreationRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
