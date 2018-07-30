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

// GetUnitsReader is a Reader for the GetUnits structure.
type GetUnitsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUnitsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUnitsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUnitsOK creates a GetUnitsOK with default headers values
func NewGetUnitsOK() *GetUnitsOK {
	return &GetUnitsOK{}
}

/*GetUnitsOK handles this case with default header values.

List of organisation details
*/
type GetUnitsOK struct {
	Payload *models.OrganisationDetailsListResponse
}

func (o *GetUnitsOK) Error() string {
	return fmt.Sprintf("[GET /units][%d] getUnitsOK  %+v", 200, o.Payload)
}

func (o *GetUnitsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrganisationDetailsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
