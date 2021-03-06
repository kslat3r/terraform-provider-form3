// Code generated by go-swagger; DO NOT EDIT.

package associations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/terraform-provider-form3/models"
)

// GetSepareconciliationReader is a Reader for the GetSepareconciliation structure.
type GetSepareconciliationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSepareconciliationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSepareconciliationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSepareconciliationOK creates a GetSepareconciliationOK with default headers values
func NewGetSepareconciliationOK() *GetSepareconciliationOK {
	return &GetSepareconciliationOK{}
}

/*GetSepareconciliationOK handles this case with default header values.

List of associations
*/
type GetSepareconciliationOK struct {
	Payload *models.SepaReconciliationAssociationDetailsListResponse
}

func (o *GetSepareconciliationOK) Error() string {
	return fmt.Sprintf("[GET /separeconciliation][%d] getSepareconciliationOK  %+v", 200, o.Payload)
}

func (o *GetSepareconciliationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SepaReconciliationAssociationDetailsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
