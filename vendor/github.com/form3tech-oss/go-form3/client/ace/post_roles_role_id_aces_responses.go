// Code generated by go-swagger; DO NOT EDIT.

package ace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/models"
)

// PostRolesRoleIDAcesReader is a Reader for the PostRolesRoleIDAces structure.
type PostRolesRoleIDAcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRolesRoleIDAcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostRolesRoleIDAcesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRolesRoleIDAcesCreated creates a PostRolesRoleIDAcesCreated with default headers values
func NewPostRolesRoleIDAcesCreated() *PostRolesRoleIDAcesCreated {
	return &PostRolesRoleIDAcesCreated{}
}

/*PostRolesRoleIDAcesCreated handles this case with default header values.

Ace creation response
*/
type PostRolesRoleIDAcesCreated struct {
	Payload *models.AceCreationResponse
}

func (o *PostRolesRoleIDAcesCreated) Error() string {
	return fmt.Sprintf("[POST /roles/{role_id}/aces][%d] postRolesRoleIdAcesCreated  %+v", 201, o.Payload)
}

func (o *PostRolesRoleIDAcesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AceCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}