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

// CreateProvidersLdapItemReader is a Reader for the CreateProvidersLdapItem structure.
type CreateProvidersLdapItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProvidersLdapItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateProvidersLdapItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateProvidersLdapItemDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateProvidersLdapItemOK creates a CreateProvidersLdapItemOK with default headers values
func NewCreateProvidersLdapItemOK() *CreateProvidersLdapItemOK {
	return &CreateProvidersLdapItemOK{}
}

/*CreateProvidersLdapItemOK handles this case with default header values.

Create a new LDAP provider.
*/
type CreateProvidersLdapItemOK struct {
	Payload *models.CreateResponse
}

func (o *CreateProvidersLdapItemOK) Error() string {
	return fmt.Sprintf("[POST /platform/3/auth/providers/ldap][%d] createProvidersLdapItemOK  %+v", 200, o.Payload)
}

func (o *CreateProvidersLdapItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProvidersLdapItemDefault creates a CreateProvidersLdapItemDefault with default headers values
func NewCreateProvidersLdapItemDefault(code int) *CreateProvidersLdapItemDefault {
	return &CreateProvidersLdapItemDefault{
		_statusCode: code,
	}
}

/*CreateProvidersLdapItemDefault handles this case with default header values.

Unexpected error
*/
type CreateProvidersLdapItemDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create providers ldap item default response
func (o *CreateProvidersLdapItemDefault) Code() int {
	return o._statusCode
}

func (o *CreateProvidersLdapItemDefault) Error() string {
	return fmt.Sprintf("[POST /platform/3/auth/providers/ldap][%d] createProvidersLdapItem default  %+v", o._statusCode, o.Payload)
}

func (o *CreateProvidersLdapItemDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}