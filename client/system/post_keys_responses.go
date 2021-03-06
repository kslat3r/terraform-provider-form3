// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/terraform-provider-form3/models"
)

// PostKeysReader is a Reader for the PostKeys structure.
type PostKeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostKeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostKeysCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostKeysCreated creates a PostKeysCreated with default headers values
func NewPostKeysCreated() *PostKeysCreated {
	return &PostKeysCreated{}
}

/*PostKeysCreated handles this case with default header values.

creation response
*/
type PostKeysCreated struct {
	Payload *models.KeyCreationResponse
}

func (o *PostKeysCreated) Error() string {
	return fmt.Sprintf("[POST /keys][%d] postKeysCreated  %+v", 201, o.Payload)
}

func (o *PostKeysCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KeyCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
