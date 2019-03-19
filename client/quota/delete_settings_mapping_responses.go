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

// DeleteSettingsMappingReader is a Reader for the DeleteSettingsMapping structure.
type DeleteSettingsMappingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSettingsMappingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSettingsMappingNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteSettingsMappingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSettingsMappingNoContent creates a DeleteSettingsMappingNoContent with default headers values
func NewDeleteSettingsMappingNoContent() *DeleteSettingsMappingNoContent {
	return &DeleteSettingsMappingNoContent{}
}

/*DeleteSettingsMappingNoContent handles this case with default header values.

Success.
*/
type DeleteSettingsMappingNoContent struct {
}

func (o *DeleteSettingsMappingNoContent) Error() string {
	return fmt.Sprintf("[DELETE /platform/1/quota/settings/mappings/{SettingsMappingId}][%d] deleteSettingsMappingNoContent ", 204)
}

func (o *DeleteSettingsMappingNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSettingsMappingDefault creates a DeleteSettingsMappingDefault with default headers values
func NewDeleteSettingsMappingDefault(code int) *DeleteSettingsMappingDefault {
	return &DeleteSettingsMappingDefault{
		_statusCode: code,
	}
}

/*DeleteSettingsMappingDefault handles this case with default header values.

Unexpected error
*/
type DeleteSettingsMappingDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete settings mapping default response
func (o *DeleteSettingsMappingDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSettingsMappingDefault) Error() string {
	return fmt.Sprintf("[DELETE /platform/1/quota/settings/mappings/{SettingsMappingId}][%d] deleteSettingsMapping default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSettingsMappingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}