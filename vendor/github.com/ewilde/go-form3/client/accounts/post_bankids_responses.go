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

// PostBankidsReader is a Reader for the PostBankids structure.
type PostBankidsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostBankidsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostBankidsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostBankidsCreated creates a PostBankidsCreated with default headers values
func NewPostBankidsCreated() *PostBankidsCreated {
	return &PostBankidsCreated{}
}

/*PostBankidsCreated handles this case with default header values.

BankId creation response
*/
type PostBankidsCreated struct {
	Payload *models.BankIDCreationResponse
}

func (o *PostBankidsCreated) Error() string {
	return fmt.Sprintf("[POST /bankids][%d] postBankidsCreated  %+v", 201, o.Payload)
}

func (o *PostBankidsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BankIDCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
