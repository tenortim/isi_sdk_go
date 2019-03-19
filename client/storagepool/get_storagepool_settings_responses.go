// Code generated by go-swagger; DO NOT EDIT.

package storagepool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// GetStoragepoolSettingsReader is a Reader for the GetStoragepoolSettings structure.
type GetStoragepoolSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStoragepoolSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetStoragepoolSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetStoragepoolSettingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStoragepoolSettingsOK creates a GetStoragepoolSettingsOK with default headers values
func NewGetStoragepoolSettingsOK() *GetStoragepoolSettingsOK {
	return &GetStoragepoolSettingsOK{}
}

/*GetStoragepoolSettingsOK handles this case with default header values.

List all settings.
*/
type GetStoragepoolSettingsOK struct {
	Payload *models.StoragepoolSettings
}

func (o *GetStoragepoolSettingsOK) Error() string {
	return fmt.Sprintf("[GET /platform/1/storagepool/settings][%d] getStoragepoolSettingsOK  %+v", 200, o.Payload)
}

func (o *GetStoragepoolSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StoragepoolSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStoragepoolSettingsDefault creates a GetStoragepoolSettingsDefault with default headers values
func NewGetStoragepoolSettingsDefault(code int) *GetStoragepoolSettingsDefault {
	return &GetStoragepoolSettingsDefault{
		_statusCode: code,
	}
}

/*GetStoragepoolSettingsDefault handles this case with default header values.

Unexpected error
*/
type GetStoragepoolSettingsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get storagepool settings default response
func (o *GetStoragepoolSettingsDefault) Code() int {
	return o._statusCode
}

func (o *GetStoragepoolSettingsDefault) Error() string {
	return fmt.Sprintf("[GET /platform/1/storagepool/settings][%d] getStoragepoolSettings default  %+v", o._statusCode, o.Payload)
}

func (o *GetStoragepoolSettingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}