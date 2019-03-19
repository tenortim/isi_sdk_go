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

// GetProvidersLocalReader is a Reader for the GetProvidersLocal structure.
type GetProvidersLocalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProvidersLocalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetProvidersLocalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetProvidersLocalDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProvidersLocalOK creates a GetProvidersLocalOK with default headers values
func NewGetProvidersLocalOK() *GetProvidersLocalOK {
	return &GetProvidersLocalOK{}
}

/*GetProvidersLocalOK handles this case with default header values.

List all local providers.
*/
type GetProvidersLocalOK struct {
	Payload *models.ProvidersLocal
}

func (o *GetProvidersLocalOK) Error() string {
	return fmt.Sprintf("[GET /platform/1/auth/providers/local][%d] getProvidersLocalOK  %+v", 200, o.Payload)
}

func (o *GetProvidersLocalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProvidersLocal)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProvidersLocalDefault creates a GetProvidersLocalDefault with default headers values
func NewGetProvidersLocalDefault(code int) *GetProvidersLocalDefault {
	return &GetProvidersLocalDefault{
		_statusCode: code,
	}
}

/*GetProvidersLocalDefault handles this case with default header values.

Unexpected error
*/
type GetProvidersLocalDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get providers local default response
func (o *GetProvidersLocalDefault) Code() int {
	return o._statusCode
}

func (o *GetProvidersLocalDefault) Error() string {
	return fmt.Sprintf("[GET /platform/1/auth/providers/local][%d] getProvidersLocal default  %+v", o._statusCode, o.Payload)
}

func (o *GetProvidersLocalDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}