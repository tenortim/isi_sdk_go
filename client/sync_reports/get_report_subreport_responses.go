// Code generated by go-swagger; DO NOT EDIT.

package sync_reports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// GetReportSubreportReader is a Reader for the GetReportSubreport structure.
type GetReportSubreportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReportSubreportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetReportSubreportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetReportSubreportDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetReportSubreportOK creates a GetReportSubreportOK with default headers values
func NewGetReportSubreportOK() *GetReportSubreportOK {
	return &GetReportSubreportOK{}
}

/*GetReportSubreportOK handles this case with default header values.

View a single SyncIQ subreport.
*/
type GetReportSubreportOK struct {
	Payload *models.ReportSubreports
}

func (o *GetReportSubreportOK) Error() string {
	return fmt.Sprintf("[GET /platform/1/sync/reports/{Rid}/subreports/{ReportSubreportId}][%d] getReportSubreportOK  %+v", 200, o.Payload)
}

func (o *GetReportSubreportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReportSubreports)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportSubreportDefault creates a GetReportSubreportDefault with default headers values
func NewGetReportSubreportDefault(code int) *GetReportSubreportDefault {
	return &GetReportSubreportDefault{
		_statusCode: code,
	}
}

/*GetReportSubreportDefault handles this case with default header values.

Unexpected error
*/
type GetReportSubreportDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get report subreport default response
func (o *GetReportSubreportDefault) Code() int {
	return o._statusCode
}

func (o *GetReportSubreportDefault) Error() string {
	return fmt.Sprintf("[GET /platform/1/sync/reports/{Rid}/subreports/{ReportSubreportId}][%d] getReportSubreport default  %+v", o._statusCode, o.Payload)
}

func (o *GetReportSubreportDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
