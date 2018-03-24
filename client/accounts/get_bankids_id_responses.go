// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ewilde/go-form3/models"
)

// GetBankidsIDReader is a Reader for the GetBankidsID structure.
type GetBankidsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBankidsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBankidsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBankidsIDOK creates a GetBankidsIDOK with default headers values
func NewGetBankidsIDOK() *GetBankidsIDOK {
	return &GetBankidsIDOK{}
}

/*GetBankidsIDOK handles this case with default header values.

BankId details
*/
type GetBankidsIDOK struct {
	Payload *models.BankIDDetailsResponse
}

func (o *GetBankidsIDOK) Error() string {
	return fmt.Sprintf("[GET /bankids/{id}][%d] getBankidsIdOK  %+v", 200, o.Payload)
}

func (o *GetBankidsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BankIDDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
