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

// ListAuthRolesReader is a Reader for the ListAuthRoles structure.
type ListAuthRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAuthRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListAuthRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListAuthRolesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAuthRolesOK creates a ListAuthRolesOK with default headers values
func NewListAuthRolesOK() *ListAuthRolesOK {
	return &ListAuthRolesOK{}
}

/*ListAuthRolesOK handles this case with default header values.

List all roles.
*/
type ListAuthRolesOK struct {
	Payload *models.AuthRolesExtended
}

func (o *ListAuthRolesOK) Error() string {
	return fmt.Sprintf("[GET /platform/1/auth/roles][%d] listAuthRolesOK  %+v", 200, o.Payload)
}

func (o *ListAuthRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthRolesExtended)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAuthRolesDefault creates a ListAuthRolesDefault with default headers values
func NewListAuthRolesDefault(code int) *ListAuthRolesDefault {
	return &ListAuthRolesDefault{
		_statusCode: code,
	}
}

/*ListAuthRolesDefault handles this case with default header values.

Unexpected error
*/
type ListAuthRolesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list auth roles default response
func (o *ListAuthRolesDefault) Code() int {
	return o._statusCode
}

func (o *ListAuthRolesDefault) Error() string {
	return fmt.Sprintf("[GET /platform/1/auth/roles][%d] listAuthRoles default  %+v", o._statusCode, o.Payload)
}

func (o *ListAuthRolesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
