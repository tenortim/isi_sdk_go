// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// DeleteAuthUserReader is a Reader for the DeleteAuthUser structure.
type DeleteAuthUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAuthUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteAuthUserNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteAuthUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAuthUserNoContent creates a DeleteAuthUserNoContent with default headers values
func NewDeleteAuthUserNoContent() *DeleteAuthUserNoContent {
	return &DeleteAuthUserNoContent{}
}

/*DeleteAuthUserNoContent handles this case with default header values.

Success.
*/
type DeleteAuthUserNoContent struct {
}

func (o *DeleteAuthUserNoContent) Error() string {
	return fmt.Sprintf("[DELETE /platform/1/auth/users/{AuthUserId}][%d] deleteAuthUserNoContent ", 204)
}

func (o *DeleteAuthUserNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAuthUserDefault creates a DeleteAuthUserDefault with default headers values
func NewDeleteAuthUserDefault(code int) *DeleteAuthUserDefault {
	return &DeleteAuthUserDefault{
		_statusCode: code,
	}
}

/*DeleteAuthUserDefault handles this case with default header values.

Unexpected error
*/
type DeleteAuthUserDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete auth user default response
func (o *DeleteAuthUserDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAuthUserDefault) Error() string {
	return fmt.Sprintf("[DELETE /platform/1/auth/users/{AuthUserId}][%d] deleteAuthUser default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteAuthUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}