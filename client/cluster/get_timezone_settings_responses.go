// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// GetTimezoneSettingsReader is a Reader for the GetTimezoneSettings structure.
type GetTimezoneSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTimezoneSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTimezoneSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetTimezoneSettingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTimezoneSettingsOK creates a GetTimezoneSettingsOK with default headers values
func NewGetTimezoneSettingsOK() *GetTimezoneSettingsOK {
	return &GetTimezoneSettingsOK{}
}

/*GetTimezoneSettingsOK handles this case with default header values.

Retrieve the cluster timezone.
*/
type GetTimezoneSettingsOK struct {
	Payload *models.TimezoneSettings
}

func (o *GetTimezoneSettingsOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/cluster/timezone/settings][%d] getTimezoneSettingsOK  %+v", 200, o.Payload)
}

func (o *GetTimezoneSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TimezoneSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTimezoneSettingsDefault creates a GetTimezoneSettingsDefault with default headers values
func NewGetTimezoneSettingsDefault(code int) *GetTimezoneSettingsDefault {
	return &GetTimezoneSettingsDefault{
		_statusCode: code,
	}
}

/*GetTimezoneSettingsDefault handles this case with default header values.

Unexpected error
*/
type GetTimezoneSettingsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get timezone settings default response
func (o *GetTimezoneSettingsDefault) Code() int {
	return o._statusCode
}

func (o *GetTimezoneSettingsDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/cluster/timezone/settings][%d] getTimezoneSettings default  %+v", o._statusCode, o.Payload)
}

func (o *GetTimezoneSettingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}