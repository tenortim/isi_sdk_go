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

// CreateAuthUserReader is a Reader for the CreateAuthUser structure.
type CreateAuthUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAuthUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateAuthUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateAuthUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateAuthUserOK creates a CreateAuthUserOK with default headers values
func NewCreateAuthUserOK() *CreateAuthUserOK {
	return &CreateAuthUserOK{}
}

/*CreateAuthUserOK handles this case with default header values.

Create a new user.
*/
type CreateAuthUserOK struct {
	Payload *models.CreateResponse
}

func (o *CreateAuthUserOK) Error() string {
	return fmt.Sprintf("[POST /platform/1/auth/users][%d] createAuthUserOK  %+v", 200, o.Payload)
}

func (o *CreateAuthUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAuthUserDefault creates a CreateAuthUserDefault with default headers values
func NewCreateAuthUserDefault(code int) *CreateAuthUserDefault {
	return &CreateAuthUserDefault{
		_statusCode: code,
	}
}

/*CreateAuthUserDefault handles this case with default header values.

Unexpected error
*/
type CreateAuthUserDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create auth user default response
func (o *CreateAuthUserDefault) Code() int {
	return o._statusCode
}

func (o *CreateAuthUserDefault) Error() string {
	return fmt.Sprintf("[POST /platform/1/auth/users][%d] createAuthUser default  %+v", o._statusCode, o.Payload)
}

func (o *CreateAuthUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}