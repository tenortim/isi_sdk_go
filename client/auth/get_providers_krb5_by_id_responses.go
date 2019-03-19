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

// GetProvidersKrb5ByIDReader is a Reader for the GetProvidersKrb5ByID structure.
type GetProvidersKrb5ByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProvidersKrb5ByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetProvidersKrb5ByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetProvidersKrb5ByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProvidersKrb5ByIDOK creates a GetProvidersKrb5ByIDOK with default headers values
func NewGetProvidersKrb5ByIDOK() *GetProvidersKrb5ByIDOK {
	return &GetProvidersKrb5ByIDOK{}
}

/*GetProvidersKrb5ByIDOK handles this case with default header values.

Retrieve the KRB5 provider.
*/
type GetProvidersKrb5ByIDOK struct {
	Payload *models.ProvidersKrb5
}

func (o *GetProvidersKrb5ByIDOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/auth/providers/krb5/{ProvidersKrb5Id}][%d] getProvidersKrb5ByIdOK  %+v", 200, o.Payload)
}

func (o *GetProvidersKrb5ByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProvidersKrb5)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProvidersKrb5ByIDDefault creates a GetProvidersKrb5ByIDDefault with default headers values
func NewGetProvidersKrb5ByIDDefault(code int) *GetProvidersKrb5ByIDDefault {
	return &GetProvidersKrb5ByIDDefault{
		_statusCode: code,
	}
}

/*GetProvidersKrb5ByIDDefault handles this case with default header values.

Unexpected error
*/
type GetProvidersKrb5ByIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get providers krb5 by Id default response
func (o *GetProvidersKrb5ByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetProvidersKrb5ByIDDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/auth/providers/krb5/{ProvidersKrb5Id}][%d] getProvidersKrb5ById default  %+v", o._statusCode, o.Payload)
}

func (o *GetProvidersKrb5ByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
