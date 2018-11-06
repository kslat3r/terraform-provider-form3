// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/models"
)

// GetBicsIDReader is a Reader for the GetBicsID structure.
type GetBicsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBicsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBicsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBicsIDOK creates a GetBicsIDOK with default headers values
func NewGetBicsIDOK() *GetBicsIDOK {
	return &GetBicsIDOK{}
}

/*GetBicsIDOK handles this case with default header values.

Bic details
*/
type GetBicsIDOK struct {
	Payload *models.BicDetailsResponse
}

func (o *GetBicsIDOK) Error() string {
	return fmt.Sprintf("[GET /bics/{id}][%d] getBicsIdOK  %+v", 200, o.Payload)
}

func (o *GetBicsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BicDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
