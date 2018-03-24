// Code generated by go-swagger; DO NOT EDIT.

package users

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

	models "github.com/ewilde/go-form3/models"
)

// NewPatchUsersUserIDParams creates a new PatchUsersUserIDParams object
// with the default values initialized.
func NewPatchUsersUserIDParams() *PatchUsersUserIDParams {
	var ()
	return &PatchUsersUserIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchUsersUserIDParamsWithTimeout creates a new PatchUsersUserIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchUsersUserIDParamsWithTimeout(timeout time.Duration) *PatchUsersUserIDParams {
	var ()
	return &PatchUsersUserIDParams{

		timeout: timeout,
	}
}

// NewPatchUsersUserIDParamsWithContext creates a new PatchUsersUserIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchUsersUserIDParamsWithContext(ctx context.Context) *PatchUsersUserIDParams {
	var ()
	return &PatchUsersUserIDParams{

		Context: ctx,
	}
}

// NewPatchUsersUserIDParamsWithHTTPClient creates a new PatchUsersUserIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchUsersUserIDParamsWithHTTPClient(client *http.Client) *PatchUsersUserIDParams {
	var ()
	return &PatchUsersUserIDParams{
		HTTPClient: client,
	}
}

/*PatchUsersUserIDParams contains all the parameters to send to the API endpoint
for the patch users user ID operation typically these are written to a http.Request
*/
type PatchUsersUserIDParams struct {

	/*UserUpdateRequest*/
	UserUpdateRequest *models.UserCreation
	/*UserID
	  User Id

	*/
	UserID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch users user ID params
func (o *PatchUsersUserIDParams) WithTimeout(timeout time.Duration) *PatchUsersUserIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch users user ID params
func (o *PatchUsersUserIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch users user ID params
func (o *PatchUsersUserIDParams) WithContext(ctx context.Context) *PatchUsersUserIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch users user ID params
func (o *PatchUsersUserIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch users user ID params
func (o *PatchUsersUserIDParams) WithHTTPClient(client *http.Client) *PatchUsersUserIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch users user ID params
func (o *PatchUsersUserIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserUpdateRequest adds the userUpdateRequest to the patch users user ID params
func (o *PatchUsersUserIDParams) WithUserUpdateRequest(userUpdateRequest *models.UserCreation) *PatchUsersUserIDParams {
	o.SetUserUpdateRequest(userUpdateRequest)
	return o
}

// SetUserUpdateRequest adds the userUpdateRequest to the patch users user ID params
func (o *PatchUsersUserIDParams) SetUserUpdateRequest(userUpdateRequest *models.UserCreation) {
	o.UserUpdateRequest = userUpdateRequest
}

// WithUserID adds the userID to the patch users user ID params
func (o *PatchUsersUserIDParams) WithUserID(userID strfmt.UUID) *PatchUsersUserIDParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the patch users user ID params
func (o *PatchUsersUserIDParams) SetUserID(userID strfmt.UUID) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchUsersUserIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.UserUpdateRequest != nil {
		if err := r.SetBodyParam(o.UserUpdateRequest); err != nil {
			return err
		}
	}

	// path param user_id
	if err := r.SetPathParam("user_id", o.UserID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
