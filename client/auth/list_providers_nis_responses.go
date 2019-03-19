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

// ListProvidersNisReader is a Reader for the ListProvidersNis structure.
type ListProvidersNisReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProvidersNisReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListProvidersNisOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListProvidersNisDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProvidersNisOK creates a ListProvidersNisOK with default headers values
func NewListProvidersNisOK() *ListProvidersNisOK {
	return &ListProvidersNisOK{}
}

/*ListProvidersNisOK handles this case with default header values.

List all NIS providers.
*/
type ListProvidersNisOK struct {
	Payload *models.ProvidersNisExtended
}

func (o *ListProvidersNisOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/auth/providers/nis][%d] listProvidersNisOK  %+v", 200, o.Payload)
}

func (o *ListProvidersNisOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProvidersNisExtended)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProvidersNisDefault creates a ListProvidersNisDefault with default headers values
func NewListProvidersNisDefault(code int) *ListProvidersNisDefault {
	return &ListProvidersNisDefault{
		_statusCode: code,
	}
}

/*ListProvidersNisDefault handles this case with default header values.

Unexpected error
*/
type ListProvidersNisDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list providers nis default response
func (o *ListProvidersNisDefault) Code() int {
	return o._statusCode
}

func (o *ListProvidersNisDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/auth/providers/nis][%d] listProvidersNis default  %+v", o._statusCode, o.Payload)
}

func (o *ListProvidersNisDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
