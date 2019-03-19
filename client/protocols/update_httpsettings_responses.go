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

// UpdateHTTPSettingsReader is a Reader for the UpdateHTTPSettings structure.
type UpdateHTTPSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateHTTPSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateHTTPSettingsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateHTTPSettingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateHTTPSettingsNoContent creates a UpdateHTTPSettingsNoContent with default headers values
func NewUpdateHTTPSettingsNoContent() *UpdateHTTPSettingsNoContent {
	return &UpdateHTTPSettingsNoContent{}
}

/*UpdateHTTPSettingsNoContent handles this case with default header values.

Success.
*/
type UpdateHTTPSettingsNoContent struct {
}

func (o *UpdateHTTPSettingsNoContent) Error() string {
	return fmt.Sprintf("[PUT /platform/3/protocols/http/settings][%d] updateHttpSettingsNoContent ", 204)
}

func (o *UpdateHTTPSettingsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateHTTPSettingsDefault creates a UpdateHTTPSettingsDefault with default headers values
func NewUpdateHTTPSettingsDefault(code int) *UpdateHTTPSettingsDefault {
	return &UpdateHTTPSettingsDefault{
		_statusCode: code,
	}
}

/*UpdateHTTPSettingsDefault handles this case with default header values.

Unexpected error
*/
type UpdateHTTPSettingsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update Http settings default response
func (o *UpdateHTTPSettingsDefault) Code() int {
	return o._statusCode
}

func (o *UpdateHTTPSettingsDefault) Error() string {
	return fmt.Sprintf("[PUT /platform/3/protocols/http/settings][%d] updateHttpSettings default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateHTTPSettingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}