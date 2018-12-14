// Code generated by go-swagger; DO NOT EDIT.

package associations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/models"
)

// GetPayportReader is a Reader for the GetPayport structure.
type GetPayportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPayportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPayportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPayportOK creates a GetPayportOK with default headers values
func NewGetPayportOK() *GetPayportOK {
	return &GetPayportOK{}
}

/*GetPayportOK handles this case with default header values.

List of associations
*/
type GetPayportOK struct {
	Payload *models.PayportAssociationDetailsListResponse
}

func (o *GetPayportOK) Error() string {
	return fmt.Sprintf("[GET /payport][%d] getPayportOK  %+v", 200, o.Payload)
}

func (o *GetPayportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PayportAssociationDetailsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
