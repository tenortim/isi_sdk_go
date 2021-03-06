// Code generated by go-swagger; DO NOT EDIT.

package sync

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// GetTargetReportsReader is a Reader for the GetTargetReports structure.
type GetTargetReportsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTargetReportsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTargetReportsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetTargetReportsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTargetReportsOK creates a GetTargetReportsOK with default headers values
func NewGetTargetReportsOK() *GetTargetReportsOK {
	return &GetTargetReportsOK{}
}

/*GetTargetReportsOK handles this case with default header values.

Get a list of SyncIQ target reports.  By default 10 reports are returned per policy, unless otherwise specified by 'reports_per_policy'.
*/
type GetTargetReportsOK struct {
	Payload *models.TargetReportsExtended
}

func (o *GetTargetReportsOK) Error() string {
	return fmt.Sprintf("[GET /platform/1/sync/target/reports][%d] getTargetReportsOK  %+v", 200, o.Payload)
}

func (o *GetTargetReportsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TargetReportsExtended)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTargetReportsDefault creates a GetTargetReportsDefault with default headers values
func NewGetTargetReportsDefault(code int) *GetTargetReportsDefault {
	return &GetTargetReportsDefault{
		_statusCode: code,
	}
}

/*GetTargetReportsDefault handles this case with default header values.

Unexpected error
*/
type GetTargetReportsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get target reports default response
func (o *GetTargetReportsDefault) Code() int {
	return o._statusCode
}

func (o *GetTargetReportsDefault) Error() string {
	return fmt.Sprintf("[GET /platform/1/sync/target/reports][%d] getTargetReports default  %+v", o._statusCode, o.Payload)
}

func (o *GetTargetReportsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
