// Code generated by go-swagger; DO NOT EDIT.

package cloud

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// ListSettingsReportingEulaReader is a Reader for the ListSettingsReportingEula structure.
type ListSettingsReportingEulaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSettingsReportingEulaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListSettingsReportingEulaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListSettingsReportingEulaDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListSettingsReportingEulaOK creates a ListSettingsReportingEulaOK with default headers values
func NewListSettingsReportingEulaOK() *ListSettingsReportingEulaOK {
	return &ListSettingsReportingEulaOK{}
}

/*ListSettingsReportingEulaOK handles this case with default header values.

View telemetry collection EULA acceptance and content URI.
*/
type ListSettingsReportingEulaOK struct {
	Payload *models.SettingsReportingEulaItem
}

func (o *ListSettingsReportingEulaOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/cloud/settings/reporting-eula][%d] listSettingsReportingEulaOK  %+v", 200, o.Payload)
}

func (o *ListSettingsReportingEulaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SettingsReportingEulaItem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSettingsReportingEulaDefault creates a ListSettingsReportingEulaDefault with default headers values
func NewListSettingsReportingEulaDefault(code int) *ListSettingsReportingEulaDefault {
	return &ListSettingsReportingEulaDefault{
		_statusCode: code,
	}
}

/*ListSettingsReportingEulaDefault handles this case with default header values.

Unexpected error
*/
type ListSettingsReportingEulaDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list settings reporting eula default response
func (o *ListSettingsReportingEulaDefault) Code() int {
	return o._statusCode
}

func (o *ListSettingsReportingEulaDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/cloud/settings/reporting-eula][%d] listSettingsReportingEula default  %+v", o._statusCode, o.Payload)
}

func (o *ListSettingsReportingEulaDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
