// Code generated by go-swagger; DO NOT EDIT.

package protocols

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// DeleteNdmpUserReader is a Reader for the DeleteNdmpUser structure.
type DeleteNdmpUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNdmpUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteNdmpUserNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteNdmpUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteNdmpUserNoContent creates a DeleteNdmpUserNoContent with default headers values
func NewDeleteNdmpUserNoContent() *DeleteNdmpUserNoContent {
	return &DeleteNdmpUserNoContent{}
}

/*DeleteNdmpUserNoContent handles this case with default header values.

Success.
*/
type DeleteNdmpUserNoContent struct {
}

func (o *DeleteNdmpUserNoContent) Error() string {
	return fmt.Sprintf("[DELETE /platform/3/protocols/ndmp/users/{NdmpUserId}][%d] deleteNdmpUserNoContent ", 204)
}

func (o *DeleteNdmpUserNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNdmpUserDefault creates a DeleteNdmpUserDefault with default headers values
func NewDeleteNdmpUserDefault(code int) *DeleteNdmpUserDefault {
	return &DeleteNdmpUserDefault{
		_statusCode: code,
	}
}

/*DeleteNdmpUserDefault handles this case with default header values.

Unexpected error
*/
type DeleteNdmpUserDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete ndmp user default response
func (o *DeleteNdmpUserDefault) Code() int {
	return o._statusCode
}

func (o *DeleteNdmpUserDefault) Error() string {
	return fmt.Sprintf("[DELETE /platform/3/protocols/ndmp/users/{NdmpUserId}][%d] deleteNdmpUser default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteNdmpUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}