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

// GetProvidersLdapByIDReader is a Reader for the GetProvidersLdapByID structure.
type GetProvidersLdapByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProvidersLdapByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetProvidersLdapByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetProvidersLdapByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProvidersLdapByIDOK creates a GetProvidersLdapByIDOK with default headers values
func NewGetProvidersLdapByIDOK() *GetProvidersLdapByIDOK {
	return &GetProvidersLdapByIDOK{}
}

/*GetProvidersLdapByIDOK handles this case with default header values.

Retrieve the LDAP provider.
*/
type GetProvidersLdapByIDOK struct {
	Payload *models.ProvidersLdap
}

func (o *GetProvidersLdapByIDOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/auth/providers/ldap/{ProvidersLdapId}][%d] getProvidersLdapByIdOK  %+v", 200, o.Payload)
}

func (o *GetProvidersLdapByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProvidersLdap)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProvidersLdapByIDDefault creates a GetProvidersLdapByIDDefault with default headers values
func NewGetProvidersLdapByIDDefault(code int) *GetProvidersLdapByIDDefault {
	return &GetProvidersLdapByIDDefault{
		_statusCode: code,
	}
}

/*GetProvidersLdapByIDDefault handles this case with default header values.

Unexpected error
*/
type GetProvidersLdapByIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get providers ldap by Id default response
func (o *GetProvidersLdapByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetProvidersLdapByIDDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/auth/providers/ldap/{ProvidersLdapId}][%d] getProvidersLdapById default  %+v", o._statusCode, o.Payload)
}

func (o *GetProvidersLdapByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
