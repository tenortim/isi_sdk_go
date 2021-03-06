// Code generated by go-swagger; DO NOT EDIT.

package quota

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// ListSettingsMappingsReader is a Reader for the ListSettingsMappings structure.
type ListSettingsMappingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSettingsMappingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListSettingsMappingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListSettingsMappingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListSettingsMappingsOK creates a ListSettingsMappingsOK with default headers values
func NewListSettingsMappingsOK() *ListSettingsMappingsOK {
	return &ListSettingsMappingsOK{}
}

/*ListSettingsMappingsOK handles this case with default header values.

List all rules.
*/
type ListSettingsMappingsOK struct {
	Payload *models.SettingsMappings
}

func (o *ListSettingsMappingsOK) Error() string {
	return fmt.Sprintf("[GET /platform/1/quota/settings/mappings][%d] listSettingsMappingsOK  %+v", 200, o.Payload)
}

func (o *ListSettingsMappingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SettingsMappings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSettingsMappingsDefault creates a ListSettingsMappingsDefault with default headers values
func NewListSettingsMappingsDefault(code int) *ListSettingsMappingsDefault {
	return &ListSettingsMappingsDefault{
		_statusCode: code,
	}
}

/*ListSettingsMappingsDefault handles this case with default header values.

Unexpected error
*/
type ListSettingsMappingsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list settings mappings default response
func (o *ListSettingsMappingsDefault) Code() int {
	return o._statusCode
}

func (o *ListSettingsMappingsDefault) Error() string {
	return fmt.Sprintf("[GET /platform/1/quota/settings/mappings][%d] listSettingsMappings default  %+v", o._statusCode, o.Payload)
}

func (o *ListSettingsMappingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
