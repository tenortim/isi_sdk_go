// Code generated by go-swagger; DO NOT EDIT.

package sync_target

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// GetReportsReportSubreportsReader is a Reader for the GetReportsReportSubreports structure.
type GetReportsReportSubreportsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReportsReportSubreportsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetReportsReportSubreportsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetReportsReportSubreportsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetReportsReportSubreportsOK creates a GetReportsReportSubreportsOK with default headers values
func NewGetReportsReportSubreportsOK() *GetReportsReportSubreportsOK {
	return &GetReportsReportSubreportsOK{}
}

/*GetReportsReportSubreportsOK handles this case with default header values.

Get a list of SyncIQ target subreports for a report.
*/
type GetReportsReportSubreportsOK struct {
	Payload *models.ReportsReportSubreportsExtended
}

func (o *GetReportsReportSubreportsOK) Error() string {
	return fmt.Sprintf("[GET /platform/1/sync/target/reports/{Rid}/subreports][%d] getReportsReportSubreportsOK  %+v", 200, o.Payload)
}

func (o *GetReportsReportSubreportsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReportsReportSubreportsExtended)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportsReportSubreportsDefault creates a GetReportsReportSubreportsDefault with default headers values
func NewGetReportsReportSubreportsDefault(code int) *GetReportsReportSubreportsDefault {
	return &GetReportsReportSubreportsDefault{
		_statusCode: code,
	}
}

/*GetReportsReportSubreportsDefault handles this case with default header values.

Unexpected error
*/
type GetReportsReportSubreportsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get reports report subreports default response
func (o *GetReportsReportSubreportsDefault) Code() int {
	return o._statusCode
}

func (o *GetReportsReportSubreportsDefault) Error() string {
	return fmt.Sprintf("[GET /platform/1/sync/target/reports/{Rid}/subreports][%d] getReportsReportSubreports default  %+v", o._statusCode, o.Payload)
}

func (o *GetReportsReportSubreportsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
