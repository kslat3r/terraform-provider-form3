// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/models"
)

// PostPaymentsIDReturnsReturnIDReversalsReader is a Reader for the PostPaymentsIDReturnsReturnIDReversals structure.
type PostPaymentsIDReturnsReturnIDReversalsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPaymentsIDReturnsReturnIDReversalsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostPaymentsIDReturnsReturnIDReversalsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostPaymentsIDReturnsReturnIDReversalsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPaymentsIDReturnsReturnIDReversalsCreated creates a PostPaymentsIDReturnsReturnIDReversalsCreated with default headers values
func NewPostPaymentsIDReturnsReturnIDReversalsCreated() *PostPaymentsIDReturnsReturnIDReversalsCreated {
	return &PostPaymentsIDReturnsReturnIDReversalsCreated{}
}

/*PostPaymentsIDReturnsReturnIDReversalsCreated handles this case with default header values.

Reversal creation response
*/
type PostPaymentsIDReturnsReturnIDReversalsCreated struct {
	Payload *models.ReturnReversalCreationResponse
}

func (o *PostPaymentsIDReturnsReturnIDReversalsCreated) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/returns/{returnId}/reversals][%d] postPaymentsIdReturnsReturnIdReversalsCreated  %+v", 201, o.Payload)
}

func (o *PostPaymentsIDReturnsReturnIDReversalsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReturnReversalCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReturnsReturnIDReversalsBadRequest creates a PostPaymentsIDReturnsReturnIDReversalsBadRequest with default headers values
func NewPostPaymentsIDReturnsReturnIDReversalsBadRequest() *PostPaymentsIDReturnsReturnIDReversalsBadRequest {
	return &PostPaymentsIDReturnsReturnIDReversalsBadRequest{}
}

/*PostPaymentsIDReturnsReturnIDReversalsBadRequest handles this case with default header values.

Reversal creation error
*/
type PostPaymentsIDReturnsReturnIDReversalsBadRequest struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReturnsReturnIDReversalsBadRequest) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/returns/{returnId}/reversals][%d] postPaymentsIdReturnsReturnIdReversalsBadRequest  %+v", 400, o.Payload)
}

func (o *PostPaymentsIDReturnsReturnIDReversalsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
