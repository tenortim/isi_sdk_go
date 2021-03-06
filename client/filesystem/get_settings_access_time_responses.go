// Code generated by go-swagger; DO NOT EDIT.

package filesystem

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// GetSettingsAccessTimeReader is a Reader for the GetSettingsAccessTime structure.
type GetSettingsAccessTimeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSettingsAccessTimeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSettingsAccessTimeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetSettingsAccessTimeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSettingsAccessTimeOK creates a GetSettingsAccessTimeOK with default headers values
func NewGetSettingsAccessTimeOK() *GetSettingsAccessTimeOK {
	return &GetSettingsAccessTimeOK{}
}

/*GetSettingsAccessTimeOK handles this case with default header values.

Retrieve settings for access time.
*/
type GetSettingsAccessTimeOK struct {
	Payload *models.SettingsAccessTime
}

func (o *GetSettingsAccessTimeOK) Error() string {
	return fmt.Sprintf("[GET /platform/1/filesystem/settings/access-time][%d] getSettingsAccessTimeOK  %+v", 200, o.Payload)
}

func (o *GetSettingsAccessTimeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SettingsAccessTime)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSettingsAccessTimeDefault creates a GetSettingsAccessTimeDefault with default headers values
func NewGetSettingsAccessTimeDefault(code int) *GetSettingsAccessTimeDefault {
	return &GetSettingsAccessTimeDefault{
		_statusCode: code,
	}
}

/*GetSettingsAccessTimeDefault handles this case with default header values.

Unexpected error
*/
type GetSettingsAccessTimeDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get settings access time default response
func (o *GetSettingsAccessTimeDefault) Code() int {
	return o._statusCode
}

func (o *GetSettingsAccessTimeDefault) Error() string {
	return fmt.Sprintf("[GET /platform/1/filesystem/settings/access-time][%d] getSettingsAccessTime default  %+v", o._statusCode, o.Payload)
}

func (o *GetSettingsAccessTimeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
