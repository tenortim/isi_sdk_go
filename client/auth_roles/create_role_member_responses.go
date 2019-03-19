// Code generated by go-swagger; DO NOT EDIT.

package auth_roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// CreateRoleMemberReader is a Reader for the CreateRoleMember structure.
type CreateRoleMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRoleMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateRoleMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateRoleMemberDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateRoleMemberOK creates a CreateRoleMemberOK with default headers values
func NewCreateRoleMemberOK() *CreateRoleMemberOK {
	return &CreateRoleMemberOK{}
}

/*CreateRoleMemberOK handles this case with default header values.

Add a member to the role.
*/
type CreateRoleMemberOK struct {
	Payload *models.CreateResponse
}

func (o *CreateRoleMemberOK) Error() string {
	return fmt.Sprintf("[POST /platform/1/auth/roles/{Role}/members][%d] createRoleMemberOK  %+v", 200, o.Payload)
}

func (o *CreateRoleMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRoleMemberDefault creates a CreateRoleMemberDefault with default headers values
func NewCreateRoleMemberDefault(code int) *CreateRoleMemberDefault {
	return &CreateRoleMemberDefault{
		_statusCode: code,
	}
}

/*CreateRoleMemberDefault handles this case with default header values.

Unexpected error
*/
type CreateRoleMemberDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create role member default response
func (o *CreateRoleMemberDefault) Code() int {
	return o._statusCode
}

func (o *CreateRoleMemberDefault) Error() string {
	return fmt.Sprintf("[POST /platform/1/auth/roles/{Role}/members][%d] createRoleMember default  %+v", o._statusCode, o.Payload)
}

func (o *CreateRoleMemberDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}