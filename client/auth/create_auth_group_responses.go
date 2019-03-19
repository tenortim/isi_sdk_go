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

// CreateAuthGroupReader is a Reader for the CreateAuthGroup structure.
type CreateAuthGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAuthGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateAuthGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateAuthGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateAuthGroupOK creates a CreateAuthGroupOK with default headers values
func NewCreateAuthGroupOK() *CreateAuthGroupOK {
	return &CreateAuthGroupOK{}
}

/*CreateAuthGroupOK handles this case with default header values.

Create a new group.
*/
type CreateAuthGroupOK struct {
	Payload *models.CreateResponse
}

func (o *CreateAuthGroupOK) Error() string {
	return fmt.Sprintf("[POST /platform/1/auth/groups][%d] createAuthGroupOK  %+v", 200, o.Payload)
}

func (o *CreateAuthGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAuthGroupDefault creates a CreateAuthGroupDefault with default headers values
func NewCreateAuthGroupDefault(code int) *CreateAuthGroupDefault {
	return &CreateAuthGroupDefault{
		_statusCode: code,
	}
}

/*CreateAuthGroupDefault handles this case with default header values.

Unexpected error
*/
type CreateAuthGroupDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create auth group default response
func (o *CreateAuthGroupDefault) Code() int {
	return o._statusCode
}

func (o *CreateAuthGroupDefault) Error() string {
	return fmt.Sprintf("[POST /platform/1/auth/groups][%d] createAuthGroup default  %+v", o._statusCode, o.Payload)
}

func (o *CreateAuthGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
