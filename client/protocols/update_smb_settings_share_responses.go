// Code generated by go-swagger; DO NOT EDIT.

package protocols

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// UpdateSmbSettingsShareReader is a Reader for the UpdateSmbSettingsShare structure.
type UpdateSmbSettingsShareReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSmbSettingsShareReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateSmbSettingsShareNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateSmbSettingsShareDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateSmbSettingsShareNoContent creates a UpdateSmbSettingsShareNoContent with default headers values
func NewUpdateSmbSettingsShareNoContent() *UpdateSmbSettingsShareNoContent {
	return &UpdateSmbSettingsShareNoContent{}
}

/*UpdateSmbSettingsShareNoContent handles this case with default header values.

Success.
*/
type UpdateSmbSettingsShareNoContent struct {
}

func (o *UpdateSmbSettingsShareNoContent) Error() string {
	return fmt.Sprintf("[PUT /platform/3/protocols/smb/settings/share][%d] updateSmbSettingsShareNoContent ", 204)
}

func (o *UpdateSmbSettingsShareNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSmbSettingsShareDefault creates a UpdateSmbSettingsShareDefault with default headers values
func NewUpdateSmbSettingsShareDefault(code int) *UpdateSmbSettingsShareDefault {
	return &UpdateSmbSettingsShareDefault{
		_statusCode: code,
	}
}

/*UpdateSmbSettingsShareDefault handles this case with default header values.

Unexpected error
*/
type UpdateSmbSettingsShareDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update smb settings share default response
func (o *UpdateSmbSettingsShareDefault) Code() int {
	return o._statusCode
}

func (o *UpdateSmbSettingsShareDefault) Error() string {
	return fmt.Sprintf("[PUT /platform/3/protocols/smb/settings/share][%d] updateSmbSettingsShare default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSmbSettingsShareDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
