// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/models"
)

// GetKeysKeyIDReader is a Reader for the GetKeysKeyID structure.
type GetKeysKeyIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKeysKeyIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetKeysKeyIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetKeysKeyIDOK creates a GetKeysKeyIDOK with default headers values
func NewGetKeysKeyIDOK() *GetKeysKeyIDOK {
	return &GetKeysKeyIDOK{}
}

/*GetKeysKeyIDOK handles this case with default header values.

Key details
*/
type GetKeysKeyIDOK struct {
	Payload *models.KeyDetailsResponse
}

func (o *GetKeysKeyIDOK) Error() string {
	return fmt.Sprintf("[GET /keys/{key_id}][%d] getKeysKeyIdOK  %+v", 200, o.Payload)
}

func (o *GetKeysKeyIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KeyDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
