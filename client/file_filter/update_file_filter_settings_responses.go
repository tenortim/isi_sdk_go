// Code generated by go-swagger; DO NOT EDIT.

package file_filter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// UpdateFileFilterSettingsReader is a Reader for the UpdateFileFilterSettings structure.
type UpdateFileFilterSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateFileFilterSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateFileFilterSettingsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateFileFilterSettingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateFileFilterSettingsNoContent creates a UpdateFileFilterSettingsNoContent with default headers values
func NewUpdateFileFilterSettingsNoContent() *UpdateFileFilterSettingsNoContent {
	return &UpdateFileFilterSettingsNoContent{}
}

/*UpdateFileFilterSettingsNoContent handles this case with default header values.

Success.
*/
type UpdateFileFilterSettingsNoContent struct {
}

func (o *UpdateFileFilterSettingsNoContent) Error() string {
	return fmt.Sprintf("[PUT /platform/3/file-filter/settings][%d] updateFileFilterSettingsNoContent ", 204)
}

func (o *UpdateFileFilterSettingsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateFileFilterSettingsDefault creates a UpdateFileFilterSettingsDefault with default headers values
func NewUpdateFileFilterSettingsDefault(code int) *UpdateFileFilterSettingsDefault {
	return &UpdateFileFilterSettingsDefault{
		_statusCode: code,
	}
}

/*UpdateFileFilterSettingsDefault handles this case with default header values.

Unexpected error
*/
type UpdateFileFilterSettingsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update file filter settings default response
func (o *UpdateFileFilterSettingsDefault) Code() int {
	return o._statusCode
}

func (o *UpdateFileFilterSettingsDefault) Error() string {
	return fmt.Sprintf("[PUT /platform/3/file-filter/settings][%d] updateFileFilterSettings default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateFileFilterSettingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
