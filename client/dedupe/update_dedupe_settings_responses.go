// Code generated by go-swagger; DO NOT EDIT.

package dedupe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// UpdateDedupeSettingsReader is a Reader for the UpdateDedupeSettings structure.
type UpdateDedupeSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDedupeSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateDedupeSettingsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateDedupeSettingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateDedupeSettingsNoContent creates a UpdateDedupeSettingsNoContent with default headers values
func NewUpdateDedupeSettingsNoContent() *UpdateDedupeSettingsNoContent {
	return &UpdateDedupeSettingsNoContent{}
}

/*UpdateDedupeSettingsNoContent handles this case with default header values.

Success.
*/
type UpdateDedupeSettingsNoContent struct {
}

func (o *UpdateDedupeSettingsNoContent) Error() string {
	return fmt.Sprintf("[PUT /platform/1/dedupe/settings][%d] updateDedupeSettingsNoContent ", 204)
}

func (o *UpdateDedupeSettingsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateDedupeSettingsDefault creates a UpdateDedupeSettingsDefault with default headers values
func NewUpdateDedupeSettingsDefault(code int) *UpdateDedupeSettingsDefault {
	return &UpdateDedupeSettingsDefault{
		_statusCode: code,
	}
}

/*UpdateDedupeSettingsDefault handles this case with default header values.

Unexpected error
*/
type UpdateDedupeSettingsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update dedupe settings default response
func (o *UpdateDedupeSettingsDefault) Code() int {
	return o._statusCode
}

func (o *UpdateDedupeSettingsDefault) Error() string {
	return fmt.Sprintf("[PUT /platform/1/dedupe/settings][%d] updateDedupeSettings default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateDedupeSettingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}