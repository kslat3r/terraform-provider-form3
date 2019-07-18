// Code generated by go-swagger; DO NOT EDIT.

package admins

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
)

// NewPostAdminsUserIDCredentialsParams creates a new PostAdminsUserIDCredentialsParams object
// with the default values initialized.
func NewPostAdminsUserIDCredentialsParams() *PostAdminsUserIDCredentialsParams {
	var ()
	return &PostAdminsUserIDCredentialsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAdminsUserIDCredentialsParamsWithTimeout creates a new PostAdminsUserIDCredentialsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAdminsUserIDCredentialsParamsWithTimeout(timeout time.Duration) *PostAdminsUserIDCredentialsParams {
	var ()
	return &PostAdminsUserIDCredentialsParams{

		timeout: timeout,
	}
}

// NewPostAdminsUserIDCredentialsParamsWithContext creates a new PostAdminsUserIDCredentialsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAdminsUserIDCredentialsParamsWithContext(ctx context.Context) *PostAdminsUserIDCredentialsParams {
	var ()
	return &PostAdminsUserIDCredentialsParams{

		Context: ctx,
	}
}

// NewPostAdminsUserIDCredentialsParamsWithHTTPClient creates a new PostAdminsUserIDCredentialsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAdminsUserIDCredentialsParamsWithHTTPClient(client *http.Client) *PostAdminsUserIDCredentialsParams {
	var ()
	return &PostAdminsUserIDCredentialsParams{
		HTTPClient: client,
	}
}

/*PostAdminsUserIDCredentialsParams contains all the parameters to send to the API endpoint
for the post admins user ID credentials operation typically these are written to a http.Request
*/
type PostAdminsUserIDCredentialsParams struct {

	/*UserID
	  User Id

	*/
	UserID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post admins user ID credentials params
func (o *PostAdminsUserIDCredentialsParams) WithTimeout(timeout time.Duration) *PostAdminsUserIDCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post admins user ID credentials params
func (o *PostAdminsUserIDCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post admins user ID credentials params
func (o *PostAdminsUserIDCredentialsParams) WithContext(ctx context.Context) *PostAdminsUserIDCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post admins user ID credentials params
func (o *PostAdminsUserIDCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post admins user ID credentials params
func (o *PostAdminsUserIDCredentialsParams) WithHTTPClient(client *http.Client) *PostAdminsUserIDCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post admins user ID credentials params
func (o *PostAdminsUserIDCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the post admins user ID credentials params
func (o *PostAdminsUserIDCredentialsParams) WithUserID(userID strfmt.UUID) *PostAdminsUserIDCredentialsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the post admins user ID credentials params
func (o *PostAdminsUserIDCredentialsParams) SetUserID(userID strfmt.UUID) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PostAdminsUserIDCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param user_id
	if err := r.SetPathParam("user_id", o.UserID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
