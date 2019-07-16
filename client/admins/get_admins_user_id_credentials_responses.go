// Code generated by go-swagger; DO NOT EDIT.

package admins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/terraform-provider-form3/models"
)

// GetAdminsUserIDCredentialsReader is a Reader for the GetAdminsUserIDCredentials structure.
type GetAdminsUserIDCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAdminsUserIDCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAdminsUserIDCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAdminsUserIDCredentialsOK creates a GetAdminsUserIDCredentialsOK with default headers values
func NewGetAdminsUserIDCredentialsOK() *GetAdminsUserIDCredentialsOK {
	return &GetAdminsUserIDCredentialsOK{}
}

/*GetAdminsUserIDCredentialsOK handles this case with default header values.

List of admin credentials for user
*/
type GetAdminsUserIDCredentialsOK struct {
	Payload *models.UserCredentialListResponse
}

func (o *GetAdminsUserIDCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /admins/{user_id}/credentials][%d] getAdminsUserIdCredentialsOK  %+v", 200, o.Payload)
}

func (o *GetAdminsUserIDCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserCredentialListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}