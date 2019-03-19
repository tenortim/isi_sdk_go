// Code generated by go-swagger; DO NOT EDIT.

package audit

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// GetSettingsGlobalReader is a Reader for the GetSettingsGlobal structure.
type GetSettingsGlobalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSettingsGlobalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSettingsGlobalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetSettingsGlobalDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSettingsGlobalOK creates a GetSettingsGlobalOK with default headers values
func NewGetSettingsGlobalOK() *GetSettingsGlobalOK {
	return &GetSettingsGlobalOK{}
}

/*GetSettingsGlobalOK handles this case with default header values.

View Global Audit settings.
*/
type GetSettingsGlobalOK struct {
	Payload *models.SettingsGlobalExtended
}

func (o *GetSettingsGlobalOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/audit/settings/global][%d] getSettingsGlobalOK  %+v", 200, o.Payload)
}

func (o *GetSettingsGlobalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SettingsGlobalExtended)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSettingsGlobalDefault creates a GetSettingsGlobalDefault with default headers values
func NewGetSettingsGlobalDefault(code int) *GetSettingsGlobalDefault {
	return &GetSettingsGlobalDefault{
		_statusCode: code,
	}
}

/*GetSettingsGlobalDefault handles this case with default header values.

Unexpected error
*/
type GetSettingsGlobalDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get settings global default response
func (o *GetSettingsGlobalDefault) Code() int {
	return o._statusCode
}

func (o *GetSettingsGlobalDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/audit/settings/global][%d] getSettingsGlobal default  %+v", o._statusCode, o.Payload)
}

func (o *GetSettingsGlobalDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}