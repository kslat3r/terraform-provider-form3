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

// PostSepasctReader is a Reader for the PostSepasct structure.
type PostSepasctReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSepasctReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostSepasctCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostSepasctCreated creates a PostSepasctCreated with default headers values
func NewPostSepasctCreated() *PostSepasctCreated {
	return &PostSepasctCreated{}
}

/*PostSepasctCreated handles this case with default header values.

creation response
*/
type PostSepasctCreated struct {
	Payload *models.SepaSctAssociationCreationResponse
}

func (o *PostSepasctCreated) Error() string {
	return fmt.Sprintf("[POST /sepasct][%d] postSepasctCreated  %+v", 201, o.Payload)
}

func (o *PostSepasctCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SepaSctAssociationCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
