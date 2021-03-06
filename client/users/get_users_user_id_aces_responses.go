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

// GetUsersUserIDAcesReader is a Reader for the GetUsersUserIDAces structure.
type GetUsersUserIDAcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersUserIDAcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUsersUserIDAcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUsersUserIDAcesOK creates a GetUsersUserIDAcesOK with default headers values
func NewGetUsersUserIDAcesOK() *GetUsersUserIDAcesOK {
	return &GetUsersUserIDAcesOK{}
}

/*GetUsersUserIDAcesOK handles this case with default header values.

List of access control entries for this user
*/
type GetUsersUserIDAcesOK struct {
	Payload *models.AceDetailsListResponse
}

func (o *GetUsersUserIDAcesOK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/aces][%d] getUsersUserIdAcesOK  %+v", 200, o.Payload)
}

func (o *GetUsersUserIDAcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AceDetailsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
