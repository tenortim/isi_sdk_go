// Code generated by go-swagger; DO NOT EDIT.

package dedupe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// GetDedupeReportReader is a Reader for the GetDedupeReport structure.
type GetDedupeReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDedupeReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDedupeReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetDedupeReportDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDedupeReportOK creates a GetDedupeReportOK with default headers values
func NewGetDedupeReportOK() *GetDedupeReportOK {
	return &GetDedupeReportOK{}
}

/*GetDedupeReportOK handles this case with default header values.

Retrieve a report for a single dedupe job.
*/
type GetDedupeReportOK struct {
	Payload *models.DedupeReports
}

func (o *GetDedupeReportOK) Error() string {
	return fmt.Sprintf("[GET /platform/1/dedupe/reports/{DedupeReportId}][%d] getDedupeReportOK  %+v", 200, o.Payload)
}

func (o *GetDedupeReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DedupeReports)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDedupeReportDefault creates a GetDedupeReportDefault with default headers values
func NewGetDedupeReportDefault(code int) *GetDedupeReportDefault {
	return &GetDedupeReportDefault{
		_statusCode: code,
	}
}

/*GetDedupeReportDefault handles this case with default header values.

Unexpected error
*/
type GetDedupeReportDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get dedupe report default response
func (o *GetDedupeReportDefault) Code() int {
	return o._statusCode
}

func (o *GetDedupeReportDefault) Error() string {
	return fmt.Sprintf("[GET /platform/1/dedupe/reports/{DedupeReportId}][%d] getDedupeReport default  %+v", o._statusCode, o.Payload)
}

func (o *GetDedupeReportDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
