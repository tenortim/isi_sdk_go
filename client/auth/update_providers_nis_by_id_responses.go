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

// UpdateProvidersNisByIDReader is a Reader for the UpdateProvidersNisByID structure.
type UpdateProvidersNisByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProvidersNisByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateProvidersNisByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateProvidersNisByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateProvidersNisByIDNoContent creates a UpdateProvidersNisByIDNoContent with default headers values
func NewUpdateProvidersNisByIDNoContent() *UpdateProvidersNisByIDNoContent {
	return &UpdateProvidersNisByIDNoContent{}
}

/*UpdateProvidersNisByIDNoContent handles this case with default header values.

Success.
*/
type UpdateProvidersNisByIDNoContent struct {
}

func (o *UpdateProvidersNisByIDNoContent) Error() string {
	return fmt.Sprintf("[PUT /platform/3/auth/providers/nis/{ProvidersNisId}][%d] updateProvidersNisByIdNoContent ", 204)
}

func (o *UpdateProvidersNisByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateProvidersNisByIDDefault creates a UpdateProvidersNisByIDDefault with default headers values
func NewUpdateProvidersNisByIDDefault(code int) *UpdateProvidersNisByIDDefault {
	return &UpdateProvidersNisByIDDefault{
		_statusCode: code,
	}
}

/*UpdateProvidersNisByIDDefault handles this case with default header values.

Unexpected error
*/
type UpdateProvidersNisByIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update providers nis by Id default response
func (o *UpdateProvidersNisByIDDefault) Code() int {
	return o._statusCode
}

func (o *UpdateProvidersNisByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /platform/3/auth/providers/nis/{ProvidersNisId}][%d] updateProvidersNisById default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateProvidersNisByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
