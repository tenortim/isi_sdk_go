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

// GetProvidersLocalByIDReader is a Reader for the GetProvidersLocalByID structure.
type GetProvidersLocalByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProvidersLocalByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetProvidersLocalByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetProvidersLocalByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProvidersLocalByIDOK creates a GetProvidersLocalByIDOK with default headers values
func NewGetProvidersLocalByIDOK() *GetProvidersLocalByIDOK {
	return &GetProvidersLocalByIDOK{}
}

/*GetProvidersLocalByIDOK handles this case with default header values.

Retrieve the local provider.
*/
type GetProvidersLocalByIDOK struct {
	Payload *models.ProvidersLocal
}

func (o *GetProvidersLocalByIDOK) Error() string {
	return fmt.Sprintf("[GET /platform/1/auth/providers/local/{ProvidersLocalId}][%d] getProvidersLocalByIdOK  %+v", 200, o.Payload)
}

func (o *GetProvidersLocalByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProvidersLocal)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProvidersLocalByIDDefault creates a GetProvidersLocalByIDDefault with default headers values
func NewGetProvidersLocalByIDDefault(code int) *GetProvidersLocalByIDDefault {
	return &GetProvidersLocalByIDDefault{
		_statusCode: code,
	}
}

/*GetProvidersLocalByIDDefault handles this case with default header values.

Unexpected error
*/
type GetProvidersLocalByIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get providers local by Id default response
func (o *GetProvidersLocalByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetProvidersLocalByIDDefault) Error() string {
	return fmt.Sprintf("[GET /platform/1/auth/providers/local/{ProvidersLocalId}][%d] getProvidersLocalById default  %+v", o._statusCode, o.Payload)
}

func (o *GetProvidersLocalByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}