// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/terraform-provider-form3/models"
)

// GetUsersUserIDCredentialsPublicKeyReader is a Reader for the GetUsersUserIDCredentialsPublicKey structure.
type GetUsersUserIDCredentialsPublicKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersUserIDCredentialsPublicKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUsersUserIDCredentialsPublicKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUsersUserIDCredentialsPublicKeyOK creates a GetUsersUserIDCredentialsPublicKeyOK with default headers values
func NewGetUsersUserIDCredentialsPublicKeyOK() *GetUsersUserIDCredentialsPublicKeyOK {
	return &GetUsersUserIDCredentialsPublicKeyOK{}
}

/*GetUsersUserIDCredentialsPublicKeyOK handles this case with default header values.

List of public keys for user
*/
type GetUsersUserIDCredentialsPublicKeyOK struct {
	Payload *models.UserPublicKeyListResponse
}

func (o *GetUsersUserIDCredentialsPublicKeyOK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials/public_key][%d] getUsersUserIdCredentialsPublicKeyOK  %+v", 200, o.Payload)
}

func (o *GetUsersUserIDCredentialsPublicKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserPublicKeyListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
