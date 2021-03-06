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

// CreateSettingsKrb5DomainReader is a Reader for the CreateSettingsKrb5Domain structure.
type CreateSettingsKrb5DomainReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSettingsKrb5DomainReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateSettingsKrb5DomainOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateSettingsKrb5DomainDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateSettingsKrb5DomainOK creates a CreateSettingsKrb5DomainOK with default headers values
func NewCreateSettingsKrb5DomainOK() *CreateSettingsKrb5DomainOK {
	return &CreateSettingsKrb5DomainOK{}
}

/*CreateSettingsKrb5DomainOK handles this case with default header values.

Create a new krb5 domain.
*/
type CreateSettingsKrb5DomainOK struct {
	Payload *models.CreateResponse
}

func (o *CreateSettingsKrb5DomainOK) Error() string {
	return fmt.Sprintf("[POST /platform/1/auth/settings/krb5/domains][%d] createSettingsKrb5DomainOK  %+v", 200, o.Payload)
}

func (o *CreateSettingsKrb5DomainOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSettingsKrb5DomainDefault creates a CreateSettingsKrb5DomainDefault with default headers values
func NewCreateSettingsKrb5DomainDefault(code int) *CreateSettingsKrb5DomainDefault {
	return &CreateSettingsKrb5DomainDefault{
		_statusCode: code,
	}
}

/*CreateSettingsKrb5DomainDefault handles this case with default header values.

Unexpected error
*/
type CreateSettingsKrb5DomainDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create settings krb5 domain default response
func (o *CreateSettingsKrb5DomainDefault) Code() int {
	return o._statusCode
}

func (o *CreateSettingsKrb5DomainDefault) Error() string {
	return fmt.Sprintf("[POST /platform/1/auth/settings/krb5/domains][%d] createSettingsKrb5Domain default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSettingsKrb5DomainDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
