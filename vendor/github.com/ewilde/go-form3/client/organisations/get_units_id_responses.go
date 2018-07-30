// Code generated by go-swagger; DO NOT EDIT.

package organisations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ewilde/go-form3/models"
)

// GetUnitsIDReader is a Reader for the GetUnitsID structure.
type GetUnitsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUnitsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUnitsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUnitsIDOK creates a GetUnitsIDOK with default headers values
func NewGetUnitsIDOK() *GetUnitsIDOK {
	return &GetUnitsIDOK{}
}

/*GetUnitsIDOK handles this case with default header values.

Organisation details
*/
type GetUnitsIDOK struct {
	Payload *models.OrganisationDetailsResponse
}

func (o *GetUnitsIDOK) Error() string {
	return fmt.Sprintf("[GET /units/{id}][%d] getUnitsIdOK  %+v", 200, o.Payload)
}

func (o *GetUnitsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrganisationDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
