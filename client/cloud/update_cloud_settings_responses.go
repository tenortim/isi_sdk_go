// Code generated by go-swagger; DO NOT EDIT.

package cloud

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// UpdateCloudSettingsReader is a Reader for the UpdateCloudSettings structure.
type UpdateCloudSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCloudSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateCloudSettingsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateCloudSettingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateCloudSettingsNoContent creates a UpdateCloudSettingsNoContent with default headers values
func NewUpdateCloudSettingsNoContent() *UpdateCloudSettingsNoContent {
	return &UpdateCloudSettingsNoContent{}
}

/*UpdateCloudSettingsNoContent handles this case with default header values.

Success.
*/
type UpdateCloudSettingsNoContent struct {
}

func (o *UpdateCloudSettingsNoContent) Error() string {
	return fmt.Sprintf("[PUT /platform/3/cloud/settings][%d] updateCloudSettingsNoContent ", 204)
}

func (o *UpdateCloudSettingsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateCloudSettingsDefault creates a UpdateCloudSettingsDefault with default headers values
func NewUpdateCloudSettingsDefault(code int) *UpdateCloudSettingsDefault {
	return &UpdateCloudSettingsDefault{
		_statusCode: code,
	}
}

/*UpdateCloudSettingsDefault handles this case with default header values.

Unexpected error
*/
type UpdateCloudSettingsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update cloud settings default response
func (o *UpdateCloudSettingsDefault) Code() int {
	return o._statusCode
}

func (o *UpdateCloudSettingsDefault) Error() string {
	return fmt.Sprintf("[PUT /platform/3/cloud/settings][%d] updateCloudSettings default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateCloudSettingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}