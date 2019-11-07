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

// DeleteSepareconciliationIDReader is a Reader for the DeleteSepareconciliationID structure.
type DeleteSepareconciliationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSepareconciliationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSepareconciliationIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteSepareconciliationIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteSepareconciliationIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteSepareconciliationIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewDeleteSepareconciliationIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteSepareconciliationIDNoContent creates a DeleteSepareconciliationIDNoContent with default headers values
func NewDeleteSepareconciliationIDNoContent() *DeleteSepareconciliationIDNoContent {
	return &DeleteSepareconciliationIDNoContent{}
}

/*DeleteSepareconciliationIDNoContent handles this case with default header values.

Association deleted
*/
type DeleteSepareconciliationIDNoContent struct {
}

func (o *DeleteSepareconciliationIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /separeconciliation/{id}][%d] deleteSepareconciliationIdNoContent ", 204)
}

func (o *DeleteSepareconciliationIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSepareconciliationIDBadRequest creates a DeleteSepareconciliationIDBadRequest with default headers values
func NewDeleteSepareconciliationIDBadRequest() *DeleteSepareconciliationIDBadRequest {
	return &DeleteSepareconciliationIDBadRequest{}
}

/*DeleteSepareconciliationIDBadRequest handles this case with default header values.

Bad Request
*/
type DeleteSepareconciliationIDBadRequest struct {
	Payload *models.APIError
}

func (o *DeleteSepareconciliationIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /separeconciliation/{id}][%d] deleteSepareconciliationIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteSepareconciliationIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSepareconciliationIDForbidden creates a DeleteSepareconciliationIDForbidden with default headers values
func NewDeleteSepareconciliationIDForbidden() *DeleteSepareconciliationIDForbidden {
	return &DeleteSepareconciliationIDForbidden{}
}

/*DeleteSepareconciliationIDForbidden handles this case with default header values.

Forbidden
*/
type DeleteSepareconciliationIDForbidden struct {
	Payload *models.APIError
}

func (o *DeleteSepareconciliationIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /separeconciliation/{id}][%d] deleteSepareconciliationIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteSepareconciliationIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSepareconciliationIDNotFound creates a DeleteSepareconciliationIDNotFound with default headers values
func NewDeleteSepareconciliationIDNotFound() *DeleteSepareconciliationIDNotFound {
	return &DeleteSepareconciliationIDNotFound{}
}

/*DeleteSepareconciliationIDNotFound handles this case with default header values.

Not Found
*/
type DeleteSepareconciliationIDNotFound struct {
	Payload *models.APIError
}

func (o *DeleteSepareconciliationIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /separeconciliation/{id}][%d] deleteSepareconciliationIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSepareconciliationIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSepareconciliationIDConflict creates a DeleteSepareconciliationIDConflict with default headers values
func NewDeleteSepareconciliationIDConflict() *DeleteSepareconciliationIDConflict {
	return &DeleteSepareconciliationIDConflict{}
}

/*DeleteSepareconciliationIDConflict handles this case with default header values.

Conflict
*/
type DeleteSepareconciliationIDConflict struct {
	Payload *models.APIError
}

func (o *DeleteSepareconciliationIDConflict) Error() string {
	return fmt.Sprintf("[DELETE /separeconciliation/{id}][%d] deleteSepareconciliationIdConflict  %+v", 409, o.Payload)
}

func (o *DeleteSepareconciliationIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
