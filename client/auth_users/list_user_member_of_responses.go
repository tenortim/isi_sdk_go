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

// ListUserMemberOfReader is a Reader for the ListUserMemberOf structure.
type ListUserMemberOfReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListUserMemberOfReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListUserMemberOfOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListUserMemberOfDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListUserMemberOfOK creates a ListUserMemberOfOK with default headers values
func NewListUserMemberOfOK() *ListUserMemberOfOK {
	return &ListUserMemberOfOK{}
}

/*ListUserMemberOfOK handles this case with default header values.

List all groups the user is a member of.
*/
type ListUserMemberOfOK struct {
	Payload *models.UserMemberOf
}

func (o *ListUserMemberOfOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/auth/users/{User}/member-of][%d] listUserMemberOfOK  %+v", 200, o.Payload)
}

func (o *ListUserMemberOfOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserMemberOf)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUserMemberOfDefault creates a ListUserMemberOfDefault with default headers values
func NewListUserMemberOfDefault(code int) *ListUserMemberOfDefault {
	return &ListUserMemberOfDefault{
		_statusCode: code,
	}
}

/*ListUserMemberOfDefault handles this case with default header values.

Unexpected error
*/
type ListUserMemberOfDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list user member of default response
func (o *ListUserMemberOfDefault) Code() int {
	return o._statusCode
}

func (o *ListUserMemberOfDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/auth/users/{User}/member-of][%d] listUserMemberOf default  %+v", o._statusCode, o.Payload)
}

func (o *ListUserMemberOfDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
