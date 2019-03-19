// Code generated by go-swagger; DO NOT EDIT.

package auth_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// DeleteUserMemberOfMemberOfReader is a Reader for the DeleteUserMemberOfMemberOf structure.
type DeleteUserMemberOfMemberOfReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserMemberOfMemberOfReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteUserMemberOfMemberOfNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteUserMemberOfMemberOfDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteUserMemberOfMemberOfNoContent creates a DeleteUserMemberOfMemberOfNoContent with default headers values
func NewDeleteUserMemberOfMemberOfNoContent() *DeleteUserMemberOfMemberOfNoContent {
	return &DeleteUserMemberOfMemberOfNoContent{}
}

/*DeleteUserMemberOfMemberOfNoContent handles this case with default header values.

Success.
*/
type DeleteUserMemberOfMemberOfNoContent struct {
}

func (o *DeleteUserMemberOfMemberOfNoContent) Error() string {
	return fmt.Sprintf("[DELETE /platform/3/auth/users/{User}/member-of/{UserMemberOfMemberOf}][%d] deleteUserMemberOfMemberOfNoContent ", 204)
}

func (o *DeleteUserMemberOfMemberOfNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserMemberOfMemberOfDefault creates a DeleteUserMemberOfMemberOfDefault with default headers values
func NewDeleteUserMemberOfMemberOfDefault(code int) *DeleteUserMemberOfMemberOfDefault {
	return &DeleteUserMemberOfMemberOfDefault{
		_statusCode: code,
	}
}

/*DeleteUserMemberOfMemberOfDefault handles this case with default header values.

Unexpected error
*/
type DeleteUserMemberOfMemberOfDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete user member of member of default response
func (o *DeleteUserMemberOfMemberOfDefault) Code() int {
	return o._statusCode
}

func (o *DeleteUserMemberOfMemberOfDefault) Error() string {
	return fmt.Sprintf("[DELETE /platform/3/auth/users/{User}/member-of/{UserMemberOfMemberOf}][%d] deleteUserMemberOfMemberOf default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteUserMemberOfMemberOfDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}