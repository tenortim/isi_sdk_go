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

// GetFsaSettingsReader is a Reader for the GetFsaSettings structure.
type GetFsaSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFsaSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFsaSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetFsaSettingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFsaSettingsOK creates a GetFsaSettingsOK with default headers values
func NewGetFsaSettingsOK() *GetFsaSettingsOK {
	return &GetFsaSettingsOK{}
}

/*GetFsaSettingsOK handles this case with default header values.

List all settings.
*/
type GetFsaSettingsOK struct {
	Payload *models.FsaSettings
}

func (o *GetFsaSettingsOK) Error() string {
	return fmt.Sprintf("[GET /platform/1/fsa/settings][%d] getFsaSettingsOK  %+v", 200, o.Payload)
}

func (o *GetFsaSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsaSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFsaSettingsDefault creates a GetFsaSettingsDefault with default headers values
func NewGetFsaSettingsDefault(code int) *GetFsaSettingsDefault {
	return &GetFsaSettingsDefault{
		_statusCode: code,
	}
}

/*GetFsaSettingsDefault handles this case with default header values.

Unexpected error
*/
type GetFsaSettingsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get fsa settings default response
func (o *GetFsaSettingsDefault) Code() int {
	return o._statusCode
}

func (o *GetFsaSettingsDefault) Error() string {
	return fmt.Sprintf("[GET /platform/1/fsa/settings][%d] getFsaSettings default  %+v", o._statusCode, o.Payload)
}

func (o *GetFsaSettingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
