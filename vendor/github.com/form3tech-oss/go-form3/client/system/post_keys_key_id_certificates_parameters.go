// Code generated by go-swagger; DO NOT EDIT.

package system

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

// NewPostKeysKeyIDCertificatesParams creates a new PostKeysKeyIDCertificatesParams object
// with the default values initialized.
func NewPostKeysKeyIDCertificatesParams() *PostKeysKeyIDCertificatesParams {
	var ()
	return &PostKeysKeyIDCertificatesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostKeysKeyIDCertificatesParamsWithTimeout creates a new PostKeysKeyIDCertificatesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostKeysKeyIDCertificatesParamsWithTimeout(timeout time.Duration) *PostKeysKeyIDCertificatesParams {
	var ()
	return &PostKeysKeyIDCertificatesParams{

		timeout: timeout,
	}
}

// NewPostKeysKeyIDCertificatesParamsWithContext creates a new PostKeysKeyIDCertificatesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostKeysKeyIDCertificatesParamsWithContext(ctx context.Context) *PostKeysKeyIDCertificatesParams {
	var ()
	return &PostKeysKeyIDCertificatesParams{

		Context: ctx,
	}
}

// NewPostKeysKeyIDCertificatesParamsWithHTTPClient creates a new PostKeysKeyIDCertificatesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostKeysKeyIDCertificatesParamsWithHTTPClient(client *http.Client) *PostKeysKeyIDCertificatesParams {
	var ()
	return &PostKeysKeyIDCertificatesParams{
		HTTPClient: client,
	}
}

/*PostKeysKeyIDCertificatesParams contains all the parameters to send to the API endpoint
for the post keys key ID certificates operation typically these are written to a http.Request
*/
type PostKeysKeyIDCertificatesParams struct {

	/*CertificateCreationRequest*/
	CertificateCreationRequest *models.CertificateCreation
	/*KeyID
	  Key Id

	*/
	KeyID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post keys key ID certificates params
func (o *PostKeysKeyIDCertificatesParams) WithTimeout(timeout time.Duration) *PostKeysKeyIDCertificatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post keys key ID certificates params
func (o *PostKeysKeyIDCertificatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post keys key ID certificates params
func (o *PostKeysKeyIDCertificatesParams) WithContext(ctx context.Context) *PostKeysKeyIDCertificatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post keys key ID certificates params
func (o *PostKeysKeyIDCertificatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post keys key ID certificates params
func (o *PostKeysKeyIDCertificatesParams) WithHTTPClient(client *http.Client) *PostKeysKeyIDCertificatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post keys key ID certificates params
func (o *PostKeysKeyIDCertificatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCertificateCreationRequest adds the certificateCreationRequest to the post keys key ID certificates params
func (o *PostKeysKeyIDCertificatesParams) WithCertificateCreationRequest(certificateCreationRequest *models.CertificateCreation) *PostKeysKeyIDCertificatesParams {
	o.SetCertificateCreationRequest(certificateCreationRequest)
	return o
}

// SetCertificateCreationRequest adds the certificateCreationRequest to the post keys key ID certificates params
func (o *PostKeysKeyIDCertificatesParams) SetCertificateCreationRequest(certificateCreationRequest *models.CertificateCreation) {
	o.CertificateCreationRequest = certificateCreationRequest
}

// WithKeyID adds the keyID to the post keys key ID certificates params
func (o *PostKeysKeyIDCertificatesParams) WithKeyID(keyID strfmt.UUID) *PostKeysKeyIDCertificatesParams {
	o.SetKeyID(keyID)
	return o
}

// SetKeyID adds the keyId to the post keys key ID certificates params
func (o *PostKeysKeyIDCertificatesParams) SetKeyID(keyID strfmt.UUID) {
	o.KeyID = keyID
}

// WriteToRequest writes these params to a swagger request
func (o *PostKeysKeyIDCertificatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CertificateCreationRequest != nil {
		if err := r.SetBodyParam(o.CertificateCreationRequest); err != nil {
			return err
		}
	}

	// path param key_id
	if err := r.SetPathParam("key_id", o.KeyID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}