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

// UpdateHdfsSettingsReader is a Reader for the UpdateHdfsSettings structure.
type UpdateHdfsSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateHdfsSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateHdfsSettingsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateHdfsSettingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateHdfsSettingsNoContent creates a UpdateHdfsSettingsNoContent with default headers values
func NewUpdateHdfsSettingsNoContent() *UpdateHdfsSettingsNoContent {
	return &UpdateHdfsSettingsNoContent{}
}

/*UpdateHdfsSettingsNoContent handles this case with default header values.

Success.
*/
type UpdateHdfsSettingsNoContent struct {
}

func (o *UpdateHdfsSettingsNoContent) Error() string {
	return fmt.Sprintf("[PUT /platform/3/protocols/hdfs/settings][%d] updateHdfsSettingsNoContent ", 204)
}

func (o *UpdateHdfsSettingsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateHdfsSettingsDefault creates a UpdateHdfsSettingsDefault with default headers values
func NewUpdateHdfsSettingsDefault(code int) *UpdateHdfsSettingsDefault {
	return &UpdateHdfsSettingsDefault{
		_statusCode: code,
	}
}

/*UpdateHdfsSettingsDefault handles this case with default header values.

Unexpected error
*/
type UpdateHdfsSettingsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update hdfs settings default response
func (o *UpdateHdfsSettingsDefault) Code() int {
	return o._statusCode
}

func (o *UpdateHdfsSettingsDefault) Error() string {
	return fmt.Sprintf("[PUT /platform/3/protocols/hdfs/settings][%d] updateHdfsSettings default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateHdfsSettingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
