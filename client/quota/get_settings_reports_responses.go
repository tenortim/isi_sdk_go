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

// GetSettingsReportsReader is a Reader for the GetSettingsReports structure.
type GetSettingsReportsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSettingsReportsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSettingsReportsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetSettingsReportsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSettingsReportsOK creates a GetSettingsReportsOK with default headers values
func NewGetSettingsReportsOK() *GetSettingsReportsOK {
	return &GetSettingsReportsOK{}
}

/*GetSettingsReportsOK handles this case with default header values.

List all settings.
*/
type GetSettingsReportsOK struct {
	Payload *models.SettingsReports
}

func (o *GetSettingsReportsOK) Error() string {
	return fmt.Sprintf("[GET /platform/1/quota/settings/reports][%d] getSettingsReportsOK  %+v", 200, o.Payload)
}

func (o *GetSettingsReportsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SettingsReports)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSettingsReportsDefault creates a GetSettingsReportsDefault with default headers values
func NewGetSettingsReportsDefault(code int) *GetSettingsReportsDefault {
	return &GetSettingsReportsDefault{
		_statusCode: code,
	}
}

/*GetSettingsReportsDefault handles this case with default header values.

Unexpected error
*/
type GetSettingsReportsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get settings reports default response
func (o *GetSettingsReportsDefault) Code() int {
	return o._statusCode
}

func (o *GetSettingsReportsDefault) Error() string {
	return fmt.Sprintf("[GET /platform/1/quota/settings/reports][%d] getSettingsReports default  %+v", o._statusCode, o.Payload)
}

func (o *GetSettingsReportsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}