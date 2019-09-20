// Code generated by go-swagger; DO NOT EDIT.

package associations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteProductsIDReader is a Reader for the DeleteProductsID structure.
type DeleteProductsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProductsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteProductsIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteProductsIDNoContent creates a DeleteProductsIDNoContent with default headers values
func NewDeleteProductsIDNoContent() *DeleteProductsIDNoContent {
	return &DeleteProductsIDNoContent{}
}

/*DeleteProductsIDNoContent handles this case with default header values.

Association deleted
*/
type DeleteProductsIDNoContent struct {
}

func (o *DeleteProductsIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /products/{id}][%d] deleteProductsIdNoContent ", 204)
}

func (o *DeleteProductsIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}