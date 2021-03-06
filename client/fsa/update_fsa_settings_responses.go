// Code generated by go-swagger; DO NOT EDIT.

package fsa

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// UpdateFsaSettingsReader is a Reader for the UpdateFsaSettings structure.
type UpdateFsaSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateFsaSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateFsaSettingsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateFsaSettingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateFsaSettingsNoContent creates a UpdateFsaSettingsNoContent with default headers values
func NewUpdateFsaSettingsNoContent() *UpdateFsaSettingsNoContent {
	return &UpdateFsaSettingsNoContent{}
}

/*UpdateFsaSettingsNoContent handles this case with default header values.

Success.
*/
type UpdateFsaSettingsNoContent struct {
}

func (o *UpdateFsaSettingsNoContent) Error() string {
	return fmt.Sprintf("[PUT /platform/1/fsa/settings][%d] updateFsaSettingsNoContent ", 204)
}

func (o *UpdateFsaSettingsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateFsaSettingsDefault creates a UpdateFsaSettingsDefault with default headers values
func NewUpdateFsaSettingsDefault(code int) *UpdateFsaSettingsDefault {
	return &UpdateFsaSettingsDefault{
		_statusCode: code,
	}
}

/*UpdateFsaSettingsDefault handles this case with default header values.

Unexpected error
*/
type UpdateFsaSettingsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update fsa settings default response
func (o *UpdateFsaSettingsDefault) Code() int {
	return o._statusCode
}

func (o *UpdateFsaSettingsDefault) Error() string {
	return fmt.Sprintf("[PUT /platform/1/fsa/settings][%d] updateFsaSettings default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateFsaSettingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
